package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

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

const NUM_TRIES_DEPLOYMENT = 10
const TIMEOUT_DEPLOYMENT = 5 * time.Second

var (
	_ resource.Resource                = &deploymentResource{}
	_ resource.ResourceWithConfigure   = &deploymentResource{}
	_ resource.ResourceWithImportState = &deploymentResource{}
)

type deploymentResource struct {
	client *client.Client
}

type deploymentResourceModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	ServingUrl types.String `tfsdk:"serving_url"`
	Asset      types.String `tfsdk:"asset"`
	SpaceID    types.String `tfsdk:"space_id"`
	Online     types.Bool   `tfsdk:"online"`
	Batch      types.Bool   `tfsdk:"batch"`
	URL        types.String `tfsdk:"url"`
}

func NewDeploymentResource() resource.Resource {
	return &deploymentResource{}
}

func (r *deploymentResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.Client)
}

func (r *deploymentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_deployment"
}

func (r *deploymentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"serving_url": schema.StringAttribute{
				Optional: true,
			},
			"asset": schema.StringAttribute{
				Optional: true,
			},
			"space_id": schema.StringAttribute{
				Required: true,
			},
			"online": schema.BoolAttribute{
				Optional: true,
			},
			"batch": schema.BoolAttribute{
				Optional: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *deploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan deploymentResourceModel
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

	var servingUrlInterface interface{}
	if plan.ServingUrl.ValueString() != "" {
		servingUrlJson := fmt.Sprintf(`{ "serving_name": "%s" }`, plan.ServingUrl.ValueString())
		err = json.Unmarshal([]byte(servingUrlJson), &servingUrlInterface)
		if err != nil {
			resp.Diagnostics.AddError("Unable to parse Serving URL Json", err.Error())
			return
		}
	}

	deployment, response, err := wmlClient.DeploymentsCreate(&watsonmachinelearningv4.DeploymentsCreateOptions{
		Name:    core.StringPtr(plan.Name.ValueString()),
		SpaceID: core.StringPtr(plan.SpaceID.ValueString()),
		Asset: &watsonmachinelearningv4.Rel{
			ID: core.StringPtr(plan.Asset.ValueString()),
		},
		Online: utils.If(plan.Online.ValueBool(), &watsonmachinelearningv4.DeploymentEntityRequestOnline{
			Parameters: servingUrlInterface,
		}, nil),
		Batch: utils.If(plan.Batch.ValueBool(), &watsonmachinelearningv4.DeploymentEntityRequestBatch{}, nil),
	})

	if err != nil {
		resp.Diagnostics.AddError("Error Deploying Model", "Could not deploy model, unexpected error: "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Deploy Model", err.Error())
		return
	}

	var servingUrl string

	for i := 1; i < NUM_TRIES_DEPLOYMENT; i++ {
		deployment, response, err := wmlClient.DeploymentsGet(&watsonmachinelearningv4.DeploymentsGetOptions{
			DeploymentID: deployment.Metadata.ID,
			SpaceID:      core.StringPtr(plan.SpaceID.ValueString()),
		})
		if err != nil {
			resp.Diagnostics.AddError("Error Getting Deployment", "Could not read Deployment ID "+*deployment.Metadata.ID+": "+err.Error())
			return
		}
		if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
			resp.Diagnostics.AddError("Unable to Get Deployment", err.Error())
			return
		}
		if len(deployment.Entity.Status.ServingUrls) > 0 {
			servingUrl = deployment.Entity.Status.ServingUrls[0]
			break
		}
		time.Sleep(TIMEOUT_DEPLOYMENT)
	}

	if servingUrl == "" {
		resp.Diagnostics.AddError("Error Getting Deployment URL", "Could not get deployment URL for "+*deployment.Metadata.ID+": "+err.Error())
	}

	plan.ID = types.StringValue(*deployment.Metadata.ID)
	plan.URL = types.StringValue(deployment.Entity.Status.ServingUrls[0])

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deploymentResourceModel
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

	deployment, response, err := wmlClient.DeploymentsGet(&watsonmachinelearningv4.DeploymentsGetOptions{
		DeploymentID: core.StringPtr(state.ID.ValueString()),
		SpaceID:      core.StringPtr(state.SpaceID.ValueString()),
	})
	if utils.Contains(utils.HTTP_OK, response.StatusCode) {
		state.ID = types.StringValue(*deployment.Metadata.ID)
		state.URL = types.StringValue(deployment.Entity.Status.ServingUrls[0])
		diags = resp.State.Set(ctx, &state)
	} else if response.StatusCode == 404 {
		diags = resp.State.Set(ctx, &deploymentResourceModel{})
	} else {
		resp.Diagnostics.AddError("Error Getting Deployment", "Could not read Deployment ID "+state.ID.ValueString()+". Error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var updateableFields = []string{"Name", "Asset", "ServingUrl"}
	var diags diag.Diagnostics

	var state deploymentResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan deploymentResourceModel
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jsonPatches []watsonmachinelearningv4.JSONPatchOperation
	for _, field := range updateableFields {
		if utils.GetAttr(&plan, field).Interface().(types.String).ValueString() != utils.GetAttr(&state, field).Interface().(types.String).ValueString() {
			switch field {
			case "Asset":
				jsonPatches = append(jsonPatches, watsonmachinelearningv4.JSONPatchOperation{
					Op:    core.StringPtr("replace"),
					Path:  core.StringPtr(fmt.Sprintf(`/%s`, strings.ToLower(field))),
					Value: &watsonmachinelearningv4.Rel{ID: core.StringPtr(plan.Asset.ValueString())},
				})
			case "ServingUrl":
				if plan.ServingUrl.ValueString() != "" {
					var servingUrlInterface interface{}
					servingUrlJson := fmt.Sprintf(`{ "serving_name": "%s" }`, plan.ServingUrl.ValueString())
					err := json.Unmarshal([]byte(servingUrlJson), &servingUrlInterface)
					if err != nil {
						resp.Diagnostics.AddError("Unable to parse Serving URL Json", err.Error())
						return
					}
					jsonPatches = append(jsonPatches, watsonmachinelearningv4.JSONPatchOperation{
						Op:    core.StringPtr("replace"),
						Path:  core.StringPtr("/online/parameters"),
						Value: servingUrlInterface,
					})
				}
			default:
				jsonPatches = append(jsonPatches, watsonmachinelearningv4.JSONPatchOperation{
					Op:    core.StringPtr("replace"),
					Path:  core.StringPtr(fmt.Sprintf(`/%s`, strings.ToLower(field))),
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

	deployment, response, err := wmlClient.DeploymentsUpdate(&watsonmachinelearningv4.DeploymentsUpdateOptions{
		DeploymentID: core.StringPtr(plan.ID.ValueString()),
		SpaceID:      core.StringPtr(plan.SpaceID.ValueString()),
		JSONPatch:    jsonPatches,
	})
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Deployment", "Could not update deployment ID "+plan.ID.ValueString()+": "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Update Deployment", err.Error())
		return
	}

	plan.ID = types.StringValue(*deployment.Metadata.ID)
	plan.URL = types.StringValue(deployment.Entity.Status.ServingUrls[0])

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state deploymentResourceModel
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

	response, err := wmlClient.DeploymentsDelete(&watsonmachinelearningv4.DeploymentsDeleteOptions{
		DeploymentID: core.StringPtr(state.ID.ValueString()),
		SpaceID:      core.StringPtr(state.SpaceID.ValueString()),
	})

	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Deployment", "Could not delete deployment ID "+state.ID.ValueString()+": "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Delete Model", err.Error())
		return
	}
}

func (r *deploymentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
