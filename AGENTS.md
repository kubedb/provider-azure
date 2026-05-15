# AGENTS.md - KubeDB provider-azure

This file provides instructions for AI coding agents working in this Crossplane provider repository.

## Project Overview

`provider-azure` is a Crossplane provider for Microsoft Azure, generated using [Upjet](https://github.com/crossplane/upjet) from the Terraform `hashicorp/azurerm` provider (v3.61.0). It exposes XRM-conformant Managed Resources backed by Azure services for use by KubeDB and other Crossplane consumers.

- Module: `kubedb.dev/provider-azure`
- Go: `1.22.0` (toolchain `go1.22.2`); Makefile `GO_REQUIRED_VERSION = 1.25`
- Provider root API group: `azure.kubedb.com`
- Terraform provider source: `hashicorp/azurerm` v3.61.0 (Terraform v1.3.3)
- Resource prefix: `azure` (see `config/provider.go`)
- 89 CRDs ship under `package/crds/`

## Build & Development Commands

```bash
# Re-run the upjet code generation pipeline (regenerates apis/, internal/controller/, package/crds/)
go run cmd/generator/main.go "$PWD"

# Build provider and generator binaries
make build

# Run controller locally out-of-cluster against current kubeconfig (with --debug)
make run

# Full build, push image, install xpkg
make all

# Generate Cobertura coverage report (excludes zz_ files)
make cobertura

# Initialize git submodules (required before first make)
make submodules
```

### End-to-end testing

```bash
# Spin up controlplane, deploy locally-built provider package, run uptest
make e2e

# Run only uptest against an already-deployed provider
make uptest UPTEST_EXAMPLE_LIST="examples/dbformysql/..."
```

### Terraform schema regeneration

```bash
# Regenerates config/schema.json from the Terraform provider
make config/schema.json

# Pull Terraform provider docs (sparse-checkout of website/docs/r)
make pull-docs
```

## Project Structure

```
apis/                       # Generated Go types for CRDs (group-versioned)
  authorization/v1alpha1/   # azurerm role assignments
  azure/v1alpha1/           # "root" group: ResourceGroup, Subscription, ProviderRegistration
  cache/v1alpha1/           # Redis Cache, Redis Enterprise
  cosmosdb/v1alpha1/        # Cosmos DB accounts, databases, containers
  dbformariadb/v1alpha1/    # Azure Database for MariaDB
  dbformysql/v1alpha1/      # Azure Database for MySQL (single + flexible server)
  dbforpostgresql/v1alpha1/ # Azure Database for PostgreSQL (single + flexible)
  keyvault/v1alpha1/        # Key Vault, keys, secrets, certificates
  network/v1alpha1/         # VNet, subnets, private endpoints
  sql/v1alpha1/             # Azure SQL servers and databases
  storage/v1alpha1/         # Storage accounts, containers, blobs
  rconfig/                  # Cross-resource reference helpers
  v1alpha1/                 # StoreConfig (external secret stores)
  v1beta1/                  # ProviderConfig, ProviderConfigUsage
  generate.go               # `go generate` directives for upjet
  register_crd.go, zz_register.go  # SchemeBuilder aggregation
cmd/
  provider/main.go              # Controller manager entry point
  generator/main.go             # Runs upjet pipeline + dynamic-controller generator
  dynamic-controller/           # Custom generator emitting zz_dynamic_crd_controller.go
config/
  provider.go              # Upjet Provider configuration (groups, root group, refs)
  external_name.go         # External-name extractors (e.g. Key Vault URL parser)
  overrides.go             # Per-resource customisations
  schema.json              # Captured Terraform provider schema (embedded)
  provider-metadata.yaml   # Terraform docs metadata (embedded)
  {group}/config.go        # Per-group resource configuration (e.g. dbformysql, sql)
internal/
  clients/azure.go         # TerraformSetupBuilder; builds tfproviderazurerm config
  controller/              # Generated per-resource controllers + zz_setup.go
  controller/zz_dynamic_crd_controller.go  # CustomResourceReconciler entrypoint
  features/features.go     # ESS + ManagementPolicies feature flags
package/
  crossplane.yaml          # Crossplane provider package manifest
  crds/                    # 89 generated CRDs
examples/                  # Hand-written example manifests grouped by service
examples-generated/        # Examples generated from Terraform docs
cluster/
  images/provider-azure/   # Dockerfile + terraformrc.hcl baked into image
  test/setup.sh            # uptest setup script
hack/
  boilerplate.go.txt       # License header injected by generator
  prepare.sh               # One-shot template bootstrap (already executed)
build/                     # crossplane/build submodule (Makefiles)
```

## Key Packages / APIs

- `cmd/provider/main.go` - Wires `controller-runtime` manager, registers all schemes via `apis.AddToScheme`, configures `tjcontroller.Options{Provider: config.GetProvider(), SetupFn: clients.TerraformSetupBuilder(...)}`, then calls `controller.NewCustomResourceReconciler(mgr, o).SetupWithManager(mgr)`.
- `config/provider.go` - `GetProvider()` constructs the `ujconfig.Provider` (`resourcePrefix = "azure"`, `WithRootGroup("azure.kubedb.com")`, `WithFeaturesPackage("internal/features")`), assembling per-service configs from `config/{authorization,base,cache,cosmosdb,dbformariadb,dbformysql,dbforpostgresql,keyvault,network,sql,storage}`.
- `internal/clients/azure.go` - `TerraformSetupBuilder` reads the `ProviderConfig`, extracts Azure service principal credentials (`subscriptionId`, `clientId`, `clientSecret`, `tenantId`) and emits a Terraform setup with `features = {}` and `skip_provider_registration`.
- `cmd/dynamic-controller/setup.go` - Custom generator that walks the upjet `PackageMonolith` output and writes a dynamic reconciler (`zz_dynamic_crd_controller.go`) capable of watching CRDs that appear at runtime.
- `internal/features/features.go` - Flags `EnableAlphaExternalSecretStores` and `EnableBetaManagementPolicies`, gated by `--enable-external-secret-stores` / `--enable-management-policies` CLI flags.
- `apis/v1beta1` - `ProviderConfig` (Azure credentials reference) and `ProviderConfigUsage` definitions.
- `apis/v1alpha1` - `StoreConfig` for External Secret Stores alpha feature.

### Provider CLI flags (`cmd/provider/main.go`)

```
--debug, -d                       Run with debug logging
--sync, -s                        Controller manager sync period (default 1h)
--leader-election, -l             Enable leader election (LEADER_ELECTION env)
--terraform-version               Required (TERRAFORM_VERSION env)
--terraform-provider-source       Required (TERRAFORM_PROVIDER_SOURCE env)
--terraform-provider-version      Required (TERRAFORM_PROVIDER_VERSION env)
--max-reconcile-rate              Global rate per second (default 10)
--namespace                       Default secret store scope (POD_NAMESPACE env)
--enable-external-secret-stores   Alpha ESS support
--enable-management-policies      Beta Management Policies support
```

## Testing

- No first-party unit tests are committed; generated `zz_*.go` files are excluded from coverage via `make cobertura`.
- End-to-end testing uses [uptest](https://github.com/crossplane/uptest) (`UPTEST_VERSION = v0.2.1`) plus KUTTL: `make e2e` or `make uptest UPTEST_EXAMPLE_LIST=...`.
- Examples used as uptest cases live under `examples/` (hand-curated) and `examples-generated/` (auto-derived from Terraform docs).
- Local cluster spun up via `make local-deploy`, which uses `build/makelib/controlplane.mk` and KIND `v0.15.0`.

## Dependencies

| Dependency | Version | Purpose |
|------------|---------|---------|
| `github.com/crossplane/crossplane-runtime` | v1.14.0-rc.0 | Managed resource framework |
| `github.com/crossplane/upjet` | v1.0.0 | Terraform-to-Crossplane code generation |
| `github.com/crossplane/crossplane-tools` | latest | CRD/managed code generators |
| `github.com/hashicorp/terraform-plugin-sdk/v2` | v2.24.0 | Terraform schema and runtime |
| `sigs.k8s.io/controller-runtime` | v0.17.2 | Kubernetes controller framework |
| `sigs.k8s.io/controller-tools` | v0.13.0 | CRD manifest generation |
| `k8s.io/{api,apimachinery,client-go}` | v0.29.2 | Kubernetes client libraries |
| `gopkg.in/alecthomas/kingpin.v2` | v2.2.6 | CLI flag parsing |
| `github.com/muvaf/typewriter` | unstable | Code-generation writer used by upjet |

### Runtime tooling (Makefile)

| Tool | Version | Notes |
|------|---------|-------|
| Terraform | 1.3.3 | Downloaded by Makefile + baked into image |
| terraform-provider-azurerm | 3.61.0 | Native plugin baked into image |
| KIND | v0.15.0 | Local controlplane |
| `up` CLI | v0.14.0 | Crossplane package management |
| uptest | v0.2.1 | E2E harness |
| Container base | `alpine:3.17.1` | See `cluster/images/provider-azure/Dockerfile` |

## Code Conventions

- All files prefixed `zz_` are generated; never hand-edit. Regenerate via `go run cmd/generator/main.go "$PWD"`.
- License header for generated code lives in `hack/boilerplate.go.txt`.
- Per-service customizations belong in `config/{group}/config.go` and are wired into `config.GetProvider()`; do not modify generated `apis/{group}/...` files directly.
- External-name extraction logic for Azure resource IDs lives in `config/external_name.go` (e.g. `keyVaultURLIDConf` parses `https://{vault}.vault.azure.net/...` URLs).
- The "root" group `azure.kubedb.com` (registered via `WithRootGroup`) holds resources without a natural Terraform group (e.g. `ResourceGroup`, `Subscription`, `ProviderRegistration`).
- Image registries: provider images are published to both `ghcr.io/kubedb` and `xpkg.upbound.io/upbound`.
- Submodule `build/` (crossplane/build) provides the Makefile layers (`common.mk`, `golang.mk`, `k8s_tools.mk`, `image.mk`, etc.); update it with `make submodules`.
- Bumping the Azure Terraform provider: edit `TERRAFORM_PROVIDER_VERSION` and `TERRAFORM_NATIVE_PROVIDER_BINARY` in the Makefile, then run `make config/schema.json pull-docs` followed by the generator.
