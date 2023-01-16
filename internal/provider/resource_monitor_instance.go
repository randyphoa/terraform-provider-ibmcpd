package provider

import (
	"context"
	"encoding/json"
	"os"
	"reflect"

	"terraform-provider-ibmcpd/internal/go-sdk/client"
	"terraform-provider-ibmcpd/internal/go-sdk/watsonopenscalev2"
	"terraform-provider-ibmcpd/internal/utils"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &monitorInstanceResource{}
	_ resource.ResourceWithConfigure   = &monitorInstanceResource{}
	_ resource.ResourceWithImportState = &monitorInstanceResource{}
)

type monitorInstanceResource struct {
	client *client.Client
}

type monitorInstanceResourceModel struct {
	ID                  types.String      `tfsdk:"id"`
	DataMartID          types.String      `tfsdk:"data_mart_id"`
	SubscriptionID      types.String      `tfsdk:"subscription_id"`
	MonitorDefinitionID types.String      `tfsdk:"monitor_definition_id"`
	Parameters          *parametersModel  `tfsdk:"parameters"`
	Thresholds          []thresholdsModel `tfsdk:"thresholds"`

	DriftArchivePath types.String `tfsdk:"drift_archive_path"`
}

type parametersModel struct {
	MinFeedbackDataSize types.Int64 `tfsdk:"min_feedback_data_size"`

	MinSamples       types.Int64  `tfsdk:"min_samples"`
	DriftThreshold   types.Number `tfsdk:"drift_threshold"`
	TrainDriftModel  types.Bool   `tfsdk:"train_drift_model"`
	EnableModelDrift types.Bool   `tfsdk:"enable_model_drift"`
	EnableDataDrift  types.Bool   `tfsdk:"enable_data_drift"`

	Features          []fairnessFeatureModel `tfsdk:"features"`
	FavourableClass   []types.String         `tfsdk:"favourable_class"`
	UnfavourableClass []types.String         `tfsdk:"unfavourable_class"`
	MinRecords        types.Int64            `tfsdk:"min_records"`

	Enabled types.Bool `tfsdk:"enabled"`
}

type thresholdsModel struct {
	MetricID types.String `tfsdk:"metric_id"`
	Type     types.String `tfsdk:"type"`
	Value    types.Number `tfsdk:"value"`
}

type fairnessFeatureModel struct {
	Feature   types.String   `tfsdk:"feature"`
	Majority  []types.String `tfsdk:"majority"`
	Minority  []types.String `tfsdk:"minority"`
	Threshold types.Number   `tfsdk:"threshold"`
}

type parametersQuality struct {
	MinFeedbackDataSize int64 `json:"min_feedback_data_size"`
}

type parametersDrift struct {
	MinSamples       int64   `json:"min_samples"`
	DriftThreshold   float64 `json:"drift_threshold"`
	TrainDriftModel  bool    `json:"train_drift_model"`
	EnableModelDrift bool    `json:"enable_model_drift"`
	EnableDataDrift  bool    `json:"enable_data_drift"`
}

type parametersFairness struct {
	Features          []parametersFairnessFeature `json:"features"`
	FavourableClass   []string                    `json:"favourable_class"`
	UnfavourableClass []string                    `json:"unfavourable_class"`
	MinRecords        int64                       `json:"min_records"`
}

type parametersFairnessFeature struct {
	Feature   string   `json:"feature"`
	Majority  []string `json:"majority"`
	Minority  []string `json:"minority"`
	Threshold float64  `json:"threshold"`
}

type parametersExplainability struct {
	Enabled bool `json:"enabled"`
}

func NewMonitorInstanceResource() resource.Resource {
	return &monitorInstanceResource{}
}

func (r *monitorInstanceResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.Client)
}

func (r *monitorInstanceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monitor_instance"
}

