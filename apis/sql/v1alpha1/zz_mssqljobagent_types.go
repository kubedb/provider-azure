// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type MSSQLJobAgentInitParameters struct {

	// The Azure Region where the Elastic Job Agent should exist. Changing this forces a new Elastic Job Agent to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name which should be used for this Elastic Job Agent. Changing this forces a new Elastic Job Agent to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A mapping of tags which should be assigned to the Database.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MSSQLJobAgentObservation struct {

	// The ID of the database to store metadata for the Elastic Job Agent. Changing this forces a new Elastic Job Agent to be created.
	DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// The ID of the Elastic Job Agent.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Azure Region where the Elastic Job Agent should exist. Changing this forces a new Elastic Job Agent to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name which should be used for this Elastic Job Agent. Changing this forces a new Elastic Job Agent to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A mapping of tags which should be assigned to the Database.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MSSQLJobAgentParameters struct {

	// The ID of the database to store metadata for the Elastic Job Agent. Changing this forces a new Elastic Job Agent to be created.
	// +crossplane:generate:reference:type=kubedb.dev/provider-azure/apis/sql/v1alpha1.MSSQLDatabase
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DatabaseID *string `json:"databaseId,omitempty" tf:"database_id,omitempty"`

	// Reference to a MSSQLDatabase in sql to populate databaseId.
	// +kubebuilder:validation:Optional
	DatabaseIDRef *v1.Reference `json:"databaseIdRef,omitempty" tf:"-"`

	// Selector for a MSSQLDatabase in sql to populate databaseId.
	// +kubebuilder:validation:Optional
	DatabaseIDSelector *v1.Selector `json:"databaseIdSelector,omitempty" tf:"-"`

	// The Azure Region where the Elastic Job Agent should exist. Changing this forces a new Elastic Job Agent to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name which should be used for this Elastic Job Agent. Changing this forces a new Elastic Job Agent to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A mapping of tags which should be assigned to the Database.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// MSSQLJobAgentSpec defines the desired state of MSSQLJobAgent
type MSSQLJobAgentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLJobAgentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MSSQLJobAgentInitParameters `json:"initProvider,omitempty"`
}

// MSSQLJobAgentStatus defines the observed state of MSSQLJobAgent.
type MSSQLJobAgentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLJobAgentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLJobAgent is the Schema for the MSSQLJobAgents API. Manages an Elastic Job Agent.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLJobAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MSSQLJobAgentSpec   `json:"spec"`
	Status MSSQLJobAgentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLJobAgentList contains a list of MSSQLJobAgents
type MSSQLJobAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLJobAgent `json:"items"`
}

// Repository type metadata.
var (
	MSSQLJobAgent_Kind             = "MSSQLJobAgent"
	MSSQLJobAgent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLJobAgent_Kind}.String()
	MSSQLJobAgent_KindAPIVersion   = MSSQLJobAgent_Kind + "." + CRDGroupVersion.String()
	MSSQLJobAgent_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLJobAgent_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLJobAgent{}, &MSSQLJobAgentList{})
}