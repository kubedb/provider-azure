/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ModuleObservation struct {

	// Configuration options for the module (e.g. ERROR_RATE 0.00 INITIAL_SIZE 400). Changing this forces a new resource to be created. Defaults to "".
	Args *string `json:"args,omitempty" tf:"args,omitempty"`

	// The name which should be used for this module. Possible values are RedisBloom, RedisTimeSeries, RediSearch and RedisJSON. Changing this forces a new Redis Enterprise Database to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ModuleParameters struct {

	// Configuration options for the module (e.g. ERROR_RATE 0.00 INITIAL_SIZE 400). Changing this forces a new resource to be created. Defaults to "".
	// +kubebuilder:validation:Optional
	Args *string `json:"args,omitempty" tf:"args,omitempty"`

	// The name which should be used for this module. Possible values are RedisBloom, RedisTimeSeries, RediSearch and RedisJSON. Changing this forces a new Redis Enterprise Database to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type RedisEnterpriseDatabaseObservation struct {

	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are Encrypted and Plaintext. Defaults to Encrypted. Changing this forces a new Redis Enterprise Database to be created.
	ClientProtocol *string `json:"clientProtocol,omitempty" tf:"client_protocol,omitempty"`

	// The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Clustering policy - default is OSSCluster. Specified at create time. Possible values are EnterpriseCluster and OSSCluster. Defaults to OSSCluster. Changing this forces a new Redis Enterprise Database to be created.
	ClusteringPolicy *string `json:"clusteringPolicy,omitempty" tf:"clustering_policy,omitempty"`

	// Redis eviction policy - default is VolatileLRU. Possible values are AllKeysLFU, AllKeysLRU, AllKeysRandom, VolatileLRU, VolatileLFU, VolatileTTL, VolatileRandom and NoEviction. Changing this forces a new Redis Enterprise Database to be created.
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// The ID of the Redis Enterprise Database.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Nickname of the group of linked databases. Changing this force a new Redis Enterprise Geo Database to be created.
	LinkedDatabaseGroupNickname *string `json:"linkedDatabaseGroupNickname,omitempty" tf:"linked_database_group_nickname,omitempty"`

	// A list of database resources to link with this database with a maximum of 5.
	LinkedDatabaseID []*string `json:"linkedDatabaseId,omitempty" tf:"linked_database_id,omitempty"`

	// A module block as defined below. Changing this forces a new resource to be created.
	Module []ModuleObservation `json:"module,omitempty" tf:"module,omitempty"`

	// TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created. Defaults to 10000.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

type RedisEnterpriseDatabaseParameters struct {

	// Specifies whether redis clients can connect using TLS-encrypted or plaintext redis protocols. Default is TLS-encrypted. Possible values are Encrypted and Plaintext. Defaults to Encrypted. Changing this forces a new Redis Enterprise Database to be created.
	// +kubebuilder:validation:Optional
	ClientProtocol *string `json:"clientProtocol,omitempty" tf:"client_protocol,omitempty"`

	// The resource id of the Redis Enterprise Cluster to deploy this Redis Enterprise Database. Changing this forces a new Redis Enterprise Database to be created.
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// Clustering policy - default is OSSCluster. Specified at create time. Possible values are EnterpriseCluster and OSSCluster. Defaults to OSSCluster. Changing this forces a new Redis Enterprise Database to be created.
	// +kubebuilder:validation:Optional
	ClusteringPolicy *string `json:"clusteringPolicy,omitempty" tf:"clustering_policy,omitempty"`

	// Redis eviction policy - default is VolatileLRU. Possible values are AllKeysLFU, AllKeysLRU, AllKeysRandom, VolatileLRU, VolatileLFU, VolatileTTL, VolatileRandom and NoEviction. Changing this forces a new Redis Enterprise Database to be created.
	// +kubebuilder:validation:Optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// Nickname of the group of linked databases. Changing this force a new Redis Enterprise Geo Database to be created.
	// +kubebuilder:validation:Optional
	LinkedDatabaseGroupNickname *string `json:"linkedDatabaseGroupNickname,omitempty" tf:"linked_database_group_nickname,omitempty"`

	// A list of database resources to link with this database with a maximum of 5.
	// +kubebuilder:validation:Optional
	LinkedDatabaseID []*string `json:"linkedDatabaseId,omitempty" tf:"linked_database_id,omitempty"`

	// A module block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Module []ModuleParameters `json:"module,omitempty" tf:"module,omitempty"`

	// TCP port of the database endpoint. Specified at create time. Defaults to an available port. Changing this forces a new Redis Enterprise Database to be created. Defaults to 10000.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The name of the Resource Group where the Redis Enterprise Database should exist. Changing this forces a new Redis Enterprise Database to be created.
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`
}

// RedisEnterpriseDatabaseSpec defines the desired state of RedisEnterpriseDatabase
type RedisEnterpriseDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisEnterpriseDatabaseParameters `json:"forProvider"`
}

// RedisEnterpriseDatabaseStatus defines the observed state of RedisEnterpriseDatabase.
type RedisEnterpriseDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisEnterpriseDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseDatabase is the Schema for the RedisEnterpriseDatabases API. Manages a Redis Enterprise Database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisEnterpriseDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterpriseDatabaseSpec   `json:"spec"`
	Status            RedisEnterpriseDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseDatabaseList contains a list of RedisEnterpriseDatabases
type RedisEnterpriseDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterpriseDatabase `json:"items"`
}

// Repository type metadata.
var (
	RedisEnterpriseDatabase_Kind             = "RedisEnterpriseDatabase"
	RedisEnterpriseDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisEnterpriseDatabase_Kind}.String()
	RedisEnterpriseDatabase_KindAPIVersion   = RedisEnterpriseDatabase_Kind + "." + CRDGroupVersion.String()
	RedisEnterpriseDatabase_GroupVersionKind = CRDGroupVersion.WithKind(RedisEnterpriseDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisEnterpriseDatabase{}, &RedisEnterpriseDatabaseList{})
}
