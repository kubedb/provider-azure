// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

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
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		roleassignment.Setup,
		providerregistration.Setup,
		providerregistration.Setup,
		resourcegroup.Setup,
		resourcegroup.Setup,
		subscription.Setup,
		subscription.Setup,
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
		privatednszonevirtualnetworklink.Setup,
		routetable.Setup,
		securitygroup.Setup,
		subnet.Setup,
		subnetnetworksecuritygroupassociation.Setup,
		subnetroutetableassociation.Setup,
		virtualnetwork.Setup,
		virtualnetworkpeering.Setup,
		providerconfig.Setup,
		mssqldatabase.Setup,
		mssqldatabasevulnerabilityassessmentrulebaseline.Setup,
		mssqlelasticpool.Setup,
		mssqlfailovergroup.Setup,
		mssqlfirewallrule.Setup,
		mssqljobagent.Setup,
		mssqljobcredential.Setup,
		mssqlmanageddatabase.Setup,
		mssqlmanagedinstance.Setup,
		mssqlmanagedinstanceactivedirectoryadministrator.Setup,
		mssqlmanagedinstancefailovergroup.Setup,
		mssqlmanagedinstancevulnerabilityassessment.Setup,
		mssqloutboundfirewallrule.Setup,
		mssqlserver.Setup,
		mssqlserverdnsalias.Setup,
		mssqlservermicrosoftsupportauditingpolicy.Setup,
		mssqlserversecurityalertpolicy.Setup,
		mssqlservertransparentdataencryption.Setup,
		mssqlservervulnerabilityassessment.Setup,
		mssqlvirtualnetworkrule.Setup,
		accountstorage.Setup,
		container.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
