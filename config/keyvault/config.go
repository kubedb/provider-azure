package keyvault

import (
	"github.com/upbound/upjet/pkg/config"
	"kubeform.dev/provider-azure/apis/rconfig"
)

// Configure configures keyvault group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azurerm_key_vault", func(r *config.Resource) {
		// Mutually exclusive with azurerm_key_vault_access_policy
		config.MoveToStatus(r.TerraformResource, "access_policy")
	})

	p.AddResourceConfigurator("azurerm_key_vault_key", func(r *config.Resource) {
		r.References["key_vault_id"] = config.Reference{
			Type:      "Vault",
			Extractor: rconfig.ExtractResourceIDFuncPath,
		}
	})
}
