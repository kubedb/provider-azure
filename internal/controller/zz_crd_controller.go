package controller

import (
	"context"
	"sync"

	"github.com/crossplane/upjet/pkg/controller"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	roleassignment "kubedb.dev/provider-azure/internal/controller/authorization/roleassignment"
	providerregistration "kubedb.dev/provider-azure/internal/controller/azure/providerregistration"
	resourcegroup "kubedb.dev/provider-azure/internal/controller/azure/resourcegroup"
	subscription "kubedb.dev/provider-azure/internal/controller/azure/subscription"
	rediscache "kubedb.dev/provider-azure/internal/controller/cache/rediscache"
	redisenterprisecluster "kubedb.dev/provider-azure/internal/controller/cache/redisenterprisecluster"
	redisenterprisedatabase "kubedb.dev/provider-azure/internal/controller/cache/redisenterprisedatabase"
	redisfirewallrule "kubedb.dev/provider-azure/internal/controller/cache/redisfirewallrule"
	redislinkedserver "kubedb.dev/provider-azure/internal/controller/cache/redislinkedserver"
	account "kubedb.dev/provider-azure/internal/controller/cosmosdb/account"
	cassandracluster "kubedb.dev/provider-azure/internal/controller/cosmosdb/cassandracluster"
	cassandradatacenter "kubedb.dev/provider-azure/internal/controller/cosmosdb/cassandradatacenter"
	cassandrakeyspace "kubedb.dev/provider-azure/internal/controller/cosmosdb/cassandrakeyspace"
	cassandratable "kubedb.dev/provider-azure/internal/controller/cosmosdb/cassandratable"
	gremlindatabase "kubedb.dev/provider-azure/internal/controller/cosmosdb/gremlindatabase"
	gremlingraph "kubedb.dev/provider-azure/internal/controller/cosmosdb/gremlingraph"
	mongocollection "kubedb.dev/provider-azure/internal/controller/cosmosdb/mongocollection"
	mongodatabase "kubedb.dev/provider-azure/internal/controller/cosmosdb/mongodatabase"
	sqlcontainer "kubedb.dev/provider-azure/internal/controller/cosmosdb/sqlcontainer"
	sqldatabase "kubedb.dev/provider-azure/internal/controller/cosmosdb/sqldatabase"
	sqldedicatedgateway "kubedb.dev/provider-azure/internal/controller/cosmosdb/sqldedicatedgateway"
	sqlfunction "kubedb.dev/provider-azure/internal/controller/cosmosdb/sqlfunction"
	sqlroleassignment "kubedb.dev/provider-azure/internal/controller/cosmosdb/sqlroleassignment"
	sqlroledefinition "kubedb.dev/provider-azure/internal/controller/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "kubedb.dev/provider-azure/internal/controller/cosmosdb/sqlstoredprocedure"
	sqltrigger "kubedb.dev/provider-azure/internal/controller/cosmosdb/sqltrigger"
	table "kubedb.dev/provider-azure/internal/controller/cosmosdb/table"
	configuration "kubedb.dev/provider-azure/internal/controller/dbformariadb/configuration"
	database "kubedb.dev/provider-azure/internal/controller/dbformariadb/database"
	firewallrule "kubedb.dev/provider-azure/internal/controller/dbformariadb/firewallrule"
	server "kubedb.dev/provider-azure/internal/controller/dbformariadb/server"
	virtualnetworkrule "kubedb.dev/provider-azure/internal/controller/dbformariadb/virtualnetworkrule"
	activedirectoryadministrator "kubedb.dev/provider-azure/internal/controller/dbformysql/activedirectoryadministrator"
	configurationdbformysql "kubedb.dev/provider-azure/internal/controller/dbformysql/configuration"
	databasedbformysql "kubedb.dev/provider-azure/internal/controller/dbformysql/database"
	firewallruledbformysql "kubedb.dev/provider-azure/internal/controller/dbformysql/firewallrule"
	flexibledatabase "kubedb.dev/provider-azure/internal/controller/dbformysql/flexibledatabase"
	flexibleserver "kubedb.dev/provider-azure/internal/controller/dbformysql/flexibleserver"
	flexibleserverconfiguration "kubedb.dev/provider-azure/internal/controller/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "kubedb.dev/provider-azure/internal/controller/dbformysql/flexibleserverfirewallrule"
	serverdbformysql "kubedb.dev/provider-azure/internal/controller/dbformysql/server"
	virtualnetworkruledbformysql "kubedb.dev/provider-azure/internal/controller/dbformysql/virtualnetworkrule"
	activedirectoryadministratordbforpostgresql "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/activedirectoryadministrator"
	configurationdbforpostgresql "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/configuration"
	databasedbforpostgresql "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/database"
	firewallruledbforpostgresql "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/firewallrule"
	flexibleserverdbforpostgresql "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/flexibleserver"
	flexibleserverconfigurationdbforpostgresql "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallruledbforpostgresql "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/flexibleserverfirewallrule"
	serverdbforpostgresql "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/server"
	serverkey "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/serverkey"
	virtualnetworkruledbforpostgresql "kubedb.dev/provider-azure/internal/controller/dbforpostgresql/virtualnetworkrule"
	key "kubedb.dev/provider-azure/internal/controller/keyvault/key"
	vault "kubedb.dev/provider-azure/internal/controller/keyvault/vault"
	privatednszone "kubedb.dev/provider-azure/internal/controller/network/privatednszone"
	privatednszonevirtualnetworklink "kubedb.dev/provider-azure/internal/controller/network/privatednszonevirtualnetworklink"
	routetable "kubedb.dev/provider-azure/internal/controller/network/routetable"
	securitygroup "kubedb.dev/provider-azure/internal/controller/network/securitygroup"
	subnet "kubedb.dev/provider-azure/internal/controller/network/subnet"
	subnetnetworksecuritygroupassociation "kubedb.dev/provider-azure/internal/controller/network/subnetnetworksecuritygroupassociation"
	subnetroutetableassociation "kubedb.dev/provider-azure/internal/controller/network/subnetroutetableassociation"
	virtualnetwork "kubedb.dev/provider-azure/internal/controller/network/virtualnetwork"
	virtualnetworkpeering "kubedb.dev/provider-azure/internal/controller/network/virtualnetworkpeering"
	providerconfig "kubedb.dev/provider-azure/internal/controller/providerconfig"
	mssqldatabase "kubedb.dev/provider-azure/internal/controller/sql/mssqldatabase"
	mssqldatabasevulnerabilityassessmentrulebaseline "kubedb.dev/provider-azure/internal/controller/sql/mssqldatabasevulnerabilityassessmentrulebaseline"
	mssqlelasticpool "kubedb.dev/provider-azure/internal/controller/sql/mssqlelasticpool"
	mssqlfailovergroup "kubedb.dev/provider-azure/internal/controller/sql/mssqlfailovergroup"
	mssqlfirewallrule "kubedb.dev/provider-azure/internal/controller/sql/mssqlfirewallrule"
	mssqljobagent "kubedb.dev/provider-azure/internal/controller/sql/mssqljobagent"
	mssqljobcredential "kubedb.dev/provider-azure/internal/controller/sql/mssqljobcredential"
	mssqlmanageddatabase "kubedb.dev/provider-azure/internal/controller/sql/mssqlmanageddatabase"
	mssqlmanagedinstance "kubedb.dev/provider-azure/internal/controller/sql/mssqlmanagedinstance"
	mssqlmanagedinstanceactivedirectoryadministrator "kubedb.dev/provider-azure/internal/controller/sql/mssqlmanagedinstanceactivedirectoryadministrator"
	mssqlmanagedinstancefailovergroup "kubedb.dev/provider-azure/internal/controller/sql/mssqlmanagedinstancefailovergroup"
	mssqlmanagedinstancevulnerabilityassessment "kubedb.dev/provider-azure/internal/controller/sql/mssqlmanagedinstancevulnerabilityassessment"
	mssqloutboundfirewallrule "kubedb.dev/provider-azure/internal/controller/sql/mssqloutboundfirewallrule"
	mssqlserver "kubedb.dev/provider-azure/internal/controller/sql/mssqlserver"
	mssqlserverdnsalias "kubedb.dev/provider-azure/internal/controller/sql/mssqlserverdnsalias"
	mssqlservermicrosoftsupportauditingpolicy "kubedb.dev/provider-azure/internal/controller/sql/mssqlservermicrosoftsupportauditingpolicy"
	mssqlserversecurityalertpolicy "kubedb.dev/provider-azure/internal/controller/sql/mssqlserversecurityalertpolicy"
	mssqlservertransparentdataencryption "kubedb.dev/provider-azure/internal/controller/sql/mssqlservertransparentdataencryption"
	mssqlservervulnerabilityassessment "kubedb.dev/provider-azure/internal/controller/sql/mssqlservervulnerabilityassessment"
	mssqlvirtualnetworkrule "kubedb.dev/provider-azure/internal/controller/sql/mssqlvirtualnetworkrule"
	accountstorage "kubedb.dev/provider-azure/internal/controller/storage/account"
	container "kubedb.dev/provider-azure/internal/controller/storage/container"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	setupFns = map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{"roleassignment.authorization.azure.kubedb.com", "RoleAssignment"}:                                                           roleassignment.Setup,
		schema.GroupKind{"providerregistration.azure.kubedb.com", "ProviderRegistration"}:                                                             providerregistration.Setup,
		schema.GroupKind{"resourcegroup.azure.kubedb.com", "ResourceGroup"}:                                                                           resourcegroup.Setup,
		schema.GroupKind{"subscription.azure.kubedb.com", "Subscription"}:                                                                             subscription.Setup,
		schema.GroupKind{"rediscache.cache.azure.kubedb.com", "RedisCache"}:                                                                           rediscache.Setup,
		schema.GroupKind{"redisenterprisecluster.cache.azure.kubedb.com", "RedisEnterpriseCluster"}:                                                   redisenterprisecluster.Setup,
		schema.GroupKind{"redisenterprisedatabase.cache.azure.kubedb.com", "RedisEnterpriseDatabase"}:                                                 redisenterprisedatabase.Setup,
		schema.GroupKind{"redisfirewallrule.cache.azure.kubedb.com", "RedisFirewallRule"}:                                                             redisfirewallrule.Setup,
		schema.GroupKind{"redislinkedserver.cache.azure.kubedb.com", "RedisLinkedServer"}:                                                             redislinkedserver.Setup,
		schema.GroupKind{"account.cosmosdb.azure.kubedb.com", "Account"}:                                                                              account.Setup,
		schema.GroupKind{"cassandracluster.cosmosdb.azure.kubedb.com", "CassandraCluster"}:                                                            cassandracluster.Setup,
		schema.GroupKind{"cassandradatacenter.cosmosdb.azure.kubedb.com", "CassandraDatacenter"}:                                                      cassandradatacenter.Setup,
		schema.GroupKind{"cassandrakeyspace.cosmosdb.azure.kubedb.com", "CassandraKeySpace"}:                                                          cassandrakeyspace.Setup,
		schema.GroupKind{"cassandratable.cosmosdb.azure.kubedb.com", "CassandraTable"}:                                                                cassandratable.Setup,
		schema.GroupKind{"gremlindatabase.cosmosdb.azure.kubedb.com", "GremlinDatabase"}:                                                              gremlindatabase.Setup,
		schema.GroupKind{"gremlingraph.cosmosdb.azure.kubedb.com", "GremlinGraph"}:                                                                    gremlingraph.Setup,
		schema.GroupKind{"mongocollection.cosmosdb.azure.kubedb.com", "MongoCollection"}:                                                              mongocollection.Setup,
		schema.GroupKind{"mongodatabase.cosmosdb.azure.kubedb.com", "MongoDatabase"}:                                                                  mongodatabase.Setup,
		schema.GroupKind{"sqlcontainer.cosmosdb.azure.kubedb.com", "SQLContainer"}:                                                                    sqlcontainer.Setup,
		schema.GroupKind{"sqldatabase.cosmosdb.azure.kubedb.com", "SQLDatabase"}:                                                                      sqldatabase.Setup,
		schema.GroupKind{"sqldedicatedgateway.cosmosdb.azure.kubedb.com", "SQLDedicatedGateway"}:                                                      sqldedicatedgateway.Setup,
		schema.GroupKind{"sqlfunction.cosmosdb.azure.kubedb.com", "SQLFunction"}:                                                                      sqlfunction.Setup,
		schema.GroupKind{"sqlroleassignment.cosmosdb.azure.kubedb.com", "SQLRoleAssignment"}:                                                          sqlroleassignment.Setup,
		schema.GroupKind{"sqlroledefinition.cosmosdb.azure.kubedb.com", "SQLRoleDefinition"}:                                                          sqlroledefinition.Setup,
		schema.GroupKind{"sqlstoredprocedure.cosmosdb.azure.kubedb.com", "SQLStoredProcedure"}:                                                        sqlstoredprocedure.Setup,
		schema.GroupKind{"sqltrigger.cosmosdb.azure.kubedb.com", "SQLTrigger"}:                                                                        sqltrigger.Setup,
		schema.GroupKind{"table.cosmosdb.azure.kubedb.com", "Table"}:                                                                                  table.Setup,
		schema.GroupKind{"configuration.dbformariadb.azure.kubedb.com", "Configuration"}:                                                              configuration.Setup,
		schema.GroupKind{"database.dbformariadb.azure.kubedb.com", "Database"}:                                                                        database.Setup,
		schema.GroupKind{"firewallrule.dbformariadb.azure.kubedb.com", "FirewallRule"}:                                                                firewallrule.Setup,
		schema.GroupKind{"server.dbformariadb.azure.kubedb.com", "Server"}:                                                                            server.Setup,
		schema.GroupKind{"virtualnetworkrule.dbformariadb.azure.kubedb.com", "VirtualNetworkRule"}:                                                    virtualnetworkrule.Setup,
		schema.GroupKind{"activedirectoryadministrator.dbformysql.azure.kubedb.com", "ActiveDirectoryAdministrator"}:                                  activedirectoryadministrator.Setup,
		schema.GroupKind{"configuration.dbformysql.azure.kubedb.com", "Configuration"}:                                                                configurationdbformysql.Setup,
		schema.GroupKind{"database.dbformysql.azure.kubedb.com", "Database"}:                                                                          databasedbformysql.Setup,
		schema.GroupKind{"firewallrule.dbformysql.azure.kubedb.com", "FirewallRule"}:                                                                  firewallruledbformysql.Setup,
		schema.GroupKind{"flexibledatabase.dbformysql.azure.kubedb.com", "FlexibleDatabase"}:                                                          flexibledatabase.Setup,
		schema.GroupKind{"flexibleserver.dbformysql.azure.kubedb.com", "FlexibleServer"}:                                                              flexibleserver.Setup,
		schema.GroupKind{"flexibleserverconfiguration.dbformysql.azure.kubedb.com", "FlexibleServerConfiguration"}:                                    flexibleserverconfiguration.Setup,
		schema.GroupKind{"flexibleserverfirewallrule.dbformysql.azure.kubedb.com", "FlexibleServerFirewallRule"}:                                      flexibleserverfirewallrule.Setup,
		schema.GroupKind{"server.dbformysql.azure.kubedb.com", "Server"}:                                                                              serverdbformysql.Setup,
		schema.GroupKind{"virtualnetworkrule.dbformysql.azure.kubedb.com", "VirtualNetworkRule"}:                                                      virtualnetworkruledbformysql.Setup,
		schema.GroupKind{"activedirectoryadministrator.dbforpostgresql.azure.kubedb.com", "ActiveDirectoryAdministrator"}:                             activedirectoryadministratordbforpostgresql.Setup,
		schema.GroupKind{"configuration.dbforpostgresql.azure.kubedb.com", "Configuration"}:                                                           configurationdbforpostgresql.Setup,
		schema.GroupKind{"database.dbforpostgresql.azure.kubedb.com", "Database"}:                                                                     databasedbforpostgresql.Setup,
		schema.GroupKind{"firewallrule.dbforpostgresql.azure.kubedb.com", "FirewallRule"}:                                                             firewallruledbforpostgresql.Setup,
		schema.GroupKind{"flexibleserver.dbforpostgresql.azure.kubedb.com", "FlexibleServer"}:                                                         flexibleserverdbforpostgresql.Setup,
		schema.GroupKind{"flexibleserverconfiguration.dbforpostgresql.azure.kubedb.com", "FlexibleServerConfiguration"}:                               flexibleserverconfigurationdbforpostgresql.Setup,
		schema.GroupKind{"flexibleserverdatabase.dbforpostgresql.azure.kubedb.com", "FlexibleServerDatabase"}:                                         flexibleserverdatabase.Setup,
		schema.GroupKind{"flexibleserverfirewallrule.dbforpostgresql.azure.kubedb.com", "FlexibleServerFirewallRule"}:                                 flexibleserverfirewallruledbforpostgresql.Setup,
		schema.GroupKind{"server.dbforpostgresql.azure.kubedb.com", "Server"}:                                                                         serverdbforpostgresql.Setup,
		schema.GroupKind{"serverkey.dbforpostgresql.azure.kubedb.com", "ServerKey"}:                                                                   serverkey.Setup,
		schema.GroupKind{"virtualnetworkrule.dbforpostgresql.azure.kubedb.com", "VirtualNetworkRule"}:                                                 virtualnetworkruledbforpostgresql.Setup,
		schema.GroupKind{"key.keyvault.azure.kubedb.com", "Key"}:                                                                                      key.Setup,
		schema.GroupKind{"vault.keyvault.azure.kubedb.com", "Vault"}:                                                                                  vault.Setup,
		schema.GroupKind{"privatednszone.network.azure.kubedb.com", "PrivateDNSZone"}:                                                                 privatednszone.Setup,
		schema.GroupKind{"privatednszonevirtualnetworklink.network.azure.kubedb.com", "PrivateDNSZoneVirtualNetworkLink"}:                             privatednszonevirtualnetworklink.Setup,
		schema.GroupKind{"routetable.network.azure.kubedb.com", "RouteTable"}:                                                                         routetable.Setup,
		schema.GroupKind{"securitygroup.network.azure.kubedb.com", "SecurityGroup"}:                                                                   securitygroup.Setup,
		schema.GroupKind{"subnet.network.azure.kubedb.com", "Subnet"}:                                                                                 subnet.Setup,
		schema.GroupKind{"subnetnetworksecuritygroupassociation.network.azure.kubedb.com", "SubnetNetworkSecurityGroupAssociation"}:                   subnetnetworksecuritygroupassociation.Setup,
		schema.GroupKind{"subnetroutetableassociation.network.azure.kubedb.com", "SubnetRouteTableAssociation"}:                                       subnetroutetableassociation.Setup,
		schema.GroupKind{"virtualnetwork.network.azure.kubedb.com", "VirtualNetwork"}:                                                                 virtualnetwork.Setup,
		schema.GroupKind{"virtualnetworkpeering.network.azure.kubedb.com", "VirtualNetworkPeering"}:                                                   virtualnetworkpeering.Setup,
		schema.GroupKind{"providerconfig.controller.azure.kubedb.com", ""}:                                                                            providerconfig.Setup,
		schema.GroupKind{"mssqldatabase.sql.azure.kubedb.com", "MSSQLDatabase"}:                                                                       mssqldatabase.Setup,
		schema.GroupKind{"mssqldatabasevulnerabilityassessmentrulebaseline.sql.azure.kubedb.com", "MSSQLDatabaseVulnerabilityAssessmentRuleBaseline"}: mssqldatabasevulnerabilityassessmentrulebaseline.Setup,
		schema.GroupKind{"mssqlelasticpool.sql.azure.kubedb.com", "MSSQLElasticPool"}:                                                                 mssqlelasticpool.Setup,
		schema.GroupKind{"mssqlfailovergroup.sql.azure.kubedb.com", "MSSQLFailoverGroup"}:                                                             mssqlfailovergroup.Setup,
		schema.GroupKind{"mssqlfirewallrule.sql.azure.kubedb.com", "MSSQLFirewallRule"}:                                                               mssqlfirewallrule.Setup,
		schema.GroupKind{"mssqljobagent.sql.azure.kubedb.com", "MSSQLJobAgent"}:                                                                       mssqljobagent.Setup,
		schema.GroupKind{"mssqljobcredential.sql.azure.kubedb.com", "MSSQLJobCredential"}:                                                             mssqljobcredential.Setup,
		schema.GroupKind{"mssqlmanageddatabase.sql.azure.kubedb.com", "MSSQLManagedDatabase"}:                                                         mssqlmanageddatabase.Setup,
		schema.GroupKind{"mssqlmanagedinstance.sql.azure.kubedb.com", "MSSQLManagedInstance"}:                                                         mssqlmanagedinstance.Setup,
		schema.GroupKind{"mssqlmanagedinstanceactivedirectoryadministrator.sql.azure.kubedb.com", "MSSQLManagedInstanceActiveDirectoryAdministrator"}: mssqlmanagedinstanceactivedirectoryadministrator.Setup,
		schema.GroupKind{"mssqlmanagedinstancefailovergroup.sql.azure.kubedb.com", "MSSQLManagedInstanceFailoverGroup"}:                               mssqlmanagedinstancefailovergroup.Setup,
		schema.GroupKind{"mssqlmanagedinstancevulnerabilityassessment.sql.azure.kubedb.com", "MSSQLManagedInstanceVulnerabilityAssessment"}:           mssqlmanagedinstancevulnerabilityassessment.Setup,
		schema.GroupKind{"mssqloutboundfirewallrule.sql.azure.kubedb.com", "MSSQLOutboundFirewallRule"}:                                               mssqloutboundfirewallrule.Setup,
		schema.GroupKind{"mssqlserver.sql.azure.kubedb.com", "MSSQLServer"}:                                                                           mssqlserver.Setup,
		schema.GroupKind{"mssqlserverdnsalias.sql.azure.kubedb.com", "MSSQLServerDNSAlias"}:                                                           mssqlserverdnsalias.Setup,
		schema.GroupKind{"mssqlservermicrosoftsupportauditingpolicy.sql.azure.kubedb.com", "MSSQLServerMicrosoftSupportAuditingPolicy"}:               mssqlservermicrosoftsupportauditingpolicy.Setup,
		schema.GroupKind{"mssqlserversecurityalertpolicy.sql.azure.kubedb.com", "MSSQLServerSecurityAlertPolicy"}:                                     mssqlserversecurityalertpolicy.Setup,
		schema.GroupKind{"mssqlservertransparentdataencryption.sql.azure.kubedb.com", "MSSQLServerTransparentDataEncryption"}:                         mssqlservertransparentdataencryption.Setup,
		schema.GroupKind{"mssqlservervulnerabilityassessment.sql.azure.kubedb.com", "MSSQLServerVulnerabilityAssessment"}:                             mssqlservervulnerabilityassessment.Setup,
		schema.GroupKind{"mssqlvirtualnetworkrule.sql.azure.kubedb.com", "MSSQLVirtualNetworkRule"}:                                                   mssqlvirtualnetworkrule.Setup,
		schema.GroupKind{"account.storage.azure.kubedb.com", "Account"}:                                                                               accountstorage.Setup,
		schema.GroupKind{"container.storage.azure.kubedb.com", "Container"}:                                                                           container.Setup,
	}
)

