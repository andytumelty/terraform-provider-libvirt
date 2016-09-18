package main

import (
	"github.com/laycat/terraform-provider-libvirt/libvirt"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: libvirt.Provider,
	})
}
