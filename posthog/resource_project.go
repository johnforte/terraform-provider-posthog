package posthog

import (
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type projectResource struct {
}

var (
	_ resource.Resource              = &projectResource{}
	_ resource.ResourceWithConfigure = &projectResource{}
)

func newProjectResource() resource.Resource {
	return &projectResource{}
}
