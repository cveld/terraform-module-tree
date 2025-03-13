// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	flag "github.com/spf13/pflag"
)

var showJSON = flag.Bool("json", false, "produce JSON-formatted output")
var folders map[string][]string = make(map[string][]string)
var treecache map[string][]string = make(map[string][]string)

func moduleLoad(dir string, root string) {
	_, ok := folders[dir]
	if !ok {
		folders[dir] = []string{}
	}
	folders[dir] = append(folders[dir], root)

	children, ok := treecache[dir]
	//ok = false
	if !ok {
		// cache does not exist, populate:
		treecache[dir] = []string{}
		module, _ := tfconfig.LoadModule(dir)

		if module.Diagnostics.HasErrors() {
			fmt.Printf("Error while loading module: %s\n", module.Diagnostics.Error())
			os.Exit(1)
		}

		for _, s := range module.ModuleCalls {
			if strings.HasPrefix(s.Source, ".") {
				var path string = filepath.Join(dir, s.Source)
				treecache[dir] = append(treecache[dir], path)
			}
		}
		children = treecache[dir]
	}

	// use cache
	for _, path := range children {
		moduleLoad(path, root)
	}

}

func main() {
	flag.Parse()

	var dir string
	if flag.NArg() > 0 {
		dir = flag.Arg(0)
	} else {
		dir = "."
	}

	jsonFile, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//var result map[string]interface{}
	var result []string

	json.Unmarshal(byteValue, &result)

	fmt.Printf("%d\n", len(result))
	for i, s := range result {
		fmt.Println(i, s)
		start := time.Now()
		moduleLoad(s, s)
		elapsed := time.Since(start)
		fmt.Printf("Time elapsed: %s\n", elapsed)
	}

	file, err := json.MarshalIndent(folders, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("output.json", file, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func showModuleJSON(module *tfconfig.Module) {
	j, err := json.MarshalIndent(module, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error producing JSON: %s\n", err)
		os.Exit(2)
	}
	os.Stdout.Write(j)
	os.Stdout.Write([]byte{'\n'})
}

func showModuleMarkdown(module *tfconfig.Module) {
	err := tfconfig.RenderMarkdown(os.Stdout, module)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error rendering template: %s\n", err)
		os.Exit(2)
	}
}
