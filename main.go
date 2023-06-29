package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/johnforte/terraform-provider-posthog/posthog"
)

func main() {
	err := providerserver.Serve(context.Background(), posthog.New)
	if err != nil {
		log.Fatalf("unable to serve provider: %s", err)
	}
}
