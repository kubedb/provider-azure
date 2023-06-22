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

type VirtualNetworkRuleObservation struct {

	// The ID of the MySQL Virtual Network Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// The ID of the subnet that the MySQL server will be connected to.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type VirtualNetworkRuleParameters struct {

	// The name of the resource group where the MySQL server resides. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// The name of the SQL Server to which this MySQL virtual network rule will be applied to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ServerName *string `json:"serverName" tf:"server_name,omitempty"`

	// The ID of the subnet that the MySQL server will be connected to.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

// VirtualNetworkRuleSpec defines the desired state of VirtualNetworkRule
type VirtualNetworkRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkRuleParameters `json:"forProvider"`
}

// VirtualNetworkRuleStatus defines the observed state of VirtualNetworkRule.
type VirtualNetworkRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkRule is the Schema for the VirtualNetworkRules API. Manages a MySQL Virtual Network Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.subnetId)",message="subnetId is a required parameter"
	Spec   VirtualNetworkRuleSpec   `json:"spec"`
	Status VirtualNetworkRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkRuleList contains a list of VirtualNetworkRules
type VirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkRule `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkRule_Kind             = "VirtualNetworkRule"
	VirtualNetworkRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetworkRule_Kind}.String()
	VirtualNetworkRule_KindAPIVersion   = VirtualNetworkRule_Kind + "." + CRDGroupVersion.String()
	VirtualNetworkRule_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetworkRule_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkRule{}, &VirtualNetworkRuleList{})
}
