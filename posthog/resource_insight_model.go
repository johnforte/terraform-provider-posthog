package posthog

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"io"
)

type Insight struct {
	Id          types.Int64  `tfsdk:"id"`
	ShortId     types.String `tfsdk:"short_id"`
	Name        types.String `tfsdk:"name"`
	DerivedName types.String `tfsdk:"derived_name"`
	Filter      types.Map    `tfsdk:"filter"`
}
type InsightRequest struct {
	Name        string                `json:"name"`
	DerivedName string                `json:"derived_name"`
	Filter      map[string]attr.Value `json:"filter"`
}
type InsightResponse struct {
	Id          int64             `json:"id"`
	ShortId     string            `json:"short_id"`
	Name        string            `json:"name"`
	DerivedName string            `json:"derived_name"`
	Filter      map[string]string `json:"filter"`
}

func convertResponseToInsight(response io.Reader) Insight {
	var insight InsightResponse
	decoder := json.NewDecoder(response)
	err := decoder.Decode(&insight)
	if err != nil {
		return Insight{}
	}
	filters := map[string]attr.Value{}
	for key, val := range insight.Filter {
		filters[key] = types.StringValue(fmt.Sprintf("%v", val))
	}
	value, _ := types.MapValue(types.StringType, filters)
	return Insight{
		Id:          types.Int64Value(insight.Id),
		ShortId:     types.StringValue(insight.ShortId),
		Name:        types.StringValue(insight.Name),
		DerivedName: types.StringValue(insight.DerivedName),
		Filter:      value,
	}
}
