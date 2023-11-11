/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"kubedb.dev/provider-azure/config/cache"
	"kubedb.dev/provider-azure/config/cosmosdb"
	"kubedb.dev/provider-azure/config/dbformariadb"
	"kubedb.dev/provider-azure/config/dbformysql"
	"kubedb.dev/provider-azure/config/dbforpostgresql"
	"kubedb.dev/provider-azure/config/keyvault"
	"kubedb.dev/provider-azure/config/network"

	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "azure"
	modulePath     = "kubedb.dev/provider-azure"
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
		ujconfig.WithRootGroup("azure.kubedb.com"),
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
		cache.Configure,
		cosmosdb.Configure,
		dbformariadb.Configure,
		dbformysql.Configure,
		dbforpostgresql.Configure,
		keyvault.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc
}
