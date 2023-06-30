package posthog

import "github.com/hashicorp/terraform-plugin-framework/types"

type Insight struct {
	Alias        types.String `tfsdk:"alias"`
	ID           types.String `tfsdk:"id"`
	DeploymentID types.String `tfsdk:"deployment_id"`
	TeamID       types.String `tfsdk:"team_id"`
}
