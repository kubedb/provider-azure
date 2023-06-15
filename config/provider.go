/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"kubeform.dev/provider-azure/config/network"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "azure"
	modulePath     = "kubeform.dev/provider-azure"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithRootGroup("azure.kubeform.com"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	// API group overrides from Terraform import statements
	for _, r := range pc.Resources {
		groupKindOverride(r)
	}

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		network.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc
}
