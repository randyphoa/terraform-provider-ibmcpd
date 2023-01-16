package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"terraform-provider-ibmcpd/internal/go-sdk/client"
	"terraform-provider-ibmcpd/internal/go-sdk/watsonmachinelearningv4"
	"terraform-provider-ibmcpd/internal/go-sdk/watsonopenscalev2"
	"terraform-provider-ibmcpd/internal/utils"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sagemakerruntime"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const NUM_TRIES_SUBSCRIPTION = 30
const TIMEOUT_SUBSCRIPTION = 5 * time.Second

var (
	_ resource.Resource                = &subscriptionResource{}
	_ resource.ResourceWithConfigure   = &subscriptionResource{}
	_ resource.ResourceWithImportState = &subscriptionResource{}
	// _ resource.ResourceWithConfigValidators = &subscriptionResource{}
)

type subscriptionResource struct {
	client *client.Client
}

type subscriptionResourceModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`

	DataMartID        types.String `tfsdk:"data_mart_id"`
	ServiceProviderID types.String `tfsdk:"service_provider_id"`

	Asset                 *subscriptionAssetModel                 `tfsdk:"asset"`
	Deployment            *subscriptionDeploymentModel            `tfsdk:"deployment"`
	AssetProperties       *subscriptionAssetPropertiesModel       `tfsdk:"asset_properties"`
	TrainingDataReference *subscriptionTrainingDataReferenceModel `tfsdk:"training_data_reference"`
	TrainingDataSchema    []sparkStructFieldModel                 `tfsdk:"training_data_schema"`

	PayloadFile types.String `tfsdk:"payload_file"`

	AWSAccessKeyID     types.String `tfsdk:"aws_access_key_id"`
	AWSSecretAccessKey types.String `tfsdk:"aws_secret_access_key"`
	AWSRegion          types.String `tfsdk:"aws_region"`
}

type sparkStructFieldModel struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	Nullable types.Bool   `tfsdk:"nullable"`
}

type subscriptionAssetModel struct {
	AssetID       types.String `tfsdk:"asset_id"`
	AssetType     types.String `tfsdk:"asset_type"`
	InputDataType types.String `tfsdk:"input_data_type"`
	ProblemType   types.String `tfsdk:"problem_type"`
	URL           types.String `tfsdk:"url"`
}

type subscriptionDeploymentModel struct {
	DeploymentID   types.String `tfsdk:"deployment_id"`
	DeploymentType types.String `tfsdk:"deployment_type"`
	URL            types.String `tfsdk:"deployment_url"`
	ScoringURL     types.String `tfsdk:"scoring_url"`
	Name           types.String `tfsdk:"name"`
}

type subscriptionAssetPropertiesModel struct {
	CategoricalFields []types.String `tfsdk:"categorical_fields"`
	FeatureFields     []types.String `tfsdk:"feature_fields"`
	LabelColumn       types.String   `tfsdk:"label_column"`
	PredictionField   types.String   `tfsdk:"prediction_field"`
	ProbabilityFields []types.String `tfsdk:"probability_fields"`
}

type subscriptionTrainingDataReferenceModel struct {
	Type types.String `tfsdk:"type"`

	DatabaseName types.String `tfsdk:"database_name"`
	Hostname     types.String `tfsdk:"hostname"`
	Password     types.String `tfsdk:"password"`
	Port         types.Int64  `tfsdk:"port"`
	Ssl          types.Bool   `tfsdk:"ssl"`
	Username     types.String `tfsdk:"username"`
	SchemaName   types.String `tfsdk:"schema_name"`
	TableName    types.String `tfsdk:"table_name"`

	BucketName         types.String `tfsdk:"bucket_name"`
	FileName           types.String `tfsdk:"file_name"`
	ResourceInstanceID types.String `tfsdk:"resource_instance_id"`
	COSApiKey          types.String `tfsdk:"cos_api_key"`
	COSURL             types.String `tfsdk:"cos_url"`
	IamURL             types.String `tfsdk:"iam_url"`
}

func NewSubscriptionResource() resource.Resource {
	return &subscriptionResource{}
}

func (r *subscriptionResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.Client)
}

func (r *subscriptionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscription"
}

// func (r *subscriptionResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
// 	return []resource.ConfigValidator{
// 		resourcevalidator.AtLeastOneOf(
// 			path.MatchRoot("project_id"),
// 			path.MatchRoot("space_id"),
// 		),
// 	}
// }

func (r *subscriptionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"data_mart_id": schema.StringAttribute{
				Required: true,
			},
			"service_provider_id": schema.StringAttribute{
				Required: true,
			},
			"asset": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"asset_id": schema.StringAttribute{
						Required: true,
					},
					"asset_type": schema.StringAttribute{
						Required: true,
					},
					"input_data_type": schema.StringAttribute{
						Required: true,
					},
					"problem_type": schema.StringAttribute{
						Required: true,
					},
					"url": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"deployment": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"deployment_id": schema.StringAttribute{
						Required: true,
					},
					"deployment_type": schema.StringAttribute{
						Required: true,
					},
					"deployment_url": schema.StringAttribute{
						Optional: true,
					},
					"scoring_url": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			"asset_properties": schema.SingleNestedAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Attributes: map[string]schema.Attribute{
					"categorical_fields": schema.ListAttribute{
						ElementType: types.StringType,
						Required:    true,
					},
					"feature_fields": schema.ListAttribute{
						ElementType: types.StringType,
						Required:    true,
					},
					"label_column": schema.StringAttribute{
						Required: true,
					},
					"prediction_field": schema.StringAttribute{
						Required: true,
					},
					"probability_fields": schema.ListAttribute{
						ElementType: types.StringType,
						Required:    true,
					},
				},
			},
			"training_data_reference": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Optional: true,
					},
					"database_name": schema.StringAttribute{
						Optional: true,
					},
					"hostname": schema.StringAttribute{
						Optional: true,
					},
					"password": schema.StringAttribute{
						Optional:  true,
						Sensitive: true,
					},
					"port": schema.Int64Attribute{
						Optional: true,
					},
					"ssl": schema.BoolAttribute{
						Optional: true,
					},
					"username": schema.StringAttribute{
						Optional: true,
					},
					"schema_name": schema.StringAttribute{
						Optional: true,
					},
					"table_name": schema.StringAttribute{
						Optional: true,
					},
					"bucket_name": schema.StringAttribute{
						Optional: true,
					},
					"file_name": schema.StringAttribute{
						Optional: true,
					},
					"resource_instance_id": schema.StringAttribute{
						Optional: true,
					},
					"cos_url": schema.StringAttribute{
						Optional: true,
					},
					"cos_api_key": schema.StringAttribute{
						Optional:  true,
						Sensitive: true,
					},
					"iam_url": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			"training_data_schema": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"type": schema.StringAttribute{
							Required: true,
						},
						"nullable": schema.BoolAttribute{
							Required: true,
						},
					},
				},
			},
			"payload_file": schema.StringAttribute{
				Optional: true,
			},
			"aws_access_key_id": schema.StringAttribute{
				Optional: true,
			},
			"aws_secret_access_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"aws_region": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *subscriptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan subscriptionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to get WOS Client", err.Error())
		return
	}

	fields := make([]watsonopenscalev2.SparkStructField, len(plan.TrainingDataSchema))
	if len(plan.TrainingDataSchema) > 0 {
		for i, v := range plan.TrainingDataSchema {
			fields[i] = watsonopenscalev2.SparkStructField{
				Name:     core.StringPtr(v.Name.ValueString()),
				Type:     core.StringPtr(v.Type.ValueString()),
				Nullable: core.BoolPtr(v.Nullable.ValueBool()),
			}
		}
	}

	var connection watsonopenscalev2.TrainingDataReferenceConnectionIntf
	var location watsonopenscalev2.TrainingDataReferenceLocationIntf

	if plan.TrainingDataReference.Type.ValueString() == "db2" {
		connection = &watsonopenscalev2.DB2TrainingDataReferenceConnection{
			DatabaseName: core.StringPtr(plan.TrainingDataReference.DatabaseName.ValueString()),
			Hostname:     core.StringPtr(plan.TrainingDataReference.Hostname.ValueString()),
			Password:     core.StringPtr(plan.TrainingDataReference.Password.ValueString()),
			Port:         core.Int64Ptr(plan.TrainingDataReference.Port.ValueInt64()),
			Ssl:          core.BoolPtr(plan.TrainingDataReference.Ssl.ValueBool()),
			Username:     core.StringPtr(plan.TrainingDataReference.Username.ValueString()),
		}
		location = &watsonopenscalev2.DB2TrainingDataReferenceLocation{
			SchemaName: core.StringPtr(plan.TrainingDataReference.SchemaName.ValueString()),
			TableName:  core.StringPtr(plan.TrainingDataReference.TableName.ValueString()),
		}
	}

	if plan.TrainingDataReference.Type.ValueString() == "cos" {
		connection = &watsonopenscalev2.COSTrainingDataReferenceConnection{
			ResourceInstanceID: core.StringPtr(plan.TrainingDataReference.ResourceInstanceID.ValueString()),
			ApiKey:             core.StringPtr(plan.TrainingDataReference.COSApiKey.ValueString()),
			IamURL:             core.StringPtr(plan.TrainingDataReference.IamURL.ValueString()),
			URL:                core.StringPtr(plan.TrainingDataReference.COSURL.ValueString()),
		}
		location = &watsonopenscalev2.COSTrainingDataReferenceLocation{
			Bucket:   core.StringPtr(plan.TrainingDataReference.BucketName.ValueString()),
			FileName: core.StringPtr(plan.TrainingDataReference.FileName.ValueString()),
		}
	}

	result, response, err := wosClient.SubscriptionsAdd(&watsonopenscalev2.SubscriptionsAddOptions{
		Asset: &watsonopenscalev2.Asset{
			AssetID:       core.StringPtr(plan.Asset.AssetID.ValueString()),
			AssetType:     core.StringPtr(plan.Asset.AssetType.ValueString()),
			InputDataType: core.StringPtr(plan.Asset.InputDataType.ValueString()),
			ProblemType:   core.StringPtr(plan.Asset.ProblemType.ValueString()),
			URL:           core.StringPtr(plan.Asset.URL.ValueString()),
		},
		DataMartID: core.StringPtr(plan.DataMartID.ValueString()),
		Deployment: &watsonopenscalev2.AssetDeploymentRequest{
			DeploymentID:   core.StringPtr(plan.Deployment.DeploymentID.ValueString()),
			DeploymentType: core.StringPtr(plan.Deployment.DeploymentType.ValueString()),
			Name:           core.StringPtr(plan.Name.ValueString()),
			URL:            core.StringPtr(plan.Deployment.URL.ValueString()),
			ScoringEndpoint: &watsonopenscalev2.ScoringEndpointRequest{
				URL: core.StringPtr(plan.Deployment.ScoringURL.ValueString()),
			},
		},
		ServiceProviderID: core.StringPtr(plan.ServiceProviderID.ValueString()),
		AssetProperties: &watsonopenscalev2.AssetPropertiesRequest{
			CategoricalFields: utils.ConvertString(plan.AssetProperties.CategoricalFields),
			FeatureFields:     utils.ConvertString(plan.AssetProperties.FeatureFields),
			LabelColumn:       core.StringPtr(plan.AssetProperties.LabelColumn.ValueString()),
			PredictionField:   core.StringPtr(plan.AssetProperties.PredictionField.ValueString()),
			ProbabilityFields: utils.ConvertString(plan.AssetProperties.ProbabilityFields),
			TrainingDataReference: &watsonopenscalev2.TrainingDataReference{
				Connection: connection,
				Location:   location,
				Type:       core.StringPtr(plan.TrainingDataReference.Type.ValueString()),
			},
			TrainingDataSchema: &watsonopenscalev2.SparkStruct{
				ID:     core.StringPtr("input_data_schema"),
				Type:   core.StringPtr("struct"),
				Fields: fields,
			},
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Subscription", "Could not create subscription, unexpected error: "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Create Subscription", err.Error())
		return
	}

	tflog.Info(ctx, "Created Subscription", map[string]interface{}{"subscription_id": result.Metadata.ID})

	var subscriptionStatus string
	for i := 1; i < NUM_TRIES_SUBSCRIPTION; i++ {
		time.Sleep(TIMEOUT_SUBSCRIPTION)
		resultSubscription, response, err := wosClient.SubscriptionsGet(&watsonopenscalev2.SubscriptionsGetOptions{
			SubscriptionID: result.Metadata.ID,
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Getting Subscription", "Could not get subscription, unexpected error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to Get Subscription", err.Error())
			return
		}
		subscriptionStatus = *resultSubscription.Entity.Status.State
		tflog.Info(ctx, "Subscription Status", map[string]interface{}{"subscription_id": result.Metadata.ID, "status": subscriptionStatus})
		if subscriptionStatus == "active" {
			break
		}
	}

	if subscriptionStatus != "active" {
		wosClient.SubscriptionsDelete(&watsonopenscalev2.SubscriptionsDeleteOptions{
			SubscriptionID: result.Metadata.ID,
		})
		resp.Diagnostics.AddError("Error Creating Subscription", "Subscription ID "+*result.Metadata.ID+" status is not active. Status is "+subscriptionStatus)
		return
	}

	var dataSetID *string
	var numPayloadRecords int
	for i := 1; i < NUM_TRIES_SUBSCRIPTION; i++ {
		time.Sleep(TIMEOUT_SUBSCRIPTION)
		resultDataSets, response, err := wosClient.DataSetsList(&watsonopenscalev2.DataSetsListOptions{
			TargetTargetID:   result.Metadata.ID,
			Type:             core.StringPtr("payload_logging"),
			TargetTargetType: core.StringPtr("subscription"),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Listing Datasets", "Could not list datasets, unexpected error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to List Datasets", err.Error())
			return
		}
		if len(resultDataSets.DataSets) > 0 {
			dataSetID = resultDataSets.DataSets[0].Metadata.ID
			tflog.Info(ctx, "Created Payload Dataset", map[string]interface{}{"dataset_id": dataSetID})
			break
		}
	}
	if *dataSetID == "" {
		wosClient.SubscriptionsDelete(&watsonopenscalev2.SubscriptionsDeleteOptions{
			SubscriptionID: result.Metadata.ID,
		})
		resp.Diagnostics.AddError("Error Creating Payload Dataset", "Unable to create payload dataset for Subscription ID "+*result.Metadata.ID+".")
		return
	}

	serviceProvider, response, err := wosClient.ServiceProvidersGet(&watsonopenscalev2.ServiceProvidersGetOptions{
		ServiceProviderID: core.StringPtr(plan.ServiceProviderID.ValueString()),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error Getting Service Provider", "Could not get service provider, unexpected error: "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Get Service Provider", err.Error())
		return
	}

	var jsonPayload map[string][]interface{}
	var scoringFields []string
	var payloadContent []byte

	if plan.PayloadFile.ValueString() != "" {
		file, err := os.Open(plan.PayloadFile.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Unable to open payload file", err.Error())
			return
		}
		defer file.Close()
		payloadContent, err = io.ReadAll(file)
		if err != nil {
			resp.Diagnostics.AddError("Unable to read payload file", err.Error())
			return
		}
		json.Unmarshal(payloadContent, &jsonPayload)
		numPayloadRecords = len(jsonPayload["values"])
		scoringFields = make([]string, len(jsonPayload["fields"]))
		for i, v := range jsonPayload["fields"] {
			scoringFields[i] = v.(string)
		}
	}

	if *serviceProvider.Entity.ServiceType == "amazon_sagemaker" {
		scoringValues := make([][]interface{}, len(jsonPayload["values"]))
		for i, v := range jsonPayload["values"] {
			scoringValues[i] = v.([]interface{})
		}

		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(plan.AWSAccessKeyID.ValueString(), plan.AWSSecretAccessKey.ValueString(), ""),
			Region:      aws.String(plan.AWSRegion.ValueString())},
		)
		if err != nil {
			resp.Diagnostics.AddError("Unable to create AWS session", err.Error())
			return
		}

		scoringPayload := fmt.Sprintf(`{"input_data": [%s]}`, string(payloadContent))
		svc := sagemakerruntime.New(sess)
		response, err := svc.InvokeEndpoint(&sagemakerruntime.InvokeEndpointInput{
			EndpointName: aws.String("pipeline-endpoint"),
			ContentType:  aws.String("application/json"),
			Body:         []byte(scoringPayload),
		})
		if err != nil {
			resp.Diagnostics.AddError("Unable to score AWS endpoint", err.Error())
			return
		}

		var jsonResponse map[string][]interface{}
		err = json.Unmarshal(response.Body, &jsonResponse)
		if err != nil {
			resp.Diagnostics.AddError("Unable to parse AWS endpoint response", err.Error())
			return
		}

		values := []string{}
		for _, v := range jsonResponse["predictions"] {
			val := v.(map[string]interface{})
			values = append(values, fmt.Sprintf(`[%f, "%s"]`, val["score"], val["predicted_label"]))
		}

		var jsonValues []interface{}
		err = json.Unmarshal([]byte(fmt.Sprintf("[%s]", strings.Join(values, ","))), &jsonValues)
		if err != nil {
			panic(err)
		}

		_, resultResponse, err := wosClient.RecordsAdd(&watsonopenscalev2.RecordsAddOptions{
			DataSetID: dataSetID,
			DatasetRecordsPayloadItem: []watsonopenscalev2.DatasetRecordsPayloadItemIntf{
				&watsonopenscalev2.DatasetRecordsPayloadItem{
					Request: &watsonopenscalev2.ScoringPayloadRequestRequest{
						Fields: scoringFields,
						Values: scoringValues,
					},
					Response: &watsonopenscalev2.ScoringPayloadRequestResponse{
						Fields: []string{"score", "predicted_label"},
						Values: jsonValues,
					},
				},
			},
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Storing Payload", "Could not store payload, unexpected error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, resultResponse.StatusCode) {
			resp.Diagnostics.AddError("Unable to Store Payload", err.Error())
			return
		}
	}

	if *serviceProvider.Entity.ServiceType == "watson_machine_learning" && plan.PayloadFile.ValueString() != "" {
		wmlClient, err := r.client.WMLClient(context.Background())
		if err != nil {
			resp.Diagnostics.AddError("Unable to get WML Client", err.Error())
			return
		}

		objValues, err := json.Marshal(jsonPayload["values"])
		if err != nil {
			resp.Diagnostics.AddError("Unable to parse payload file", err.Error())
			return
		}
		var scoringValues [][]interface{}
		json.Unmarshal(objValues, &scoringValues)

		_, response, err = wmlClient.DeploymentsComputePredictions(&watsonmachinelearningv4.DeploymentsComputePredictionsOptions{
			DeploymentID: core.StringPtr(plan.Deployment.DeploymentID.ValueString()),
			InputData: []watsonmachinelearningv4.InputDataArray{
				{Fields: scoringFields, Values: scoringValues},
			},
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Scoring Payload", "Could not score payload, unexpected error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to Score Payload", err.Error())
			return
		}
	}

	for i := 1; i < NUM_TRIES_SUBSCRIPTION; i++ {
		time.Sleep(TIMEOUT_SUBSCRIPTION)
		resultRecordsList, response, err := wosClient.RecordsList(&watsonopenscalev2.RecordsListOptions{
			DataSetID:         dataSetID,
			IncludeTotalCount: core.BoolPtr(true),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Listing Payload Records", "Could not list payload records, unexpected error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to List Payload Records", err.Error())
			return
		}
		records := resultRecordsList.(*watsonopenscalev2.RecordsListResponse)
		if int(*records.TotalCount) >= numPayloadRecords {
			tflog.Info(ctx, "Stored Payload Dataset", map[string]interface{}{"payload_records": *records.TotalCount})
			break
		}
	}

	plan.ID = types.StringValue(*result.Metadata.ID)

	if _, err := os.Stat(plan.PayloadFile.ValueString()); err == nil || os.IsExist(err) {
		os.Remove(plan.PayloadFile.ValueString())
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *subscriptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state subscriptionResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		panic(err)
	}

	subscription, response, err := wosClient.SubscriptionsGet(&watsonopenscalev2.SubscriptionsGetOptions{
		SubscriptionID: core.StringPtr(state.ID.ValueString()),
	})

	if utils.Contains(utils.HTTP_OK, response.StatusCode) {
		state.ID = types.StringValue(*subscription.Metadata.ID)
		state.Name = types.StringValue(*subscription.Entity.Deployment.Name)
		diags = resp.State.Set(ctx, &state)
	} else if response.StatusCode == 404 {
		diags = resp.State.Set(ctx, &subscriptionResourceModel{})
	} else {
		resp.Diagnostics.AddError("Error Getting Subscription", "Could not read Subscription ID "+state.ID.ValueString()+": "+err.Error())
		return
	}

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *subscriptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// var updateableFields = []string{"Name"}
	var diags diag.Diagnostics

	var state subscriptionResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan subscriptionResourceModel
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jsonPatches []watsonopenscalev2.PatchDocument

	// for _, field := range updateableFields {
	// 	if utils.GetAttr(&plan, field).Interface().(types.String).ValueString() != utils.GetAttr(&state, field).Interface().(types.String).ValueString() {
	// 		jsonPatches = append(jsonPatches, watsonopenscalev2.PatchDocument{
	// 			Op:    core.StringPtr("replace"),
	// 			Path:  core.StringPtr(fmt.Sprintf(`/%s`, strings.ToLower(field))),
	// 			Value: core.StringPtr(utils.GetAttr(&plan, field).Interface().(types.String).ValueString()),
	// 		})
	// 	}
	// }
	if plan.Name.ValueString() != state.Name.ValueString() {
		jsonPatches = append(jsonPatches, watsonopenscalev2.PatchDocument{
			Op:    core.StringPtr("replace"),
			Path:  core.StringPtr("/deployment/name"),
			Value: plan.Name.ValueString(),
		})
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		panic(err)
	}

	result, response, err := wosClient.SubscriptionsUpdate(&watsonopenscalev2.SubscriptionsUpdateOptions{
		SubscriptionID: core.StringPtr(state.ID.ValueString()),
		PatchDocument:  jsonPatches,
	})

	if err != nil {
		resp.Diagnostics.AddError("Error Updating Subscription", "Could not update subscription ID "+plan.ID.ValueString()+": "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Update Subscription", err.Error())
		return
	}

	plan.ID = types.StringValue(*result.Metadata.ID)
	// plan.Name = types.StringValue(*result.Entity.Deployment.Name)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *subscriptionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state subscriptionResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		panic(err)
	}

	response, err := wosClient.SubscriptionsDelete(&watsonopenscalev2.SubscriptionsDeleteOptions{
		SubscriptionID: core.StringPtr(state.ID.ValueString()),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Subscription", "Could not delete subscription ID "+state.ID.ValueString()+": "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Delete Subscription", err.Error())
		return
	}
}

func (r *subscriptionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
