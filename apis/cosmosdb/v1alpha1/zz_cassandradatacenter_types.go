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

type CassandraDatacenterObservation struct {

	// Determines whether availability zones are enabled. Defaults to true.
	AvailabilityZonesEnabled *bool `json:"availabilityZonesEnabled,omitempty" tf:"availability_zones_enabled,omitempty"`

	// The key URI of the customer key to use for the encryption of the backup Storage Account.
	BackupStorageCustomerKeyURI *string `json:"backupStorageCustomerKeyUri,omitempty" tf:"backup_storage_customer_key_uri,omitempty"`

	// The fragment of the cassandra.yaml configuration file to be included in the cassandra.yaml for all nodes in this Cassandra Datacenter. The fragment should be Base64 encoded and only a subset of keys is allowed.
	Base64EncodedYamlFragment *string `json:"base64EncodedYamlFragment,omitempty" tf:"base64_encoded_yaml_fragment,omitempty"`

	// The ID of the Cassandra Cluster. Changing this forces a new Cassandra Datacenter to be created.
	CassandraClusterID *string `json:"cassandraClusterId,omitempty" tf:"cassandra_cluster_id,omitempty"`

	// The ID of the delegated management subnet for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	DelegatedManagementSubnetID *string `json:"delegatedManagementSubnetId,omitempty" tf:"delegated_management_subnet_id,omitempty"`

	// Determines the number of p30 disks that are attached to each node.
	DiskCount *float64 `json:"diskCount,omitempty" tf:"disk_count,omitempty"`

	// The Disk SKU that is used for this Cassandra Datacenter. Defaults to P30.
	DiskSku *string `json:"diskSku,omitempty" tf:"disk_sku,omitempty"`

	// The ID of the Cassandra Datacenter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Cassandra Datacenter should exist. Changing this forces a new Cassandra Datacenter to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The key URI of the customer key to use for the encryption of the Managed Disk.
	ManagedDiskCustomerKeyURI *string `json:"managedDiskCustomerKeyUri,omitempty" tf:"managed_disk_customer_key_uri,omitempty"`

	// The number of nodes the Cassandra Datacenter should have. The number should be equal or greater than 3. Defaults to 3.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// Determines the selected sku.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`
}

type CassandraDatacenterParameters struct {

	// Determines whether availability zones are enabled. Defaults to true.
	// +kubebuilder:validation:Optional
	AvailabilityZonesEnabled *bool `json:"availabilityZonesEnabled,omitempty" tf:"availability_zones_enabled,omitempty"`

	// The key URI of the customer key to use for the encryption of the backup Storage Account.
	// +kubebuilder:validation:Optional
	BackupStorageCustomerKeyURI *string `json:"backupStorageCustomerKeyUri,omitempty" tf:"backup_storage_customer_key_uri,omitempty"`

	// The fragment of the cassandra.yaml configuration file to be included in the cassandra.yaml for all nodes in this Cassandra Datacenter. The fragment should be Base64 encoded and only a subset of keys is allowed.
	// +kubebuilder:validation:Optional
	Base64EncodedYamlFragment *string `json:"base64EncodedYamlFragment,omitempty" tf:"base64_encoded_yaml_fragment,omitempty"`

	// The ID of the Cassandra Cluster. Changing this forces a new Cassandra Datacenter to be created.
	// +kubebuilder:validation:Required
	CassandraClusterID *string `json:"cassandraClusterId" tf:"cassandra_cluster_id,omitempty"`

	// The ID of the delegated management subnet for this Cassandra Datacenter. Changing this forces a new Cassandra Datacenter to be created.
	// +kubebuilder:validation:Optional
	DelegatedManagementSubnetID *string `json:"delegatedManagementSubnetId,omitempty" tf:"delegated_management_subnet_id,omitempty"`

	// Determines the number of p30 disks that are attached to each node.
	// +kubebuilder:validation:Optional
	DiskCount *float64 `json:"diskCount,omitempty" tf:"disk_count,omitempty"`

	// The Disk SKU that is used for this Cassandra Datacenter. Defaults to P30.
	// +kubebuilder:validation:Optional
	DiskSku *string `json:"diskSku,omitempty" tf:"disk_sku,omitempty"`

	// The Azure Region where the Cassandra Datacenter should exist. Changing this forces a new Cassandra Datacenter to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The key URI of the customer key to use for the encryption of the Managed Disk.
	// +kubebuilder:validation:Optional
	ManagedDiskCustomerKeyURI *string `json:"managedDiskCustomerKeyUri,omitempty" tf:"managed_disk_customer_key_uri,omitempty"`

	// The number of nodes the Cassandra Datacenter should have. The number should be equal or greater than 3. Defaults to 3.
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// Determines the selected sku.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`
}

// CassandraDatacenterSpec defines the desired state of CassandraDatacenter
type CassandraDatacenterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CassandraDatacenterParameters `json:"forProvider"`
}

// CassandraDatacenterStatus defines the observed state of CassandraDatacenter.
type CassandraDatacenterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CassandraDatacenterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraDatacenter is the Schema for the CassandraDatacenters API. Manages a Cassandra Datacenter.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CassandraDatacenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.delegatedManagementSubnetId)",message="delegatedManagementSubnetId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.location)",message="location is a required parameter"
	Spec   CassandraDatacenterSpec   `json:"spec"`
	Status CassandraDatacenterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraDatacenterList contains a list of CassandraDatacenters
type CassandraDatacenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CassandraDatacenter `json:"items"`
}

// Repository type metadata.
var (
	CassandraDatacenter_Kind             = "CassandraDatacenter"
	CassandraDatacenter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CassandraDatacenter_Kind}.String()
	CassandraDatacenter_KindAPIVersion   = CassandraDatacenter_Kind + "." + CRDGroupVersion.String()
	CassandraDatacenter_GroupVersionKind = CRDGroupVersion.WithKind(CassandraDatacenter_Kind)
)

func init() {
	SchemeBuilder.Register(&CassandraDatacenter{}, &CassandraDatacenterList{})
}
