package posthog

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"io"
)

type Insight struct {
	Id          types.Int64  `tfsdk:"id"`
	ShortId     types.String `tfsdk:"short_id"`
	Name        types.String `tfsdk:"name"`
	DerivedName types.String `tfsdk:"derived_name"`
}

func convertResponseToInsight(response io.Reader) Insight {
	var target Insight
	json.NewDecoder(response).Decode(&target)
	return target
}
