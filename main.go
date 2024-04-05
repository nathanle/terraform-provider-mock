package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/nathanle/terraform-provider-mock/mock"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mock.Provider,
	})
}
