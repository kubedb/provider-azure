/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
	"strings"
)

func getResourceNameFromIDURLFn(pos int) config.GetExternalNameFn {
	return func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("cannot get id")
		}
		words := strings.Split(id.(string), "/")
		return words[len(words)-pos], nil
	}
}

func keyVaultURLIDConf(resourceType string) config.ExternalName {
	e := config.NameAsIdentifier
	e.GetExternalNameFn = getResourceNameFromIDURLFn(2)
	e.GetIDFn = func(_ context.Context, _ string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		keyVaultID, ok := parameters["key_vault_id"]
		if !ok {
			return "", errors.New("cannot get key_vault_id")
		}
		words := strings.Split(keyVaultID.(string), "/")
		keyVaultName := words[len(words)-1]

		name, ok := parameters["name"]
		if !ok {
			return "", errors.New("cannot get name")
		}

		version, ok := parameters["version"]
		if !ok {
			return "", nil
		}

		return fmt.Sprintf("https://%s.vault.azure.net/%s/%s/%s",
			keyVaultName, resourceType, name, version), nil
	}
	return e
}

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{

	// base group
	"azurerm_subscription":                   config.TemplatedStringAsIdentifier("alias", "/providers/Microsoft.Subscription/aliases/{{ .external_name }}"),
	"azurerm_resource_provider_registration": config.IdentifierFromProvider,
	"azurerm_resource_group":                 config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .external_name }}"),

	// network
	"azurerm_virtual_network_peering":                   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworks/{{ .parameters.virtual_network_name }}/virtualNetworkPeerings/{{ .external_name }}"),
	"azurerm_virtual_network":                           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworks/{{ .external_name }}"),
	"azurerm_private_dns_zone":                          config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .external_name }}"),
	"azurerm_private_dns_zone_virtual_network_link":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/privateDnsZones/{{ .parameters.private_dns_zone_name }}/virtualNetworkLinks/{{ .external_name }}"),
	"azurerm_network_security_group":                    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/networkSecurityGroups/{{ .external_name }}"),
	"azurerm_subnet_network_security_group_association": config.IdentifierFromProvider,
	"azurerm_subnet_route_table_association":            config.IdentifierFromProvider,

	// subnet
	"azurerm_subnet":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/virtualNetworks/{{ .parameters.virtual_network_name }}/subnets/{{ .external_name }}"),
	"azurerm_route_table": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Network/routeTables/{{ .external_name }}"),

	// redis
	"azurerm_redis_cache":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redis/{{ .external_name }}"),
	"azurerm_redis_firewall_rule":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redis/{{ .parameters.redis_cache_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_redis_enterprise_cluster":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redisEnterprise/{{ .external_name }}"),
	"azurerm_redis_enterprise_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cluster_id }}/databases/{{ .external_name }}"),
	"azurerm_redis_linked_server":       config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Cache/redis/{{ .parameters.target_redis_cache_name }}/linkedServers/{{ .parameters.linked_redis_cache_name }}"),

	// cosmosdb
	"azurerm_cosmosdb_sql_database":  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .external_name }}"),
	"azurerm_cosmosdb_sql_container": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .parameters.database_name }}/containers/{{ .external_name }}"),
	// We switched to IdentifierFromProvider configuration because of the problem in this issue: https://github.com/crossplane/upjet/issues/32
	"azurerm_cosmosdb_sql_role_assignment": config.IdentifierFromProvider,
	// We switched to IdentifierFromProvider configuration because of the problem in this issue: https://github.com/crossplane/upjet/issues/32
	"azurerm_cosmosdb_sql_role_definition":  config.IdentifierFromProvider,
	"azurerm_cosmosdb_mongo_database":       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/mongodbDatabases/{{ .external_name }}"),
	"azurerm_cosmosdb_mongo_collection":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/mongodbDatabases/{{ .parameters.database_name }}/collections/{{ .external_name }}"),
	"azurerm_cosmosdb_cassandra_keyspace":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/cassandraKeyspaces/{{ .external_name }}"),
	"azurerm_cosmosdb_cassandra_table":      config.TemplatedStringAsIdentifier("name", "{{ .parameters.cassandra_keyspace_id }}/tables/{{ .external_name }}"),
	"azurerm_cosmosdb_gremlin_database":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/gremlinDatabases/{{ .external_name }}"),
	"azurerm_cosmosdb_gremlin_graph":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/gremlinDatabases/{{ .parameters.database_name }}/graphs/{{ .external_name }}"),
	"azurerm_cosmosdb_sql_stored_procedure": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/sqlDatabases/{{ .parameters.database_name }}/containers/{{ .parameters.container_name }}/storedProcedures/{{ .external_name }}"),
	// Contains only container_id as parameter which includes other information like resource group name etc.
	"azurerm_cosmosdb_sql_function": config.TemplatedStringAsIdentifier("name", "{{ .parameters.container_id }}/userDefinedFunctions/{{ .external_name }}"),
	"azurerm_cosmosdb_table":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/tables/{{ .external_name }}"),
	"azurerm_cosmosdb_account":      config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .external_name }}"),
	// NOTE: "azurerm_cosmosdb_notebook_workspace" resource removed because this is not available in Azure portal.
	// Please see the note https://learn.microsoft.com/en-us/azure/cosmos-db/sql/enable-notebooks
	// "azurerm_cosmosdb_notebook_workspace":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/databaseAccounts/{{ .parameters.account_name }}/notebookWorkspaces/{{ .external_name }}"),
	"azurerm_cosmosdb_sql_trigger":          config.TemplatedStringAsIdentifier("name", "{{ .parameters.container_id }}/triggers/{{ .external_name }}"),
	"azurerm_cosmosdb_cassandra_cluster":    config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DocumentDB/cassandraClusters/{{ .external_name }}"),
	"azurerm_cosmosdb_cassandra_datacenter": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cassandra_cluster_id }}/dataCenters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.DocumentDB/databaseAccounts/account1/services/SqlDedicatedGateway
	"azurerm_cosmosdb_sql_dedicated_gateway": config.TemplatedStringAsIdentifier("", "{{ .parameters.cosmosdb_account_id }}/services/SqlDedicatedGateway"),

	// mariadb
	"azurerm_mariadb_server":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .external_name }}"),
	"azurerm_mariadb_database":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	"azurerm_mariadb_firewall_rule":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_mariadb_virtual_network_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMariaDB/servers/{{ .parameters.server_name }}/virtualNetworkRules/{{ .external_name }}"),
	"azurerm_mariadb_configuration":        config.IdentifierFromProvider,

	// dbformysql
	"azurerm_mysql_server":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .external_name }}"),
	"azurerm_mysql_configuration":        config.IdentifierFromProvider,
	"azurerm_mysql_firewall_rule":        config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_mysql_virtual_network_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .parameters.server_name }}/virtualNetworkRules/{{ .external_name }}"),
	// A MySQL Active Directory Administrator can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforMySQL/servers/myserver/administrators/activeDirectory
	"azurerm_mysql_active_directory_administrator": config.IdentifierFromProvider,
	// MySQL Database's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/servers/server1/databases/database1
	"azurerm_mysql_database": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	// dbformysql flexible variants
	"azurerm_mysql_flexible_server":               config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .external_name }}"),
	"azurerm_mysql_flexible_database":             config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	"azurerm_mysql_flexible_server_configuration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .parameters.server_name }}/configurations/{{ .external_name }}"),
	"azurerm_mysql_flexible_server_firewall_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/flexibleServers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),

	// postgresql
	"azurerm_postgresql_server":                         config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .external_name }}"),
	"azurerm_postgresql_flexible_server_configuration":  config.IdentifierFromProvider,
	"azurerm_postgresql_database":                       config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	"azurerm_postgresql_active_directory_administrator": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}"),
	"azurerm_postgresql_flexible_server_database":       config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/databases/{{ .external_name }}"),
	"azurerm_postgresql_flexible_server_firewall_rule":  config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/firewallRules/{{ .external_name }}"),
	"azurerm_postgresql_firewall_rule":                  config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),
	"azurerm_postgresql_flexible_server":                config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{ .external_name }}"),
	"azurerm_postgresql_virtual_network_rule":           config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforPostgreSQL/servers/{{ .parameters.server_name }}/virtualNetworkRules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforPostgreSQL/servers/server1/keys/keyvaultname_key-name_keyversion
	"azurerm_postgresql_server_key": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforPostgreSQL/servers/server1/configurations/backslash_quote
	"azurerm_postgresql_configuration": config.IdentifierFromProvider,

	// keyvault
	"azurerm_key_vault":     config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.KeyVault/vaults/{{ .external_name }}"),
	"azurerm_key_vault_key": keyVaultURLIDConf("keys"),

	// storage
	"azurerm_storage_account":   config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Storage/storageAccounts/{{ .external_name }}"),
	"azurerm_storage_container": config.TemplatedStringAsIdentifier("name", "https://{{ .parameters.storage_account_name }}.blob.core.windows.net/{{ .external_name }}"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
