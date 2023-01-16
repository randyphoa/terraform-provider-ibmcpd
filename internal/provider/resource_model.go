package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"terraform-provider-ibmcpd/internal/go-sdk/client"
	"terraform-provider-ibmcpd/internal/go-sdk/watsonmachinelearningv4"
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
	_ resource.Resource                = &modelResource{}
	_ resource.ResourceWithConfigure   = &modelResource{}
	_ resource.ResourceWithImportState = &modelResource{}
)

type modelResource struct {
	client *client.Client
}

type modelResourceModel struct {
	ID           types.String       `tfsdk:"id"`
	Name         types.String       `tfsdk:"name"`
	Type         types.String       `tfsdk:"type"`
	ModelPath    types.String       `tfsdk:"model_path"`
	InputSchema  []inputSchemaModel `tfsdk:"input_schema"`
	LabelColumn  types.String       `tfsdk:"label_column"`
	SoftwareSpec types.String       `tfsdk:"software_spec"`
	ProjectID    types.String       `tfsdk:"project_id"`
	SpaceID      types.String       `tfsdk:"space_id"`
	AssetID      types.String       `tfsdk:"asset_id"`
	Checksum     types.String       `tfsdk:"checksum"`
}

type inputSchemaModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type inputSchemaJsonModel struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewModelResource() resource.Resource {
	return &modelResource{}
}

func (r *modelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.Client)
}

func (r *modelResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_model"
}

