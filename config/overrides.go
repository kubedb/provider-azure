package config

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

var (
	resourceGroup = map[string]string{
		"azurerm_virtual_network_peering": "network",
		"azurerm_virtual_network":         "network",
	}
	resourceKind = map[string]string{
		"azurerm_virtual_network_peering": "VirtualNetworkPeering",
		"azurerm_virtual_network":         "VirtualNetwork",
	}
)

// default api-group & kind configuration for all resources
func groupKindOverride(r *ujconfig.Resource) {
	if _, ok := resourceGroup[r.Name]; ok {
		r.ShortGroup = resourceGroup[r.Name]
	}

	if _, ok := resourceKind[r.Name]; ok {
		r.Kind = resourceKind[r.Name]
	}
}
