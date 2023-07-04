package posthog

import (
	"encoding/json"
	"fmt"
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

type InsightResponse struct {
	id           int64             `json:"id"`
	short_id     string            `json:"short_id"`
	name         string            `json:"name"`
	derived_name string            `json:"derived_name"`
	filter       map[string]string `json:"filter"`
}

func convertResponseToInsight(response io.Reader) Insight {
	var profile InsightResponse
	decoder := json.NewDecoder(response)
	err := decoder.Decode(&profile)
	if err != nil {
		return Insight{}
	}
	fmt.Println(profile)
	return Insight{
		Id:          types.Int64Value(profile.id),
		ShortId:     types.StringValue(profile.short_id),
		Name:        types.StringValue(profile.name),
		DerivedName: types.StringValue(profile.derived_name),
		Filter:      types.MapValue(profile.filter),
	}
}
