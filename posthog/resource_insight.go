package posthog

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

type insightResource struct {
}

var (
	_ resource.Resource              = &insightResource{}
	_ resource.ResourceWithConfigure = &insightResource{}
)

func newInsightResource() resource.Resource {
	return &insightResource{}
}

func (r *insightResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ``,
		Attributes: map[string]schema.Attribute{
			"query": schema.StringAttribute{
				Description:   "",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"derived_name": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
		},
	}
}
