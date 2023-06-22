/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "kubeform.dev/provider-azure/apis/cache/v1alpha1"
	v1alpha1cosmosdb "kubeform.dev/provider-azure/apis/cosmosdb/v1alpha1"
	v1alpha1dbformariadb "kubeform.dev/provider-azure/apis/dbformariadb/v1alpha1"
	v1alpha1dbformysql "kubeform.dev/provider-azure/apis/dbformysql/v1alpha1"
	v1alpha1dbforpostgresql "kubeform.dev/provider-azure/apis/dbforpostgresql/v1alpha1"
	v1alpha1keyvault "kubeform.dev/provider-azure/apis/keyvault/v1alpha1"
	v1alpha1network "kubeform.dev/provider-azure/apis/network/v1alpha1"
	v1alpha1apis "kubeform.dev/provider-azure/apis/v1alpha1"
	v1beta1 "kubeform.dev/provider-azure/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1cosmosdb.SchemeBuilder.AddToScheme,
		v1alpha1dbformariadb.SchemeBuilder.AddToScheme,
		v1alpha1dbformysql.SchemeBuilder.AddToScheme,
		v1alpha1dbforpostgresql.SchemeBuilder.AddToScheme,
		v1alpha1keyvault.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
