/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	_ "embed"
	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"kubedb.dev/provider-azure/config/authorization"
	"kubedb.dev/provider-azure/config/base"
	"kubedb.dev/provider-azure/config/cache"
	"kubedb.dev/provider-azure/config/common"
	"kubedb.dev/provider-azure/config/cosmosdb"
	"kubedb.dev/provider-azure/config/dbformariadb"
	"kubedb.dev/provider-azure/config/dbformysql"
	"kubedb.dev/provider-azure/config/dbforpostgresql"
	"kubedb.dev/provider-azure/config/keyvault"
	"kubedb.dev/provider-azure/config/network"
	"kubedb.dev/provider-azure/config/sql"
	"kubedb.dev/provider-azure/config/storage"
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
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	// "azure" group contains resources that actually do not have a specific
	// group, e.g. ResourceGroup with APIVersion "azure.upbound.io/v1beta1".
	// We need to include the controllers for this group into the base packages
	// list to get their controllers packaged together with the config package
	// controllers (provider family config package).
	//for _, c := range []string{"internal/controller/azure/resourcegroup", "internal/controller/azure/providerregistration", "internal/controller/azure/subscription"} {
	//	pc.BasePackages.ControllerMap[c] = "config"
	//}

	// API group overrides from Terraform import statements
	for _, r := range pc.Resources {
		groupKindOverride(r)
	}

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		base.Configure,
		network.Configure,
		cache.Configure,
		cosmosdb.Configure,
		dbformariadb.Configure,
		dbformysql.Configure,
		dbforpostgresql.Configure,
		keyvault.Configure,
		sql.Configure,
		storage.Configure,
		authorization.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()

	// This function runs after the special configurations were applied. However, if some references were configured in
	// the configuration files, the reference generator does not override them.
	for _, r := range pc.Resources {
		delete(r.References, "resource_group_name")
		if err := common.AddCommonReferences(r); err != nil {
			panic(err)
		}
	}
	return pc
}
