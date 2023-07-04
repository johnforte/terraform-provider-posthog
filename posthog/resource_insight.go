package posthog

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"net/http"
)

var (
	_ resource.Resource              = &insightResource{}
	_ resource.ResourceWithConfigure = &insightResource{}
)

func newInsightResource() resource.Resource {
	return &insightResource{}
}

type insightResource struct {
	client *Client
}

func (r *insightResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ``,
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"short_id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"derived_name": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}
func (r *insightResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Got: %T.", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *insightResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_insight"
}
func (r *insightResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Insight
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	insightJson := InsightRequest{
		Name:        plan.Name.ValueString(),
		DerivedName: plan.DerivedName.ValueString(),
	}
	body, err := json.Marshal(insightJson)
	if err != nil {
		diags.FromErr(err)
		return
	}
	out := r.client.doRequest(http.MethodPost, "insights", body)
	result := convertResponseToInsight(out)
	fmt.Println(body)
	diags = resp.State.Set(ctx, result)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *insightResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Insight
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	response := r.client.doRequest(http.MethodGet, fmt.Sprintf("insights/%d", state.Id), nil)
	diags = resp.State.Set(ctx, response)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *insightResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *insightResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