func (r *monitorInstanceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"data_mart_id": schema.StringAttribute{
				Required: true,
			},
			"subscription_id": schema.StringAttribute{
				Required: true,
			},
			"monitor_definition_id": schema.StringAttribute{
				Required: true,
			},
			"drift_archive_path": schema.StringAttribute{
				Optional: true,
			},
			"parameters": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"min_feedback_data_size": schema.Int64Attribute{
						Optional: true,
					},
					"min_samples": schema.Int64Attribute{
						Optional: true,
					},
					"drift_threshold": schema.NumberAttribute{
						Optional: true,
					},
					"train_drift_model": schema.BoolAttribute{
						Optional: true,
					},
					"enable_model_drift": schema.BoolAttribute{
						Optional: true,
					},
					"enable_data_drift": schema.BoolAttribute{
						Optional: true,
					},
					"enabled": schema.BoolAttribute{
						Optional: true,
					},
					"features": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"feature": schema.StringAttribute{
									Optional: true,
								},
								"majority": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"minority": schema.ListAttribute{
									ElementType: types.StringType,
									Optional:    true,
								},
								"threshold": schema.NumberAttribute{
									Optional: true,
								},
							},
						},
					},
					"favourable_class": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"unfavourable_class": schema.ListAttribute{
						ElementType: types.StringType,
						Optional:    true,
					},
					"min_records": schema.Int64Attribute{
						Optional: true,
					},
				},
			},
			"thresholds": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"metric_id": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Optional: true,
						},
						"value": schema.NumberAttribute{
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func (r *monitorInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan monitorInstanceResourceModel
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

	var jsonParameters map[string]interface{}

	switch plan.MonitorDefinitionID.ValueString() {
	case "quality":
		qualityParameters, err := json.Marshal(parametersQuality{
			MinFeedbackDataSize: plan.Parameters.MinFeedbackDataSize.ValueInt64(),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Parsing Quality Parameters", "Could not parse quality parameters, unexpected error: "+err.Error())
			return
		}
		json.Unmarshal(qualityParameters, &jsonParameters)

	case "drift":
		driftThreshold, _ := plan.Parameters.DriftThreshold.ValueBigFloat().Float64()
		driftParameters, err := json.Marshal(parametersDrift{
			MinSamples:       plan.Parameters.MinSamples.ValueInt64(),
			DriftThreshold:   driftThreshold,
			TrainDriftModel:  plan.Parameters.TrainDriftModel.ValueBool(),
			EnableModelDrift: plan.Parameters.EnableModelDrift.ValueBool(),
			EnableDataDrift:  plan.Parameters.EnableDataDrift.ValueBool(),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Parsing Drift Parameters", "Could not parse drift parameters, unexpected error: "+err.Error())
			return
		}
		json.Unmarshal(driftParameters, &jsonParameters)

	case "fairness":
		features := make([]parametersFairnessFeature, len(plan.Parameters.Features))
		for i, v := range plan.Parameters.Features {
			fairnessThreshold, _ := v.Threshold.ValueBigFloat().Float64()
			features[i] = parametersFairnessFeature{
				Feature:   v.Feature.ValueString(),
				Majority:  utils.ConvertString(v.Majority),
				Minority:  utils.ConvertString(v.Minority),
				Threshold: fairnessThreshold,
			}
		}

		fairnessParameters, err := json.Marshal(parametersFairness{
			Features:          features,
			FavourableClass:   utils.ConvertString(plan.Parameters.FavourableClass),
			UnfavourableClass: utils.ConvertString(plan.Parameters.UnfavourableClass),
			MinRecords:        plan.Parameters.MinRecords.ValueInt64(),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Parsing Drift Parameters", "Could not parse drift parameters, unexpected error: "+err.Error())
			return
		}
		json.Unmarshal(fairnessParameters, &jsonParameters)

	case "explainability":
		explainabilityParameters, err := json.Marshal(parametersExplainability{
			Enabled: plan.Parameters.Enabled.ValueBool(),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Parsing Quality Parameters", "Could not parse quality parameters, unexpected error: "+err.Error())
			return
		}
		json.Unmarshal(explainabilityParameters, &jsonParameters)

	}

	thresholds := make([]watsonopenscalev2.MetricThresholdOverride, len(plan.Thresholds))
	for i, v := range plan.Thresholds {
		valueFloat, _ := v.Value.ValueBigFloat().Float64()
		thresholds[i] = watsonopenscalev2.MetricThresholdOverride{
			MetricID: core.StringPtr(v.MetricID.ValueString()),
			Type:     core.StringPtr(v.Type.ValueString()),
			Value:    core.Float64Ptr(valueFloat),
		}
	}

	if plan.DriftArchivePath.ValueString() != "" {
		file, err := os.Open(plan.DriftArchivePath.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Unable to open drift archive", err.Error())
			return
		}
		defer file.Close()
		response, err := wosClient.DriftArchivePost(&watsonopenscalev2.DriftArchivePostOptions{
			DataMartID:       core.StringPtr(plan.DataMartID.ValueString()),
			SubscriptionID:   core.StringPtr(plan.SubscriptionID.ValueString()),
			Body:             file,
			ArchiveName:      core.StringPtr("user_drift.tar.gz"),
			EnableDataDrift:  core.BoolPtr(true),
			EnableModelDrift: core.BoolPtr(true),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Uploading Drift Archive", "Could not upload drift archive, unexpected error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to Upload Drift Archive", err.Error())
			return
		}
	}

	result, response, err := wosClient.InstancesAdd(&watsonopenscalev2.InstancesAddOptions{
		DataMartID:          core.StringPtr(plan.DataMartID.ValueString()),
		MonitorDefinitionID: core.StringPtr(plan.MonitorDefinitionID.ValueString()),
		Target: &watsonopenscalev2.Target{
			TargetID:   core.StringPtr(plan.SubscriptionID.ValueString()),
			TargetType: core.StringPtr("subscription"),
		},
		Parameters: utils.If(plan.MonitorDefinitionID.ValueString() != "mrm", jsonParameters, map[string]interface{}{}),
		Thresholds: utils.If(plan.MonitorDefinitionID.ValueString() != "mrm", thresholds, nil),
	})

	if err != nil {
		resp.Diagnostics.AddError("Error Creating Instance", "Could not create instance, unexpected error: "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Create Instance", err.Error())
		return
	}

	plan.ID = types.StringValue(*result.Metadata.ID)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *monitorInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state monitorInstanceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		panic(err)
	}

	monitorInstance, response, err := wosClient.InstancesGet(&watsonopenscalev2.InstancesGetOptions{
		MonitorInstanceID: core.StringPtr(state.ID.ValueString()),
	})
	if utils.Contains(utils.HTTP_OK, response.StatusCode) {
		state.ID = types.StringValue(*monitorInstance.Metadata.ID)
		state.MonitorDefinitionID = types.StringValue(*monitorInstance.Entity.MonitorDefinitionID)
		diags = resp.State.Set(ctx, &state)
	} else if response.StatusCode == 404 {
		diags = resp.State.Set(ctx, &monitorInstanceResourceModel{})
	} else {
		resp.Diagnostics.AddError("Error Getting Monitor Instance", "Could not read Monitor Instance ID "+state.ID.ValueString()+": "+err.Error())
		return
	}

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *monitorInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// var updateableFields = []string{"Thresholds"}
	var diags diag.Diagnostics

	var state monitorInstanceResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan monitorInstanceResourceModel
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		panic(err)
	}

	var jsonPatches []watsonopenscalev2.PatchDocument

	if !reflect.DeepEqual(plan.Thresholds, state.Thresholds) {
		thresholds := make([]watsonopenscalev2.MetricThresholdOverride, len(plan.Thresholds))
		for i, v := range plan.Thresholds {
			valueFloat, _ := v.Value.ValueBigFloat().Float64()
			thresholds[i] = watsonopenscalev2.MetricThresholdOverride{
				MetricID: core.StringPtr(v.MetricID.ValueString()),
				Type:     core.StringPtr(v.Type.ValueString()),
				Value:    core.Float64Ptr(valueFloat),
			}
		}
		// jsonThresholds, err := json.Marshal(thresholds)
		// if err != nil {
		// 	resp.Diagnostics.AddError("Unable to parse thresholds", err.Error())
		// 	return
		// }
		jsonPatches = append(jsonPatches, watsonopenscalev2.PatchDocument{
			Op:    core.StringPtr("replace"),
			Path:  core.StringPtr("/thresholds"),
			Value: thresholds,
		})

	}
	// for _, field := range updateableFields {
	// 	if utils.GetAttr(&plan, field).Interface().(types.String).ValueString() != utils.GetAttr(&state, field).Interface().(types.String).ValueString() {
	// 		switch field {
	// 		case "Thresholds":
	// 			thresholds := make([]watsonopenscalev2.MetricThresholdOverride, len(plan.Thresholds))
	// 			for i, v := range plan.Thresholds {
	// 				valueFloat, _ := v.Value.ValueBigFloat().Float64()
	// 				thresholds[i] = watsonopenscalev2.MetricThresholdOverride{
	// 					MetricID: core.StringPtr(v.MetricID.ValueString()),
	// 					Type:     core.StringPtr(v.Type.ValueString()),
	// 					Value:    core.Float64Ptr(valueFloat),
	// 				}
	// 			}

	// 			jsonThresholds, err := json.Marshal(thresholds)
	// 			if err != nil {
	// 				resp.Diagnostics.AddError("Unable to parse thresholds", err.Error())
	// 				return
	// 			}

	// 			jsonPatches = append(jsonPatches, watsonopenscalev2.JSONPatchOperation{
	// 				Op:    core.StringPtr("replace"),
	// 				Path:  core.StringPtr(fmt.Sprintf(`/%s`, strings.ToLower(field))),
	// 				Value: string(jsonThresholds),
	// 			})

	// 			// default:
	// 			// 	jsonPatches = append(jsonPatches, watsonopenscalev2.JSONPatchOperation{
	// 			// 		Op:    core.StringPtr("replace"),
	// 			// 		Path:  core.StringPtr(fmt.Sprintf(`/%s`, strings.ToLower(field))),
	// 			// 		Value: core.StringPtr(utils.GetAttr(&plan, field).Interface().(types.String).ValueString()),
	// 			// 	})
	// 		}
	// 	}
	// }

	result, response, err := wosClient.InstancesUpdate(&watsonopenscalev2.InstancesUpdateOptions{
		MonitorInstanceID: core.StringPtr(plan.ID.ValueString()),
		PatchDocument:     jsonPatches,
	})

	if err != nil {
		resp.Diagnostics.AddError("Error Updating Monitor Instance", "Could not update monitor instance ID "+plan.ID.ValueString()+": "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Update Monitor Instance", err.Error())
		return
	}

	plan.ID = types.StringValue(*result.Metadata.ID)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *monitorInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state monitorInstanceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		panic(err)
	}

	response, err := wosClient.InstancesDelete(&watsonopenscalev2.InstancesDeleteOptions{
		MonitorInstanceID: core.StringPtr(state.ID.ValueString()),
	})

	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Monitor Instance", "Could not delete monitor instance ID "+state.ID.ValueString()+": "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Delete Monitor Instance", err.Error())
		return
	}
}

func (r *monitorInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