//package controller

/*
var(
    gk2      = schema.GroupKind{"azure.kubedb.com", "ResourceGroup"}
	gk3      = schema.GroupKind{"azure.kubedb.com", "ProviderConfig"}
	setupFns = map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		gk2: resourcegroup.Setup,
		gk3: providerregistration.Setup,
	}
)
*/

var (
	setupDone = map[schema.GroupKind]bool{}
	mu        sync.RWMutex
)

//func SetupControllerList(mgr ctrl.Manager, o controller.Options) error {
//
//}

type CustomResourceReconciler struct {
	mgr ctrl.Manager
	o   controller.Options
}

func NewCustomResourceReconciler(mgr ctrl.Manager, o controller.Options) *CustomResourceReconciler {
	//if err := SetupControllerList(mgr, o); err != nil {
	//	log.Error(err, "unable to fetch CustomResourceDefinition")
	//}
	return &CustomResourceReconciler{mgr: mgr, o: o}
}

func (r *CustomResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var crd apiextensions.CustomResourceDefinition
	if err := r.mgr.GetClient().Get(ctx, req.NamespacedName, &crd); err != nil {
		log.Error(err, "unable to fetch CustomResourceDefinition")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	gk := schema.GroupKind{
		Group: crd.Spec.Group,
		Kind:  crd.Spec.Names.Kind,
	}
	mu.Lock()
	defer mu.Unlock()
	_, found := setupDone[gk]
	if found {
		return ctrl.Result{}, nil
	}
	setup, found := setupFns[gk]
	if found {
		setup(r.mgr, r.o)
		setupDone[gk] = true
	}

	return ctrl.Result{}, nil
}

func (r *CustomResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiextensions.CustomResourceDefinition{}).
		Complete(r)
}
