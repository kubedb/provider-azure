/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"fmt"
	"github.com/crossplane/upjet/pkg/pipeline"
	dynamic_controller "kubedb.dev/provider-azure/cmd/dynamic-controller"
	pconfig "kubedb.dev/provider-azure/config"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	pc := pconfig.GetProvider()
	pipeline.Run(pc, absRootDir)
	dynamic_controller.GenerateController(pc, absRootDir)
}
