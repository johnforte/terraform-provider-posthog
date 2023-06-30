package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/johnforte/terraform-provider-posthog/posthog"
)

func main() {
	err := providerserver.Serve(context.Background(), posthog.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/johnforte/posthog",
	})
	if err != nil {
		log.Fatalf("unable to serve provider: %s", err)
	}
}
