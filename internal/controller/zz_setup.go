/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	rediscache "kubeform.dev/provider-azure/internal/controller/cache/rediscache"
	redisenterprisecluster "kubeform.dev/provider-azure/internal/controller/cache/redisenterprisecluster"
	redisenterprisedatabase "kubeform.dev/provider-azure/internal/controller/cache/redisenterprisedatabase"
	redisfirewallrule "kubeform.dev/provider-azure/internal/controller/cache/redisfirewallrule"
	redislinkedserver "kubeform.dev/provider-azure/internal/controller/cache/redislinkedserver"
	account "kubeform.dev/provider-azure/internal/controller/cosmosdb/account"
	cassandracluster "kubeform.dev/provider-azure/internal/controller/cosmosdb/cassandracluster"
	cassandradatacenter "kubeform.dev/provider-azure/internal/controller/cosmosdb/cassandradatacenter"
	cassandrakeyspace "kubeform.dev/provider-azure/internal/controller/cosmosdb/cassandrakeyspace"
	cassandratable "kubeform.dev/provider-azure/internal/controller/cosmosdb/cassandratable"
	gremlindatabase "kubeform.dev/provider-azure/internal/controller/cosmosdb/gremlindatabase"
	gremlingraph "kubeform.dev/provider-azure/internal/controller/cosmosdb/gremlingraph"
	mongocollection "kubeform.dev/provider-azure/internal/controller/cosmosdb/mongocollection"
	mongodatabase "kubeform.dev/provider-azure/internal/controller/cosmosdb/mongodatabase"
	sqlcontainer "kubeform.dev/provider-azure/internal/controller/cosmosdb/sqlcontainer"
	sqldatabase "kubeform.dev/provider-azure/internal/controller/cosmosdb/sqldatabase"
	sqldedicatedgateway "kubeform.dev/provider-azure/internal/controller/cosmosdb/sqldedicatedgateway"
	sqlfunction "kubeform.dev/provider-azure/internal/controller/cosmosdb/sqlfunction"
	sqlroleassignment "kubeform.dev/provider-azure/internal/controller/cosmosdb/sqlroleassignment"
	sqlroledefinition "kubeform.dev/provider-azure/internal/controller/cosmosdb/sqlroledefinition"
	sqlstoredprocedure "kubeform.dev/provider-azure/internal/controller/cosmosdb/sqlstoredprocedure"
	sqltrigger "kubeform.dev/provider-azure/internal/controller/cosmosdb/sqltrigger"
	table "kubeform.dev/provider-azure/internal/controller/cosmosdb/table"
	configuration "kubeform.dev/provider-azure/internal/controller/dbformariadb/configuration"
	database "kubeform.dev/provider-azure/internal/controller/dbformariadb/database"
	firewallrule "kubeform.dev/provider-azure/internal/controller/dbformariadb/firewallrule"
	server "kubeform.dev/provider-azure/internal/controller/dbformariadb/server"
	virtualnetworkrule "kubeform.dev/provider-azure/internal/controller/dbformariadb/virtualnetworkrule"
	activedirectoryadministrator "kubeform.dev/provider-azure/internal/controller/dbformysql/activedirectoryadministrator"
	configurationdbformysql "kubeform.dev/provider-azure/internal/controller/dbformysql/configuration"
	databasedbformysql "kubeform.dev/provider-azure/internal/controller/dbformysql/database"
	firewallruledbformysql "kubeform.dev/provider-azure/internal/controller/dbformysql/firewallrule"
	flexibledatabase "kubeform.dev/provider-azure/internal/controller/dbformysql/flexibledatabase"
	flexibleserver "kubeform.dev/provider-azure/internal/controller/dbformysql/flexibleserver"
	flexibleserverconfiguration "kubeform.dev/provider-azure/internal/controller/dbformysql/flexibleserverconfiguration"
	flexibleserverfirewallrule "kubeform.dev/provider-azure/internal/controller/dbformysql/flexibleserverfirewallrule"
	serverdbformysql "kubeform.dev/provider-azure/internal/controller/dbformysql/server"
	virtualnetworkruledbformysql "kubeform.dev/provider-azure/internal/controller/dbformysql/virtualnetworkrule"
	activedirectoryadministratordbforpostgresql "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/activedirectoryadministrator"
	configurationdbforpostgresql "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/configuration"
	databasedbforpostgresql "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/database"
	firewallruledbforpostgresql "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/firewallrule"
	flexibleserverdbforpostgresql "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/flexibleserver"
	flexibleserverconfigurationdbforpostgresql "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallruledbforpostgresql "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/flexibleserverfirewallrule"
	serverdbforpostgresql "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/server"
	serverkey "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/serverkey"
	virtualnetworkruledbforpostgresql "kubeform.dev/provider-azure/internal/controller/dbforpostgresql/virtualnetworkrule"
	key "kubeform.dev/provider-azure/internal/controller/keyvault/key"
	vault "kubeform.dev/provider-azure/internal/controller/keyvault/vault"
	privatednszone "kubeform.dev/provider-azure/internal/controller/network/privatednszone"
	virtualnetwork "kubeform.dev/provider-azure/internal/controller/network/virtualnetwork"
	virtualnetworkpeering "kubeform.dev/provider-azure/internal/controller/network/virtualnetworkpeering"
	providerconfig "kubeform.dev/provider-azure/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		rediscache.Setup,
		redisenterprisecluster.Setup,
		redisenterprisedatabase.Setup,
		redisfirewallrule.Setup,
		redislinkedserver.Setup,
		account.Setup,
		cassandracluster.Setup,
		cassandradatacenter.Setup,
		cassandrakeyspace.Setup,
		cassandratable.Setup,
		gremlindatabase.Setup,
		gremlingraph.Setup,
		mongocollection.Setup,
		mongodatabase.Setup,
		sqlcontainer.Setup,
		sqldatabase.Setup,
		sqldedicatedgateway.Setup,
		sqlfunction.Setup,
		sqlroleassignment.Setup,
		sqlroledefinition.Setup,
		sqlstoredprocedure.Setup,
		sqltrigger.Setup,
		table.Setup,
		configuration.Setup,
		database.Setup,
		firewallrule.Setup,
		server.Setup,
		virtualnetworkrule.Setup,
		activedirectoryadministrator.Setup,
		configurationdbformysql.Setup,
		databasedbformysql.Setup,
		firewallruledbformysql.Setup,
		flexibledatabase.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverfirewallrule.Setup,
		serverdbformysql.Setup,
		virtualnetworkruledbformysql.Setup,
		activedirectoryadministratordbforpostgresql.Setup,
		configurationdbforpostgresql.Setup,
		databasedbforpostgresql.Setup,
		firewallruledbforpostgresql.Setup,
		flexibleserverdbforpostgresql.Setup,
		flexibleserverconfigurationdbforpostgresql.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallruledbforpostgresql.Setup,
		serverdbforpostgresql.Setup,
		serverkey.Setup,
		virtualnetworkruledbforpostgresql.Setup,
		key.Setup,
		vault.Setup,
		privatednszone.Setup,
		virtualnetwork.Setup,
		virtualnetworkpeering.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
