package posthog

import "github.com/hashicorp/terraform-plugin-framework/types"

type Insight struct {
	Id          types.Int64  `tfsdk:"id"`
	ShortId     types.String `tfsdk:"short_id"`
	Name        types.String `tfsdk:"name"`
	DerivedName types.String `tfsdk:"derived_name"`
}
