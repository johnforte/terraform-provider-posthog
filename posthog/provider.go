package posthog

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type posthogProvider struct{}

func New() provider.Provider {
	return &posthogProvider{}
}

func (p *posthogProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	c := newClient(config.PersonalAPIToken.ValueString(), config.ProjectId.ValueString())

	resp.DataSourceData = c
	resp.ResourceData = c
}

func (p *posthogProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "posthog"
}

func (p *posthogProvider) Schema(_ context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `PostHog Provider for Terraform`,
		Attributes: map[string]schema.Attribute{
			"personal_api_token": schema.StringAttribute{
				Optional:    true,
				Description: "This is your personal api token.",
				Sensitive:   true,
			},
			"project_id": schema.StringAttribute{
				Optional:    true,
				Description: "This is your project id to run terraform against.",
			},
		},
	}
}

func (p *posthogProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newInsightResource,
		newProjectResource,
	}
}

// No Data Sources currently just adding this as stub
func (p *posthogProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

type providerData struct {
	PersonalAPIToken types.String `tfsdk:"personal_api_token"`
	ProjectId        types.String `tfsdk:"project_id"`
}
