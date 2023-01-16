package provider

import (
	"context"

	"terraform-provider-ibmcpd/internal/go-sdk/client"
	"terraform-provider-ibmcpd/internal/utils"

	"github.com/hashicorp/terraform-plugin-framework-validators/providervalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ provider.Provider = &Provider{}
)

type Provider struct {
}

type ProviderModel struct {
	URL      types.String `tfsdk:"url"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	ApiKey   types.String `tfsdk:"api_key"`
}

func New() provider.Provider {
	return &Provider{}
}

func (p *Provider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ibmcpd"
}

func (p *Provider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provider to manage IBM Cloud Pak for Data resources.",
		Attributes: map[string]schema.Attribute{
			"url": schema.StringAttribute{
				Description: "URL for IBM Cloud Pak for Data.",
				Required:    true,
			},
			"username": schema.StringAttribute{
				Description: "Username for IBM Cloud Pak for Data.",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password for IBM Cloud Pak for Data.",
				Optional:    true,
				Sensitive:   true,
			},
			"api_key": schema.StringAttribute{
				Description: "API key for IBM Cloud Pak for Data.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

func (p *Provider) ConfigValidators(ctx context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{
		providervalidator.AtLeastOneOf(
			path.MatchRoot("password"),
			path.MatchRoot("api_key"),
		),
	}
}

func (p *Provider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config ProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if config.URL.IsUnknown() {
		resp.Diagnostics.AddAttributeError(path.Root("url"), "Unknown URL", "Unable to create API client with unknown URL: "+config.URL.ValueString())
	}

	if resp.Diagnostics.HasError() {
		return
	}

	url := config.URL.ValueString()
	username := config.Username.ValueString()
	password := config.Password.ValueString()
	apiKey := config.ApiKey.ValueString()

	if url == "" {
		resp.Diagnostics.AddAttributeError(path.Root("url"), "Missing url", "Unable to create API client with missing URL.")
	}

	if password == "" && apiKey == "" {
		resp.Diagnostics.AddAttributeError(path.Root("password"), "Missing password", "Unable to create API client with missing password or API key.")
		resp.Diagnostics.AddAttributeError(path.Root("api_key"), "Missing API key", "Unable to create API client with missing password or API key.")
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "ibmcpd_url", url)
	ctx = tflog.SetField(ctx, "ibmcpd_username", username)
	ctx = tflog.SetField(ctx, "ibmcpd_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "ibmcpd_password")

	tflog.Debug(ctx, "Creating IBM CPD client")

	// auth, err := core.NewCloudPakForDataAuthenticatorUsingAPIKey(fmt.Sprintf("%s/icp4d-api", url), username, apiKey, true, map[string]string{})
	// auth, err := core.NewIamAuthenticator(ApiKey, IamUrl, "bx", "bx", false, map[string]string{})

	auth, err := utils.GetAuthenticator(url, username, password, apiKey)
	if err != nil {
		resp.Diagnostics.AddError("Unable to authenticate IBM CPD credentials", "Error: "+err.Error())
		return
	}
	client, err := client.NewClient(auth, &client.Config{URL: url})
	if err != nil {
		resp.Diagnostics.AddError("Unable to create Client API", "Error: "+err.Error())
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured IBM CPD client", map[string]interface{}{"success": true})
}

func (p *Provider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewModelResource,
		NewDeploymentResource,
		NewSubscriptionResource,
		NewMonitorInstanceResource,
		NewRecordResource,
		NewServiceProviderResource,
	}
}

func (p *Provider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewSPAssetDataSource,
	}
}
