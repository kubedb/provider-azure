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

type RedisFirewallRuleObservation struct {

	// The highest IP address included in the range.
	EndIP *string `json:"endIp,omitempty" tf:"end_ip,omitempty"`

	// The ID of the Redis Firewall Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Redis Cache. Changing this forces a new resource to be created.
	RedisCacheName *string `json:"redisCacheName,omitempty" tf:"redis_cache_name,omitempty"`

	// The name of the resource group in which this Redis Cache exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The lowest IP address included in the range
	StartIP *string `json:"startIp,omitempty" tf:"start_ip,omitempty"`
}

type RedisFirewallRuleParameters struct {

	// The highest IP address included in the range.
	// +kubebuilder:validation:Optional
	EndIP *string `json:"endIp,omitempty" tf:"end_ip,omitempty"`

	// The name of the Redis Cache. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	RedisCacheName *string `json:"redisCacheName" tf:"redis_cache_name,omitempty"`

	// The name of the resource group in which this Redis Cache exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// The lowest IP address included in the range
	// +kubebuilder:validation:Optional
	StartIP *string `json:"startIp,omitempty" tf:"start_ip,omitempty"`
}

// RedisFirewallRuleSpec defines the desired state of RedisFirewallRule
type RedisFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisFirewallRuleParameters `json:"forProvider"`
}

// RedisFirewallRuleStatus defines the observed state of RedisFirewallRule.
type RedisFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisFirewallRule is the Schema for the RedisFirewallRules API. Manages a Firewall Rule associated with a Redis Cache.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.endIp)",message="endIp is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.startIp)",message="startIp is a required parameter"
	Spec   RedisFirewallRuleSpec   `json:"spec"`
	Status RedisFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisFirewallRuleList contains a list of RedisFirewallRules
type RedisFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	RedisFirewallRule_Kind             = "RedisFirewallRule"
	RedisFirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisFirewallRule_Kind}.String()
	RedisFirewallRule_KindAPIVersion   = RedisFirewallRule_Kind + "." + CRDGroupVersion.String()
	RedisFirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(RedisFirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisFirewallRule{}, &RedisFirewallRuleList{})
}
