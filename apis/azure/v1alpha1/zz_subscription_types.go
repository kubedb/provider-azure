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

type SubscriptionObservation struct {

	// The Azure Billing Scope ID. Can be a Microsoft Customer Account Billing Scope ID, a Microsoft Partner Account Billing Scope ID or an Enrollment Billing Scope ID.
	BillingScopeID *string `json:"billingScopeId,omitempty" tf:"billing_scope_id,omitempty"`

	// The Resource ID of the Alias.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Subscription. Changing this forces a new Subscription to be created.
	// The GUID of the Subscription.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The Name of the Subscription. This is the Display Name in the portal.
	// The Display Name for the Subscription.
	SubscriptionName *string `json:"subscriptionName,omitempty" tf:"subscription_name,omitempty"`

	// A mapping of tags to assign to the Subscription.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Tenant to which the subscription belongs.
	// The Tenant ID to which the subscription belongs
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// The workload type of the Subscription. Possible values are Production (default) and DevTest. Changing this forces a new Subscription to be created.
	// The workload type for the Subscription. Possible values are `Production` (default) and `DevTest`.
	Workload *string `json:"workload,omitempty" tf:"workload,omitempty"`
}

type SubscriptionParameters struct {

	// The Azure Billing Scope ID. Can be a Microsoft Customer Account Billing Scope ID, a Microsoft Partner Account Billing Scope ID or an Enrollment Billing Scope ID.
	// +kubebuilder:validation:Optional
	BillingScopeID *string `json:"billingScopeId,omitempty" tf:"billing_scope_id,omitempty"`

	// The ID of the Subscription. Changing this forces a new Subscription to be created.
	// The GUID of the Subscription.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The Name of the Subscription. This is the Display Name in the portal.
	// The Display Name for the Subscription.
	// +kubebuilder:validation:Optional
	SubscriptionName *string `json:"subscriptionName,omitempty" tf:"subscription_name,omitempty"`

	// A mapping of tags to assign to the Subscription.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The workload type of the Subscription. Possible values are Production (default) and DevTest. Changing this forces a new Subscription to be created.
	// The workload type for the Subscription. Possible values are `Production` (default) and `DevTest`.
	// +kubebuilder:validation:Optional
	Workload *string `json:"workload,omitempty" tf:"workload,omitempty"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subscription is the Schema for the Subscriptions API. Manages a Subscription.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.subscriptionName)",message="subscriptionName is a required parameter"
	Spec   SubscriptionSpec   `json:"spec"`
	Status SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}