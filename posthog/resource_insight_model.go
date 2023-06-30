package posthog

import "github.com/hashicorp/terraform-plugin-framework/types"

type Insight struct {
	Name        types.String `tfsdk:"name"`
	DerivedName types.String `tfsdk:"derived_name"`
}
