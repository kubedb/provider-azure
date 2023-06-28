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

type SQLRoleAssignmentObservation struct {

	// The name of the Cosmos DB Account. Changing this forces a new resource to be created.
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// The ID of the Cosmos DB SQL Role Assignment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The GUID as the name of the Cosmos DB SQL Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Principal (Client) in Azure Active Directory. Changing this forces a new resource to be created.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The name of the Resource Group in which the Cosmos DB SQL Role Assignment is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The resource ID of the Cosmos DB SQL Role Definition.
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// The data plane resource path for which access is being granted through this Cosmos DB SQL Role Assignment. Changing this forces a new resource to be created.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

type SQLRoleAssignmentParameters struct {

	// The name of the Cosmos DB Account. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// The GUID as the name of the Cosmos DB SQL Role Assignment - one will be generated if not specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Principal (Client) in Azure Active Directory. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The name of the Resource Group in which the Cosmos DB SQL Role Assignment is created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The resource ID of the Cosmos DB SQL Role Definition.
	// +kubebuilder:validation:Optional
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// The data plane resource path for which access is being granted through this Cosmos DB SQL Role Assignment. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`
}

// SQLRoleAssignmentSpec defines the desired state of SQLRoleAssignment
type SQLRoleAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLRoleAssignmentParameters `json:"forProvider"`
}

// SQLRoleAssignmentStatus defines the observed state of SQLRoleAssignment.
type SQLRoleAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLRoleAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLRoleAssignment is the Schema for the SQLRoleAssignments API. Manages a Cosmos DB SQL Role Assignment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLRoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.accountName)",message="accountName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.principalId)",message="principalId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.resourceGroupName)",message="resourceGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.roleDefinitionId)",message="roleDefinitionId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.scope)",message="scope is a required parameter"
	Spec   SQLRoleAssignmentSpec   `json:"spec"`
	Status SQLRoleAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLRoleAssignmentList contains a list of SQLRoleAssignments
type SQLRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLRoleAssignment `json:"items"`
}

// Repository type metadata.
var (
	SQLRoleAssignment_Kind             = "SQLRoleAssignment"
	SQLRoleAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLRoleAssignment_Kind}.String()
	SQLRoleAssignment_KindAPIVersion   = SQLRoleAssignment_Kind + "." + CRDGroupVersion.String()
	SQLRoleAssignment_GroupVersionKind = CRDGroupVersion.WithKind(SQLRoleAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLRoleAssignment{}, &SQLRoleAssignmentList{})
}