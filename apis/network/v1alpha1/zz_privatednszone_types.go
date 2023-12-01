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

type PrivateDNSZoneInitParameters struct {

	// An soa_record block as defined below. Changing this forces a new resource to be created.
	SoaRecord []SoaRecordInitParameters `json:"soaRecord,omitempty" tf:"soa_record,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSZoneObservation struct {

	// The Private DNS Zone ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The maximum number of record sets that can be created in this Private DNS zone.
	MaxNumberOfRecordSets *float64 `json:"maxNumberOfRecordSets,omitempty" tf:"max_number_of_record_sets,omitempty"`

	// The maximum number of virtual networks that can be linked to this Private DNS zone.
	MaxNumberOfVirtualNetworkLinks *float64 `json:"maxNumberOfVirtualNetworkLinks,omitempty" tf:"max_number_of_virtual_network_links,omitempty"`

	// The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled.
	MaxNumberOfVirtualNetworkLinksWithRegistration *float64 `json:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty" tf:"max_number_of_virtual_network_links_with_registration,omitempty"`

	// The current number of record sets in this Private DNS zone.
	NumberOfRecordSets *float64 `json:"numberOfRecordSets,omitempty" tf:"number_of_record_sets,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// An soa_record block as defined below. Changing this forces a new resource to be created.
	SoaRecord []SoaRecordObservation `json:"soaRecord,omitempty" tf:"soa_record,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSZoneParameters struct {

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=kubedb.dev/provider-azure/apis/azure/v1alpha1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// An soa_record block as defined below. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SoaRecord []SoaRecordParameters `json:"soaRecord,omitempty" tf:"soa_record,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SoaRecordInitParameters struct {

	// The email contact for the SOA record.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The expire time for the SOA record. Defaults to 2419200.
	ExpireTime *float64 `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to 10.
	MinimumTTL *float64 `json:"minimumTtl,omitempty" tf:"minimum_ttl,omitempty"`

	// The refresh time for the SOA record. Defaults to 3600.
	RefreshTime *float64 `json:"refreshTime,omitempty" tf:"refresh_time,omitempty"`

	// The retry time for the SOA record. Defaults to 300.
	RetryTime *float64 `json:"retryTime,omitempty" tf:"retry_time,omitempty"`

	// The Time To Live of the SOA Record in seconds. Defaults to 3600.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the Record Set.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SoaRecordObservation struct {

	// The email contact for the SOA record.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The expire time for the SOA record. Defaults to 2419200.
	ExpireTime *float64 `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// The fully qualified domain name of the Record Set.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The domain name of the authoritative name server for the SOA record.
	HostName *string `json:"hostName,omitempty" tf:"host_name,omitempty"`

	// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to 10.
	MinimumTTL *float64 `json:"minimumTtl,omitempty" tf:"minimum_ttl,omitempty"`

	// The refresh time for the SOA record. Defaults to 3600.
	RefreshTime *float64 `json:"refreshTime,omitempty" tf:"refresh_time,omitempty"`

	// The retry time for the SOA record. Defaults to 300.
	RetryTime *float64 `json:"retryTime,omitempty" tf:"retry_time,omitempty"`

	// The serial number for the SOA record.
	SerialNumber *float64 `json:"serialNumber,omitempty" tf:"serial_number,omitempty"`

	// The Time To Live of the SOA Record in seconds. Defaults to 3600.
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the Record Set.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SoaRecordParameters struct {

	// The email contact for the SOA record.
	// +kubebuilder:validation:Optional
	Email *string `json:"email" tf:"email,omitempty"`

	// The expire time for the SOA record. Defaults to 2419200.
	// +kubebuilder:validation:Optional
	ExpireTime *float64 `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// The minimum Time To Live for the SOA record. By convention, it is used to determine the negative caching duration. Defaults to 10.
	// +kubebuilder:validation:Optional
	MinimumTTL *float64 `json:"minimumTtl,omitempty" tf:"minimum_ttl,omitempty"`

	// The refresh time for the SOA record. Defaults to 3600.
	// +kubebuilder:validation:Optional
	RefreshTime *float64 `json:"refreshTime,omitempty" tf:"refresh_time,omitempty"`

	// The retry time for the SOA record. Defaults to 300.
	// +kubebuilder:validation:Optional
	RetryTime *float64 `json:"retryTime,omitempty" tf:"retry_time,omitempty"`

	// The Time To Live of the SOA Record in seconds. Defaults to 3600.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the Record Set.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PrivateDNSZoneSpec defines the desired state of PrivateDNSZone
type PrivateDNSZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSZoneParameters `json:"forProvider"`
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
	InitProvider PrivateDNSZoneInitParameters `json:"initProvider,omitempty"`
}

// PrivateDNSZoneStatus defines the observed state of PrivateDNSZone.
type PrivateDNSZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSZone is the Schema for the PrivateDNSZones API. Manages a Private DNS Zone.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSZoneSpec   `json:"spec"`
	Status            PrivateDNSZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSZoneList contains a list of PrivateDNSZones
type PrivateDNSZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSZone `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSZone_Kind             = "PrivateDNSZone"
	PrivateDNSZone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSZone_Kind}.String()
	PrivateDNSZone_KindAPIVersion   = PrivateDNSZone_Kind + "." + CRDGroupVersion.String()
	PrivateDNSZone_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSZone_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSZone{}, &PrivateDNSZoneList{})
}
