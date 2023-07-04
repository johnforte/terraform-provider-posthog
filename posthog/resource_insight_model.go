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
	Filter      types.Map    `tfsdk:"filter"`
}
type InsightRequest struct {
	Name        string            `json:"name"`
	DerivedName string            `json:"derived_name"`
	Filter      map[string]string `json:"filter"`
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
	return Insight{
		Id:          types.Int64Value(insight.Id),
		ShortId:     types.StringValue(insight.ShortId),
		Name:        types.StringValue(insight.Name),
		DerivedName: types.StringValue(insight.DerivedName),
		Filter:      types.MapValue(types.StringType, insight.Filter),
	}
}
