package network

import (
	"github.com/crossplane/upjet/pkg/config"
	"kubedb.dev/provider-azure/apis/rconfig"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_virtual_network_peering", func(r *config.Resource) {
		r.Kind = "VirtualNetworkPeering"
		r.References["virtual_network_name"] = config.Reference{
			Type: "VirtualNetwork",
		}
		r.References["remote_virtual_network_id"] = config.Reference{
			Type:      "VirtualNetwork",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_virtual_network", func(r *config.Resource) {
		r.Kind = "VirtualNetwork"
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"subnet"},
		}
		config.MoveToStatus(r.TerraformResource, "subnet")
	})

	p.AddResourceConfigurator("azurerm_private_dns_zone", func(r *config.Resource) {
		r.UseAsync = false
	})
}
