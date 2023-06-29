package posthog

var (
	_ resource.Resource              = &projectResource{}
	_ resource.ResourceWithConfigure = &projectResource{}
)

func newProjectResource() resource.Resource {
	return &projectResource{}
}