func (r *modelResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a model on IBM Cloud Pak for Data.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier for model.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of model.",
				Required:    true,
			},
			"type": schema.StringAttribute{
				Description: "Type of model.",
				Required:    true,
			},
			"software_spec": schema.StringAttribute{
				Description: "Software spec of model.",
				Required:    true,
			},
			"model_path": schema.StringAttribute{
				Description: "tar.gz file of model from joblib.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"input_schema": schema.ListNestedAttribute{
				Description: "Training input schema of model.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional: true,
						},
						"type": schema.StringAttribute{
							Optional: true,
						},
					},
				},
			},
			"label_column": schema.StringAttribute{
				Description: "Label column of model.",
				Optional:    true,
			},
			"project_id": schema.StringAttribute{
				Description: "Project ID of model.",
				Optional:    true,
			},
			"space_id": schema.StringAttribute{
				Description: "Space ID of model.",
				Optional:    true,
			},
			"asset_id": schema.StringAttribute{
				Description: "Asset ID of model in project. Used when promoting model in project to deployment space.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"checksum": schema.StringAttribute{
				Description: "Checksum of Python object.",
				Optional:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *modelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan modelResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wmlClient, err := r.client.WMLClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to get WML Client", err.Error())
		return
	}

	if plan.AssetID.ValueString() != "" && plan.SpaceID.ValueString() != "" {
		pathParamsMap := map[string]string{
			"asset_id": plan.AssetID.ValueString(),
		}
		builder := core.NewRequestBuilder(core.POST)
		builder = builder.WithContext(context.Background())

		url := utils.If(strings.Contains(r.client.Config.URL, "cloud.ibm.com"), "https://dataplatform.cloud.ibm.com", r.client.Config.URL)

		_, err = builder.ResolveRequestURL(url, `/projects/api/rest/catalogs/assets/{asset_id}/promote`, pathParamsMap)
		if err != nil {
			resp.Diagnostics.AddError("Unable to build promote URL", err.Error())
			return
		}
		builder.AddHeader("Content-Type", "application/json")
		builder.AddQuery("project_id", plan.ProjectID.ValueString())

		requestBody := fmt.Sprintf(`{"spaceId": "%s", "projectId": "%s", "assetType": "wml_model"}`, plan.SpaceID.ValueString(), plan.ProjectID.ValueString())
		var jsonRequestBody map[string]interface{}
		json.Unmarshal([]byte(requestBody), &jsonRequestBody)
		_, err = builder.SetBodyContentJSON(jsonRequestBody)
		if err != nil {
			resp.Diagnostics.AddError("Unable to set body content", err.Error())
			return
		}
		request, err := builder.Build()
		if err != nil {
			resp.Diagnostics.AddError("Unable to create URL builder", err.Error())
			return
		}
		err = wmlClient.Service.Options.Authenticator.Authenticate(request)
		if err != nil {
			resp.Diagnostics.AddError("Authentication error", err.Error())
			return
		}
		response, err := wmlClient.Service.Client.Do(request)
		if err != nil {
			resp.Diagnostics.AddError("Error received from HTTP client", err.Error())
			return
		}
		defer response.Body.Close()
		content, err := io.ReadAll(response.Body)
		if err != nil {
			resp.Diagnostics.AddError("Error reading from HTTP response", err.Error())
			return
		}
		var jsonResponse map[string]interface{}
		err = json.Unmarshal(content, &jsonResponse)
		if err != nil {
			resp.Diagnostics.AddError("Error parsing response json", err.Error())
			return
		}
		if jsonResponse == nil {
			resp.Diagnostics.AddError("Unable to promote model", err.Error())
			return
		}
		plan.ID = types.StringValue(jsonResponse["promotedAsset"].(map[string]interface{})["asset_id"].(string))
	} else {
		var schemas = &watsonmachinelearningv4.ModelEntitySchemas{}
		if len(plan.InputSchema) > 0 {
			var jsonInputSchema []interface{}
			inputSchema := make([]inputSchemaJsonModel, len(plan.InputSchema))
			for i, v := range plan.InputSchema {
				inputSchema[i] = inputSchemaJsonModel{
					Name: v.Name.ValueString(),
					Type: v.Type.ValueString(),
				}
			}
			objInputSchema, err := json.Marshal(inputSchema)
			if err != nil {
				resp.Diagnostics.AddError("Unable to get parse input schema", err.Error())
				return
			}
			json.Unmarshal(objInputSchema, &jsonInputSchema)
			schemas = &watsonmachinelearningv4.ModelEntitySchemas{
				Input: []watsonmachinelearningv4.DataSchema{
					{
						ID:     core.StringPtr("input_data_schema"),
						Fields: jsonInputSchema,
					},
				},
			}
		}
		model, response, err := wmlClient.ModelsCreate(&watsonmachinelearningv4.ModelsCreateOptions{
			Name: core.StringPtr(plan.Name.ValueString()),
			Type: core.StringPtr(plan.Type.ValueString()),
			SoftwareSpec: &watsonmachinelearningv4.SoftwareSpecRel{
				Name: core.StringPtr(plan.SoftwareSpec.ValueString()),
			},
			SpaceID:     utils.If(plan.SpaceID.ValueString() != "", core.StringPtr(plan.SpaceID.ValueString()), nil),
			ProjectID:   utils.If(plan.ProjectID.ValueString() != "", core.StringPtr(plan.ProjectID.ValueString()), nil),
			LabelColumn: core.StringPtr(plan.LabelColumn.ValueString()),
			Schemas:     schemas,
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Creating Model", "Could not create model, unexpected error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to Create model", err.Error())
			return
		}

		file, err := os.Open(plan.ModelPath.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Unable to Read Model File", err.Error())
			return
		}
		defer file.Close()
		_, response, err = wmlClient.ModelsUploadContent(&watsonmachinelearningv4.ModelsUploadContentOptions{
			ModelID:       model.Metadata.ID,
			ContentFormat: core.StringPtr("native"),
			SpaceID:       utils.If(plan.SpaceID.ValueString() != "", core.StringPtr(plan.SpaceID.ValueString()), nil),
			ProjectID:     utils.If(plan.ProjectID.ValueString() != "", core.StringPtr(plan.ProjectID.ValueString()), nil),
			Body:          file,
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Creating Model", "Could not create model, unexpected error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to Create model", err.Error())
			return
		}

		plan.ID = types.StringValue(*model.Metadata.ID)

		if _, err := os.Stat(plan.ModelPath.ValueString()); err == nil || os.IsExist(err) {
			os.Remove(plan.ModelPath.ValueString())
		}
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *modelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state modelResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if state.ID.ValueString() == "" {
		resp.Diagnostics.AddError("Error Getting Model", "Model ID is null.")
		return
	}

	wmlClient, err := r.client.WMLClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to get WML Client", err.Error())
		return
	}

	model, response, err := wmlClient.ModelsGet(&watsonmachinelearningv4.ModelsGetOptions{
		ModelID:   core.StringPtr(state.ID.ValueString()),
		SpaceID:   utils.If(state.SpaceID.ValueString() != "", core.StringPtr(state.SpaceID.ValueString()), nil),
		ProjectID: utils.If(state.SpaceID.ValueString() == "", core.StringPtr(state.ProjectID.ValueString()), nil),
	})
	if utils.Contains(utils.HTTP_OK, response.StatusCode) {
		state.ID = types.StringValue(*model.Metadata.ID)
		diags = resp.State.Set(ctx, &state)
	} else if response.StatusCode == 404 {
		diags = resp.State.Set(ctx, &modelResourceModel{})
	} else {
		resp.Diagnostics.AddError("Error Getting Model", "Could not read Model ID "+state.ID.ValueString()+". Error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *modelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var updateableFields = []string{"Name", "SoftwareSpec"}
	var diags diag.Diagnostics

	mapFieldPatchPath := map[string]string{"Name": "name", "SoftwareSpec": "software_spec"}

	var state modelResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan modelResourceModel
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jsonPatches []watsonmachinelearningv4.JSONPatchOperation
	for _, field := range updateableFields {
		if utils.GetAttr(&plan, field).Interface().(types.String).ValueString() != utils.GetAttr(&state, field).Interface().(types.String).ValueString() {
			switch field {
			case "SoftwareSpec":
				jsonPatches = append(jsonPatches, watsonmachinelearningv4.JSONPatchOperation{
					Op:   core.StringPtr("replace"),
					Path: core.StringPtr(fmt.Sprintf(`/%s`, mapFieldPatchPath[field])),
					Value: &watsonmachinelearningv4.SoftwareSpecRel{
						Name: core.StringPtr(plan.SoftwareSpec.ValueString()),
					},
				})

			default:
				jsonPatches = append(jsonPatches, watsonmachinelearningv4.JSONPatchOperation{
					Op:    core.StringPtr("replace"),
					Path:  core.StringPtr(fmt.Sprintf(`/%s`, mapFieldPatchPath[field])),
					Value: core.StringPtr(utils.GetAttr(&plan, field).Interface().(types.String).ValueString()),
				})
			}
		}
	}

	wmlClient, err := r.client.WMLClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to get WML Client", err.Error())
		return
	}

	if len(jsonPatches) > 0 {
		_, response, err := wmlClient.ModelsUpdate(&watsonmachinelearningv4.ModelsUpdateOptions{
			ModelID:   core.StringPtr(plan.ID.ValueString()),
			SpaceID:   utils.If(state.SpaceID.ValueString() != "", core.StringPtr(state.SpaceID.ValueString()), nil),
			ProjectID: utils.If(state.SpaceID.ValueString() == "", core.StringPtr(state.ProjectID.ValueString()), nil),
			JSONPatch: jsonPatches,
		})

		if err != nil {
			resp.Diagnostics.AddError("Error Updating Model", "Could not update model id "+plan.ID.ValueString()+". Error: "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to Update Model", err.Error())
			return
		}
	}

	if plan.ModelPath.ValueString() != state.ModelPath.ValueString() {
		file, err := os.Open(plan.ModelPath.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Unable to Read Model File", err.Error())
			return
		}
		defer file.Close()
		_, response, err := wmlClient.ModelsUploadContent(&watsonmachinelearningv4.ModelsUploadContentOptions{
			ModelID:       core.StringPtr(plan.ID.ValueString()),
			ContentFormat: core.StringPtr("native"),
			SpaceID:       utils.If(state.SpaceID.ValueString() != "", core.StringPtr(state.SpaceID.ValueString()), nil),
			ProjectID:     utils.If(state.SpaceID.ValueString() == "", core.StringPtr(state.ProjectID.ValueString()), nil),
			Body:          file,
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Uoload Content to Model", err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Error Uoload Content to Model", err.Error())
			return
		}
		if _, err := os.Stat(plan.ModelPath.ValueString()); err == nil || os.IsExist(err) {
			os.Remove(plan.ModelPath.ValueString())
		}
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *modelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state modelResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wmlClient, err := r.client.WMLClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to get WML Client", err.Error())
		return
	}

	response, err := wmlClient.ModelsDelete(&watsonmachinelearningv4.ModelsDeleteOptions{
		ModelID:   core.StringPtr(state.ID.ValueString()),
		SpaceID:   utils.If(state.SpaceID.ValueString() != "", core.StringPtr(state.SpaceID.ValueString()), nil),
		ProjectID: utils.If(state.SpaceID.ValueString() == "", core.StringPtr(state.ProjectID.ValueString()), nil),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Model", "Could not delete model ID "+state.ID.ValueString()+". Error: "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Error Deleting Model", "Could not delete model ID "+state.ID.ValueString()+". Error: "+err.Error())
		return
	}
}

func (r *modelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
