package config

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

var (
	resourceGroup = map[string]string{
		"azurerm_virtual_network_peering":                   "network",
		"azurerm_virtual_network":                           "network",
		"azurerm_private_dns_zone":                          "network",
		"azurerm_subnet":                                    "network",
		"azurerm_private_dns_zone_virtual_network_link":     "network",
		"azurerm_subnet_network_security_group_association": "network",
		"azurerm_subnet_route_table_association":            "network",
		"azurerm_route_table":                               "network",

		"azurerm_redis_cache":                               "cache",
		"azurerm_redis_firewall_rule":                       "cache",
		"azurerm_redis_enterprise_cluster":                  "cache",
		"azurerm_redis_enterprise_database":                 "cache",
		"azurerm_redis_linked_server":                       "cache",
		"azurerm_cosmosdb_sql_database":                     "cosmosdb",
		"azurerm_cosmosdb_sql_container":                    "cosmosdb",
		"azurerm_cosmosdb_sql_role_assignment":              "cosmosdb",
		"azurerm_cosmosdb_sql_role_definition":              "cosmosdb",
		"azurerm_cosmosdb_mongo_database":                   "cosmosdb",
		"azurerm_cosmosdb_mongo_collection":                 "cosmosdb",
		"azurerm_cosmosdb_cassandra_keyspace":               "cosmosdb",
		"azurerm_cosmosdb_cassandra_table":                  "cosmosdb",
		"azurerm_cosmosdb_gremlin_database":                 "cosmosdb",
		"azurerm_cosmosdb_gremlin_graph":                    "cosmosdb",
		"azurerm_cosmosdb_sql_stored_procedure":             "cosmosdb",
		"azurerm_cosmosdb_sql_function":                     "cosmosdb",
		"azurerm_cosmosdb_table":                            "cosmosdb",
		"azurerm_cosmosdb_account":                          "cosmosdb",
		"azurerm_cosmosdb_sql_trigger":                      "cosmosdb",
		"azurerm_cosmosdb_cassandra_cluster":                "cosmosdb",
		"azurerm_cosmosdb_cassandra_datacenter":             "cosmosdb",
		"azurerm_cosmosdb_sql_dedicated_gateway":            "cosmosdb",
		"azurerm_mariadb_server":                            "dbformariadb",
		"azurerm_mariadb_database":                          "dbformariadb",
		"azurerm_mariadb_firewall_rule":                     "dbformariadb",
		"azurerm_mariadb_virtual_network_rule":              "dbformariadb",
		"azurerm_mariadb_configuration":                     "dbformariadb",
		"azurerm_mysql_server":                              "dbformysql",
		"azurerm_mysql_configuration":                       "dbformysql",
		"azurerm_mysql_firewall_rule":                       "dbformysql",
		"azurerm_mysql_virtual_network_rule":                "dbformysql",
		"azurerm_mysql_active_directory_administrator":      "dbformysql",
		"azurerm_mysql_database":                            "dbformysql",
		"azurerm_mysql_flexible_server":                     "dbformysql",
		"azurerm_mysql_flexible_database":                   "dbformysql",
		"azurerm_mysql_flexible_server_configuration":       "dbformysql",
		"azurerm_mysql_flexible_server_firewall_rule":       "dbformysql",
		"azurerm_postgresql_server":                         "dbforpostgresql",
		"azurerm_postgresql_flexible_server_configuration":  "dbforpostgresql",
		"azurerm_postgresql_database":                       "dbforpostgresql",
		"azurerm_postgresql_active_directory_administrator": "dbforpostgresql",
		"azurerm_postgresql_flexible_server_database":       "dbforpostgresql",
		"azurerm_postgresql_flexible_server_firewall_rule":  "dbforpostgresql",
		"azurerm_postgresql_firewall_rule":                  "dbforpostgresql",
		"azurerm_postgresql_flexible_server":                "dbforpostgresql",
		"azurerm_postgresql_virtual_network_rule":           "dbforpostgresql",
		"azurerm_postgresql_server_key":                     "dbforpostgresql",
		"azurerm_postgresql_configuration":                  "dbforpostgresql",
		"azurerm_key_vault":                                 "keyvault",
		"azurerm_key_vault_key":                             "keyvault",

		// sql
		"azurerm_mssql_server":                                          "sql",
		"azurerm_mssql_database":                                        "sql",
		"azurerm_mssql_failover_group":                                  "sql",
		"azurerm_mssql_server_transparent_data_encryption":              "sql",
		"azurerm_mssql_virtual_network_rule":                            "sql",
		"azurerm_mssql_managed_instance":                                "sql",
		"azurerm_mssql_managed_database":                                "sql",
		"azurerm_mssql_managed_instance_active_directory_administrator": "sql",
		"azurerm_mssql_managed_instance_failover_group":                 "sql",
		"azurerm_mssql_managed_instance_vulnerability_assessment":       "sql",
		"azurerm_mssql_outbound_firewall_rule":                          "sql",
		"azurerm_mssql_server_dns_alias":                                "sql",
		"azurerm_mssql_database_extended_auditing_policy":               "sql",
		"azurerm_mssql_server_security_alert_policy":                    "sql",
		"azurerm_mssql_firewall_rule":                                   "sql",
		"azurerm_mssql_database_vulnerability_assessment_rule_baseline": "sql",
		"azurerm_mssql_elasticpool":                                     "sql",
		"azurerm_mssql_job_agent":                                       "sql",
		"azurerm_mssql_job_credential":                                  "sql",
		"azurerm_mssql_server_vulnerability_assessment":                 "sql",
		"azurerm_storage_account":                                       "storage",
		"azurerm_storage_container":                                     "storage",
		"azurerm_network_security_group":                                "network",
	}
	resourceKind = map[string]string{
		"azurerm_virtual_network_peering":                   "VirtualNetworkPeering",
		"azurerm_virtual_network":                           "VirtualNetwork",
		"azurerm_private_dns_zone":                          "PrivateDNSZone",
		"azurerm_private_dns_zone_virtual_network_link":     "PrivateDNSZoneVirtualNetworkLink",
		"azurerm_subnet":                                    "Subnet",
		"azurerm_network_security_group":                    "SecurityGroup",
		"azurerm_subnet_network_security_group_association": "SubnetNetworkSecurityGroupAssociation",
		"azurerm_subnet_route_table_association":            "SubnetRouteTableAssociation",
		"azurerm_route_table":                               "RouteTable",
		"azurerm_redis_cache":                               "RedisCache",
		"azurerm_redis_firewall_rule":                       "RedisFirewallRule",
		"azurerm_redis_enterprise_cluster":                  "RedisEnterpriseCluster",
		"azurerm_redis_enterprise_database":                 "RedisEnterpriseDatabase",
		"azurerm_redis_linked_server":                       "RedisLinkedServer",
		"azurerm_cosmosdb_sql_database":                     "SQLDatabase",
		"azurerm_cosmosdb_sql_container":                    "SQLContainer",
		"azurerm_cosmosdb_sql_role_assignment":              "SQLRoleAssignment",
		"azurerm_cosmosdb_sql_role_definition":              "SQLRoleDefinition",
		"azurerm_cosmosdb_mongo_database":                   "MongoDatabase",
		"azurerm_cosmosdb_mongo_collection":                 "MongoCollection",
		"azurerm_cosmosdb_cassandra_keyspace":               "CassandraKeySpace",
		"azurerm_cosmosdb_cassandra_table":                  "CassandraTable",
		"azurerm_cosmosdb_gremlin_database":                 "GremlinDatabase",
		"azurerm_cosmosdb_gremlin_graph":                    "GremlinGraph",
		"azurerm_cosmosdb_sql_stored_procedure":             "SQLStoredProcedure",
		"azurerm_cosmosdb_sql_function":                     "SQLFunction",
		"azurerm_cosmosdb_table":                            "Table",
		"azurerm_cosmosdb_account":                          "Account",
		"azurerm_cosmosdb_sql_trigger":                      "SQLTrigger",
		"azurerm_cosmosdb_cassandra_cluster":                "CassandraCluster",
		"azurerm_cosmosdb_cassandra_datacenter":             "CassandraDatacenter",
		"azurerm_cosmosdb_sql_dedicated_gateway":            "SQLDedicatedGateway",
		"azurerm_mariadb_server":                            "Server",
		"azurerm_mariadb_database":                          "Database",
		"azurerm_mariadb_firewall_rule":                     "FirewallRule",
		"azurerm_mariadb_virtual_network_rule":              "VirtualNetworkRule",
		"azurerm_mariadb_configuration":                     "Configuration",
		"azurerm_mysql_server":                              "Server",
		"azurerm_mysql_configuration":                       "Configuration",
		"azurerm_mysql_firewall_rule":                       "FirewallRule",
		"azurerm_mysql_virtual_network_rule":                "VirtualNetworkRule",
		"azurerm_mysql_active_directory_administrator":      "ActiveDirectoryAdministrator",
		"azurerm_mysql_database":                            "Database",
		"azurerm_mysql_flexible_server":                     "FlexibleServer",
		"azurerm_mysql_flexible_database":                   "FlexibleDatabase",
		"azurerm_mysql_flexible_server_configuration":       "FlexibleServerConfiguration",
		"azurerm_mysql_flexible_server_firewall_rule":       "FlexibleServerFirewallRule",
		"azurerm_postgresql_server":                         "Server",
		"azurerm_postgresql_flexible_server_configuration":  "FlexibleServerConfiguration",
		"azurerm_postgresql_database":                       "Database",
		"azurerm_postgresql_active_directory_administrator": "ActiveDirectoryAdministrator",
		"azurerm_postgresql_flexible_server_database":       "FlexibleServerDatabase",
		"azurerm_postgresql_flexible_server_firewall_rule":  "FlexibleServerFirewallRule",
		"azurerm_postgresql_firewall_rule":                  "FirewallRule",
		"azurerm_postgresql_flexible_server":                "FlexibleServer",
		"azurerm_postgresql_virtual_network_rule":           "VirtualNetworkRule",
		"azurerm_postgresql_server_key":                     "ServerKey",
		"azurerm_postgresql_configuration":                  "Configuration",
		"azurerm_key_vault":                                 "Vault",
		"azurerm_key_vault_key":                             "Key",

		// sql
		"azurerm_mssql_server":                                          "MSSQLServer",
		"azurerm_mssql_database":                                        "MSSQLDatabase",
		"azurerm_mssql_failover_group":                                  "MSSQLFailoverGroup",
		"azurerm_mssql_server_transparent_data_encryption":              "MSSQLServerTransparentDataEncryption",
		"azurerm_mssql_virtual_network_rule":                            "MSSQLVirtualNetworkRule",
		"azurerm_mssql_managed_instance":                                "MSSQLManagedInstance",
		"azurerm_mssql_managed_database":                                "MSSQLManagedDatabase",
		"azurerm_mssql_managed_instance_active_directory_administrator": "MSSQLManagedInstanceActiveDirectoryAdministrator",
		"azurerm_mssql_managed_instance_failover_group":                 "MSSQLManagedInstanceFailoverGroup",
		"azurerm_mssql_managed_instance_vulnerability_assessment":       "MSSQLManagedInstanceVulnerabilityAssessment",
		"azurerm_mssql_outbound_firewall_rule":                          "MSSQLOutboundFirewallRule",
		"azurerm_mssql_server_dns_alias":                                "MSSQLServerDNSAlias",
		"azurerm_mssql_database_extended_auditing_policy":               "MSSQLServerMicrosoftSupportAuditingPolicy",
		"azurerm_mssql_server_security_alert_policy":                    "MSSQLServerSecurityAlertPolicy",
		"azurerm_mssql_firewall_rule":                                   "MSSQLFirewallRule",
		"azurerm_mssql_database_vulnerability_assessment_rule_baseline": "MSSQLDatabaseVulnerabilityAssessmentRuleBaseline",
		"azurerm_mssql_elasticpool":                                     "MSSQLElasticPool",
		"azurerm_mssql_job_agent":                                       "MSSQLJobAgent",
		"azurerm_mssql_job_credential":                                  "MSSQLJobCredential",
		"azurerm_mssql_server_vulnerability_assessment":                 "MSSQLServerVulnerabilityAssessment",
		"azurerm_storage_account":                                       "Account",
		"azurerm_storage_container":                                     "Container",
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
