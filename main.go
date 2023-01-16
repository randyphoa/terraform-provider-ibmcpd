package main

import (
	"context"

	"terraform-provider-ibmcpd/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name ibmcpd

func main() {
	providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/randyphoa/terraform-provider-ibmcpd",
	})
}
