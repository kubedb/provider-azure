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

	p.AddResourceConfigurator("azurerm_virtual_network_gateway", func(r *config.Resource) {
		r.Kind = "VirtualNetworkGateway"
		r.References["ip_configuration.subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
	p.AddResourceConfigurator("azurerm_private_dns_zone", func(r *config.Resource) {
		r.UseAsync = false
	})

	p.AddResourceConfigurator("azurerm_network_security_rule", func(r *config.Resource) {
		r.UseAsync = false
		r.References["network_security_group_name"] = config.Reference{
			Type: "SecurityGroup",
		}
	})

	p.AddResourceConfigurator("azurerm_network_interface_security_group_association", func(r *config.Resource) {
		r.Kind = "NetworkInterfaceSecurityGroupAssociation"
		r.References["network_interface_id"] = config.Reference{
			Type:      "NetworkInterface",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["network_security_group_id"] = config.Reference{
			Type:      "SecurityGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_subnet_network_security_group_association", func(r *config.Resource) {
		r.Kind = "SubnetNetworkSecurityGroupAssociation"
		r.References["subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["network_security_group_id"] = config.Reference{
			Type:      "SecurityGroup",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_subnet_route_table_association", func(r *config.Resource) {
		r.Kind = "SubnetRouteTableAssociation"
		r.References["subnet_id"] = config.Reference{
			Type:      "Subnet",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
		r.References["route_table_id"] = config.Reference{
			Type:      "RouteTable",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("azurerm_network_security_group", func(r *config.Resource) {
		r.UseAsync = false
	})

}
