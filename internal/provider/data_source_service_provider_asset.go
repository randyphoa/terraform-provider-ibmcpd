package provider

import (
	"context"
	"encoding/json"
	"io"

	"terraform-provider-ibmcpd/internal/go-sdk/client"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &spAssetDataSource{}
	_ datasource.DataSourceWithConfigure = &spAssetDataSource{}
)

func NewSPAssetDataSource() datasource.DataSource {
	return &spAssetDataSource{}
}

type spAssetDataSource struct {
	client *client.Client
}

type spAssetDataSourceModel struct {
	ID                types.String    `tfsdk:"id"`
	ServiceProviderID types.String    `tfsdk:"service_provider_id"`
	SPAssets          []spAssetsModel `tfsdk:"sp_assets"`
}

type spAssetsModel struct {
	AssetID      types.String `tfsdk:"asset_id"`
	Name         types.String `tfsdk:"name"`
	DeploymentRN types.String `tfsdk:"deployment_rn"`
	DeploymentID types.String `tfsdk:"deployment_id"`
}

func (d *spAssetDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = req.ProviderData.(*client.Client)
}

func (d *spAssetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sp_assets"
}

func (d *spAssetDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists assets that are connected to the service provider.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"service_provider_id": schema.StringAttribute{
				Description: "Identifer for service provider.",
				Required:    true,
			},
			"sp_assets": schema.ListNestedAttribute{
				Description: "List of assets.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"asset_id": schema.StringAttribute{
							Description: "Identifer for asset.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name of asset.",
							Computed:    true,
						},
						"deployment_rn": schema.StringAttribute{
							Description: "Deployment endpoint of asset.",
							Computed:    true,
						},
						"deployment_id": schema.StringAttribute{
							Description: "Deployment identifer of asset.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *spAssetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var plan spAssetDataSourceModel
	diags := req.Config.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state spAssetDataSourceModel

	wosClient, err := d.client.WOSClient(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to get WOS Client", err.Error())
		return
	}

	pathParamsMap := map[string]string{
		"service_provider_id": plan.ServiceProviderID.ValueString(),
	}
	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(context.Background())
	_, err = builder.ResolveRequestURL(d.client.Config.URL, `/v1/ml_instances/{service_provider_id}/deployments?datamart_id=00000000-0000-0000-0000-000000000000&limit=10`, pathParamsMap)
	if err != nil {
		resp.Diagnostics.AddError("Unable to build SP Assets URL", err.Error())
		return
	}
	builder.AddHeader("Content-Type", "application/json")
	request, err := builder.Build()
	if err != nil {
		resp.Diagnostics.AddError("Unable to create URL builder", err.Error())
		return
	}
	err = wosClient.Service.Options.Authenticator.Authenticate(request)
	if err != nil {
		resp.Diagnostics.AddError("Authentication error", err.Error())
		return
	}
	response, err := wosClient.Service.Client.Do(request)
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

	resources := jsonResponse["resources"].([]interface{})
	for _, resource := range resources {
		metadata := resource.(map[string]interface{})["metadata"].(map[string]interface{})
		entity := resource.(map[string]interface{})["entity"].(map[string]interface{})
		asset := entity["asset"].(map[string]interface{})

		spAsset := spAssetsModel{
			AssetID:      types.StringValue(asset["asset_id"].(string)),
			Name:         types.StringValue(entity["name"].(string)),
			DeploymentRN: types.StringValue(entity["deployment_rn"].(string)),
			DeploymentID: types.StringValue(metadata["guid"].(string)),
		}
		state.SPAssets = append(state.SPAssets, spAsset)
	}

	state.ID = types.StringValue("placeholder")

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
