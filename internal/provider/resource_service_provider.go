package provider

import (
	"context"
	"strings"

	"terraform-provider-ibmcpd/internal/go-sdk/client"
	"terraform-provider-ibmcpd/internal/go-sdk/watsonopenscalev2"
	"terraform-provider-ibmcpd/internal/utils"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &serviceProviderResource{}
	_ resource.ResourceWithConfigure   = &serviceProviderResource{}
	_ resource.ResourceWithImportState = &deploymentResource{}
)

type serviceProviderResource struct {
	client *client.Client
}

type serviceProviderResourceModel struct {
	ID                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	ServiceType        types.String `tfsdk:"service_type"`
	OperationalSpaceID types.String `tfsdk:"operational_space_id"`

	DeploymentSpaceID types.String `tfsdk:"deployment_space_id"`
	UserName          types.String `tfsdk:"user_name"`
	APIKey            types.String `tfsdk:"api_key"`
	Password          types.String `tfsdk:"password"`
	Url               types.String `tfsdk:"url"`

	AWSAccessKeyID     types.String `tfsdk:"aws_access_key_id"`
	AWSSecretAccessKey types.String `tfsdk:"aws_secret_access_key"`
	AWSRegion          types.String `tfsdk:"aws_region"`
}

func NewServiceProviderResource() resource.Resource {
	return &serviceProviderResource{}
}

func (r *serviceProviderResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*client.Client)
}

func (r *serviceProviderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_provider"
}

func (r *serviceProviderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"service_type": schema.StringAttribute{
				Required: true,
			},
			"operational_space_id": schema.StringAttribute{
				Required: true,
			},
			"deployment_space_id": schema.StringAttribute{
				Optional: true,
			},

			"user_name": schema.StringAttribute{
				Optional: true,
			},
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"url": schema.StringAttribute{
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

func (r *serviceProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan serviceProviderResourceModel
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

	var credentials watsonopenscalev2.MLCredentialsIntf

	if strings.Contains(plan.Url.ValueString(), "cloud.ibm.com") {
		credentials = &watsonopenscalev2.WMLCredentialsCP4D{
			Password: core.StringPtr(plan.Password.ValueString()),
			Token:    core.StringPtr(plan.APIKey.ValueString()),
			Username: core.StringPtr(plan.UserName.ValueString()),
			URL:      core.StringPtr(plan.Url.ValueString()),
		}
	}

	if !strings.Contains(plan.Url.ValueString(), "cloud.ibm.com") {
		credentials = &watsonopenscalev2.WMLCredentialsCloud{
			Apikey:     core.StringPtr(plan.APIKey.ValueString()),
			URL:        core.StringPtr(plan.Url.ValueString()),
			InstanceID: core.StringPtr(""),
		}
	}

	if plan.AWSSecretAccessKey.ValueString() != "" {
		credentials = &watsonopenscalev2.SageMakerCredentials{
			AccessKeyID:     core.StringPtr(plan.AWSAccessKeyID.ValueString()),
			SecretAccessKey: core.StringPtr(plan.AWSSecretAccessKey.ValueString()),
			Region:          core.StringPtr(plan.AWSRegion.ValueString()),
		}
	}

	result, response, err := wosClient.ServiceProvidersAdd(&watsonopenscalev2.ServiceProvidersAddOptions{
		Credentials:        credentials,
		Name:               core.StringPtr(plan.Name.ValueString()),
		ServiceType:        core.StringPtr(plan.ServiceType.ValueString()),
		OperationalSpaceID: core.StringPtr(plan.OperationalSpaceID.ValueString()),
		DeploymentSpaceID:  core.StringPtr(plan.DeploymentSpaceID.ValueString()),
	})

	if err != nil {
		resp.Diagnostics.AddError("Error Creating Service Provider", "Could not create service provider, unexpected error: "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Create Service Provider", err.Error())
		return
	}

	plan.ID = types.StringValue(*result.Metadata.ID)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state serviceProviderResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to get WOS Client", err.Error())
		return
	}

	serviceProvider, response, err := wosClient.ServiceProvidersGet(&watsonopenscalev2.ServiceProvidersGetOptions{
		ServiceProviderID: core.StringPtr(state.ID.ValueString()),
	})
	if utils.Contains(utils.HTTP_OK, response.StatusCode) {
		state.ID = types.StringValue(*serviceProvider.Metadata.ID)
		state.Name = types.StringValue(*serviceProvider.Entity.Name)
		state.ServiceType = types.StringValue(*serviceProvider.Entity.ServiceType)
		state.OperationalSpaceID = types.StringValue(*serviceProvider.Entity.OperationalSpaceID)
		diags = resp.State.Set(ctx, &state)
	} else if response.StatusCode == 404 {
		diags = resp.State.Set(ctx, &serviceProviderResourceModel{})
	} else {
		resp.Diagnostics.AddError("Error Getting Service Provider", "Could not read Service Provider ID "+state.ID.ValueString()+": "+err.Error())
		return
	}

	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *serviceProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *serviceProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state serviceProviderResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	wosClient, err := r.client.WOSClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to get WOS Client", err.Error())
		return
	}

	response, err := wosClient.ServiceProvidersDelete(&watsonopenscalev2.ServiceProvidersDeleteOptions{
		ServiceProviderID: core.StringPtr(state.ID.ValueString()),
	})

	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Service Provider", "Could not delete Service Provider ID "+state.ID.ValueString()+": "+err.Error())
		return
	}
	if !utils.Contains(utils.HTTP_OK, response.StatusCode) {
		resp.Diagnostics.AddError("Unable to Delete Service Provider", err.Error())
		return
	}
}

func (r *serviceProviderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
