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

type SQLDedicatedGatewayObservation struct {

	// The resource ID of the CosmosDB Account. Changing this forces a new resource to be created.
	CosmosdbAccountID *string `json:"cosmosdbAccountId,omitempty" tf:"cosmosdb_account_id,omitempty"`

	// The ID of the CosmosDB SQL Dedicated Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance count for the CosmosDB SQL Dedicated Gateway. Possible value is between 1 and 5.
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The instance size for the CosmosDB SQL Dedicated Gateway. Changing this forces a new resource to be created. Possible values are Cosmos.D4s, Cosmos.D8s and Cosmos.D16s.
	InstanceSize *string `json:"instanceSize,omitempty" tf:"instance_size,omitempty"`
}

type SQLDedicatedGatewayParameters struct {

	// The resource ID of the CosmosDB Account. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	CosmosdbAccountID *string `json:"cosmosdbAccountId" tf:"cosmosdb_account_id,omitempty"`

	// The instance count for the CosmosDB SQL Dedicated Gateway. Possible value is between 1 and 5.
	// +kubebuilder:validation:Optional
	InstanceCount *float64 `json:"instanceCount,omitempty" tf:"instance_count,omitempty"`

	// The instance size for the CosmosDB SQL Dedicated Gateway. Changing this forces a new resource to be created. Possible values are Cosmos.D4s, Cosmos.D8s and Cosmos.D16s.
	// +kubebuilder:validation:Optional
	InstanceSize *string `json:"instanceSize,omitempty" tf:"instance_size,omitempty"`
}

// SQLDedicatedGatewaySpec defines the desired state of SQLDedicatedGateway
type SQLDedicatedGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLDedicatedGatewayParameters `json:"forProvider"`
}

// SQLDedicatedGatewayStatus defines the observed state of SQLDedicatedGateway.
type SQLDedicatedGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLDedicatedGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLDedicatedGateway is the Schema for the SQLDedicatedGateways API. Manages a SQL Dedicated Gateway within a Cosmos DB Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLDedicatedGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.instanceCount)",message="instanceCount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.instanceSize)",message="instanceSize is a required parameter"
	Spec   SQLDedicatedGatewaySpec   `json:"spec"`
	Status SQLDedicatedGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLDedicatedGatewayList contains a list of SQLDedicatedGateways
type SQLDedicatedGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLDedicatedGateway `json:"items"`
}

// Repository type metadata.
var (
	SQLDedicatedGateway_Kind             = "SQLDedicatedGateway"
	SQLDedicatedGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLDedicatedGateway_Kind}.String()
	SQLDedicatedGateway_KindAPIVersion   = SQLDedicatedGateway_Kind + "." + CRDGroupVersion.String()
	SQLDedicatedGateway_GroupVersionKind = CRDGroupVersion.WithKind(SQLDedicatedGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLDedicatedGateway{}, &SQLDedicatedGatewayList{})
}