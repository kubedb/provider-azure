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




type FeatureInitParameters struct {


// Specifies the name of the feature to register.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Should this feature be Registered or Unregistered?
Registered *bool `json:"registered,omitempty" tf:"registered,omitempty"`
}


type FeatureObservation struct {


// Specifies the name of the feature to register.
Name *string `json:"name,omitempty" tf:"name,omitempty"`

// Should this feature be Registered or Unregistered?
Registered *bool `json:"registered,omitempty" tf:"registered,omitempty"`
}


type FeatureParameters struct {


// Specifies the name of the feature to register.
// +kubebuilder:validation:Optional
Name *string `json:"name" tf:"name,omitempty"`

// Should this feature be Registered or Unregistered?
// +kubebuilder:validation:Optional
Registered *bool `json:"registered" tf:"registered,omitempty"`
}


type ProviderRegistrationInitParameters struct {


// A list of feature blocks as defined below.
Feature []FeatureInitParameters `json:"feature,omitempty" tf:"feature,omitempty"`

// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
Name *string `json:"name,omitempty" tf:"name,omitempty"`
}


type ProviderRegistrationObservation struct {


// A list of feature blocks as defined below.
Feature []FeatureObservation `json:"feature,omitempty" tf:"feature,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
Name *string `json:"name,omitempty" tf:"name,omitempty"`
}


type ProviderRegistrationParameters struct {


// A list of feature blocks as defined below.
// +kubebuilder:validation:Optional
Feature []FeatureParameters `json:"feature,omitempty" tf:"feature,omitempty"`

// The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.
// +kubebuilder:validation:Optional
Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// ProviderRegistrationSpec defines the desired state of ProviderRegistration
type ProviderRegistrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       ProviderRegistrationParameters `json:"forProvider"`
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
	InitProvider       ProviderRegistrationInitParameters `json:"initProvider,omitempty"`
}

// ProviderRegistrationStatus defines the observed state of ProviderRegistration.
type ProviderRegistrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          ProviderRegistrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderRegistration is the Schema for the ProviderRegistrations API. Manages the Registration of a Resource Provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ProviderRegistration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec              ProviderRegistrationSpec   `json:"spec"`
	Status            ProviderRegistrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderRegistrationList contains a list of ProviderRegistrations
type ProviderRegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProviderRegistration `json:"items"`
}

// Repository type metadata.
var (
	ProviderRegistration_Kind             = "ProviderRegistration"
	ProviderRegistration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProviderRegistration_Kind}.String()
	ProviderRegistration_KindAPIVersion   = ProviderRegistration_Kind + "." + CRDGroupVersion.String()
	ProviderRegistration_GroupVersionKind = CRDGroupVersion.WithKind(ProviderRegistration_Kind)
)

func init() {
	SchemeBuilder.Register(&ProviderRegistration{}, &ProviderRegistrationList{})
}
