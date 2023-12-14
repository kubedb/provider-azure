/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"fmt"
	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/examples"
	"github.com/crossplane/upjet/pkg/pipeline"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	pconfig "kubedb.dev/provider-azure/config"
)

type TerraformedInput struct {
	*config.Resource
	ParametersTypeName string
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	pipeline.Run(pconfig.GetProvider(), absRootDir)

	pc := pconfig.GetProvider()
	/*fmt.Println(pc.ModulePath, " !!!! ", pc.MainTemplate)
	i := 0
	for p, g := range pc.BasePackages.ControllerMap {
		fmt.Println(i, " ", p, " ", g, "", config.PackageNameMonolith)
		i += 1
	}
	j := 0
	for st, re := range pc.Resources {
		group := pc.RootGroup

		if re.ShortGroup != "" {
			group = strings.ToLower(re.ShortGroup) + "." + pc.RootGroup
		}
		sGroup := strings.Split(group, ".")[0]
		fmt.Println(j, " ", st, " ", group, " ", sGroup)
		j += 1
	}
	cnt := 0
	for _, p := range pc.BasePackages.Controller {
		path := filepath.Join(pc.ModulePath, p)
		fmt.Println(cnt, " ", path)
		cnt += 1
	}*/

	resourcesGroups := map[string]map[string]map[string]*config.Resource{}
	for name, resource := range pc.Resources {
		group := pc.RootGroup
		if resource.ShortGroup != "" {
			group = strings.ToLower(resource.ShortGroup) + "." + pc.RootGroup
		}
		if len(resourcesGroups[group]) == 0 {
			resourcesGroups[group] = map[string]map[string]*config.Resource{}
		}
		if len(resourcesGroups[group][resource.Version]) == 0 {
			resourcesGroups[group][resource.Version] = map[string]*config.Resource{}
		}
		resourcesGroups[group][resource.Version][name] = resource
	}

	exampleGen := examples.NewGenerator(rootDir, pc.ModulePath, pc.ShortName, pc.Resources)
	if err := exampleGen.SetReferenceTypes(pc.Resources); err != nil {
		panic(errors.Wrap(err, "cannot set reference types for resources"))
	}
	// Add ProviderConfig API package to the list of API version packages.
	apiVersionPkgList := make([]string, 0)
	for _, p := range pc.BasePackages.APIVersion {
		apiVersionPkgList = append(apiVersionPkgList, filepath.Join(pc.ModulePath, p))
	}
	// Add ProviderConfig controller package to the list of controller packages.
	controllerPkgMap := make(map[string][]string)
	// new API takes precedence
	for p, g := range pc.BasePackages.ControllerMap {
		path := filepath.Join(pc.ModulePath, p)
		controllerPkgMap[g] = append(controllerPkgMap[g], path)
		controllerPkgMap[config.PackageNameMonolith] = append(controllerPkgMap[config.PackageNameMonolith], path)
	}
	//nolint:staticcheck
	for _, p := range pc.BasePackages.Controller {
		path := filepath.Join(pc.ModulePath, p)
		found := false
		for _, p := range controllerPkgMap[config.PackageNameConfig] {
			if path == p {
				found = true
				break
			}
		}
		if !found {
			controllerPkgMap[config.PackageNameConfig] = append(controllerPkgMap[config.PackageNameConfig], path)
		}
		found = false
		for _, p := range controllerPkgMap[config.PackageNameMonolith] {
			if path == p {
				found = true
				break
			}
		}
		if !found {
			controllerPkgMap[config.PackageNameMonolith] = append(controllerPkgMap[config.PackageNameMonolith], path)
		}
	}
	count := 0
	for group, versions := range resourcesGroups {
		for version, resources := range versions {
			versionGen := pipeline.NewVersionGenerator(rootDir, pc.ModulePath, group, version)
			crdGen := pipeline.NewCRDGenerator(versionGen.Package(), rootDir, pc.ShortName, group, version)
			var _ = pipeline.NewTerraformedGenerator(versionGen.Package(), rootDir, group, version)
			ctrlGen := pipeline.NewControllerGenerator(rootDir, pc.ModulePath, group)

			for _, name := range sortedResources(resources) {
				_, err := crdGen.Generate(resources[name])
				if err != nil {
					panic(errors.Wrapf(err, "cannot generate crd for resource %s", name))
				}

				featuresPkgPath := ""
				if pc.FeaturesPackage != "" {
					featuresPkgPath = filepath.Join(pc.ModulePath, pc.FeaturesPackage)
				}
				ctrlPkgPath, err := ctrlGen.Generate(resources[name], versionGen.Package().Path(), featuresPkgPath)
				if err != nil {
					panic(errors.Wrapf(err, "cannot generate controller for resource %s", name))
				}
				sGroup := strings.Split(group, ".")[0]
				controllerPkgMap[sGroup] = append(controllerPkgMap[sGroup], ctrlPkgPath)
				controllerPkgMap[config.PackageNameMonolith] = append(controllerPkgMap[config.PackageNameMonolith], ctrlPkgPath)
				if err := exampleGen.Generate(group, version, resources[name]); err != nil {
					panic(errors.Wrapf(err, "cannot generate example manifest for resource %s", name))
				}
				count++
			}

			if err := versionGen.Generate(); err != nil {
				panic(errors.Wrap(err, "cannot generate version files"))
			}
			apiVersionPkgList = append(apiVersionPkgList, versionGen.Package().Path())
		}
	}

	if err := exampleGen.StoreExamples(); err != nil {
		panic(errors.Wrapf(err, "cannot store examples"))
	}

	if err := pipeline.NewRegisterGenerator(rootDir, pc.ModulePath).Generate(apiVersionPkgList); err != nil {
		panic(errors.Wrap(err, "cannot generate register file"))
	}

	// Generate the provider,
	// i.e. the setup function and optionally the provider's main program.
	i := 0
	for ke, va := range controllerPkgMap {
		fmt.Println(i, " ", ke)
		j := 0
		for _, v := range va {
			fmt.Println(j, " ", v)
			j += 1
		}

		i += 1
	}
	if err := pipeline.NewProviderGenerator(rootDir, pc.ModulePath).Generate(controllerPkgMap, pc.MainTemplate); err != nil {
		panic(errors.Wrap(err, "cannot generate setup file"))
	}

	// NOTE(muvaf): gosec linter requires that the whole command is hard-coded.
	// So, we set the directory of the command instead of passing in the directory
	// as an argument to "find".
	apisCmd := exec.Command("bash", "-c", "goimports -w $(find . -iname 'zz_*')")
	apisCmd.Dir = filepath.Clean(filepath.Join(rootDir, "apis"))
	if out, err := apisCmd.CombinedOutput(); err != nil {
		panic(errors.Wrap(err, "cannot run goimports for apis folder: "+string(out)))
	}

	internalCmd := exec.Command("bash", "-c", "goimports -w $(find . -iname 'zz_*')")
	internalCmd.Dir = filepath.Clean(filepath.Join(rootDir, "internal"))
	if out, err := internalCmd.CombinedOutput(); err != nil {
		panic(errors.Wrap(err, "cannot run goimports for internal folder: "+string(out)))
	}

	fmt.Printf("\nGenerated %d resources!!!\n", count)
}

func sortedResources(m map[string]*config.Resource) []string {
	result := make([]string, len(m))
	i := 0
	for g := range m {
		result[i] = g
		i++
	}
	sort.Strings(result)
	return result
}
