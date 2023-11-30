/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/terraform"

	"kubedb.dev/provider-azure/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal azure credentials as JSON"

	// Azure service principal credentials file JSON keys
	keyAzureSubscriptionID = "subscriptionId"
	keyAzureClientID       = "clientId"
	keyAzureClientSecret   = "clientSecret"
	keyAzureTenantID       = "tenantId"
	// Terraform Provider configuration block keys
	keyClientID       = "client_id"
	keySubscriptionID = "subscription_id"
	keyTenantID       = "tenant_id"
	keyClientSecret   = "client_secret"

	keyTerraformFeatures        = "features"
	keySkipProviderRegistration = "skip_provider_registration"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]interface{}{
			keyTerraformFeatures: struct{}{},
			// Terraform AzureRM provider tries to register all resource providers
			// in Azure just in case if the provider of the resource you're
			// trying to create is not registered and the returned error is
			// ambiguous. However, this requires service principal to have provider
			// registration permissions which are irrelevant in most contexts.
			// For details, see https://github.com/upbound/provider-azure/issues/104
			keySkipProviderRegistration: true,
		}
		// using spAuth for azure authentication
		if _, ok := creds[keyAzureSubscriptionID]; ok {
			ps.Configuration[keySubscriptionID] = creds[keyAzureSubscriptionID]
		} else {
			return ps, errors.New("invalid subscription id")
		}
		if _, ok := creds[keyAzureTenantID]; ok {
			ps.Configuration[keyTenantID] = creds[keyAzureTenantID]
		} else {
			return ps, errors.New("invalid tenant id")
		}
		if _, ok := creds[keyAzureClientID]; ok {
			ps.Configuration[keyClientID] = creds[keyAzureClientID]
		} else {
			return ps, errors.New("invalid client id")
		}
		if _, ok := creds[keyAzureClientSecret]; ok {
			ps.Configuration[keyClientSecret] = creds[keyAzureClientSecret]
		} else {
			return ps, errors.New("invalid client secret")
		}
		return ps, nil
	}
}
