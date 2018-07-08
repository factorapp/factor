// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/factorapp/factor/codegen"
	"github.com/gobuffalo/envy"
	"github.com/spf13/cobra"
)

var appName string
var directories = []string{
	"app",
	"assets",
	"client",
	"components",
	"models",
	"routes",
	"server",
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize a new factor application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		appName = args[0]
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		// make directories under appName
		err = makeDirectories(cwd)
		if err != nil {
			fmt.Println("error making directories:", err)
			return
		}
		appPkg := filepath.Join(envy.CurrentPackage(), appName)
		if appPkg == "" {
			fmt.Println("couldn't get the current package for the app")
			return
		}
		fmt.Println("app package", appPkg)
		// put the new files there
		populateApp(cwd, appPkg)
	},
}

// appPath is the location of the app under the GOPATH
func populateApp(cwd, appPkg string) error {
	filename := "index.html"
	filePath := filepath.Join(cwd, appName, "app", filename)
	err := writeTemplate(filePath, indexTemplate)
	if err != nil {
		return err
	}

	filename = "wasm_exec.js"
	filePath = filepath.Join(cwd, appName, "app", filename)
	err = writeTemplate(filePath, codegen.WasmJS)
	if err != nil {
		return err
	}
	filename = "global.css"
	filePath = filepath.Join(cwd, appName, "assets", filename)
	err = writeTemplate(filePath, codegen.GlobalCSS)
	if err != nil {
		return err
	}
	filename = "main.go"
	filePath = filepath.Join(cwd, appName, "client", filename)
	clientGoMain, err := codegen.ClientGoMain(appPkg)
	if err != nil {
		return err
	}
	err = writeTemplate(filePath, clientGoMain)
	if err != nil {
		return err
	}
	filename = "Nav.html"
	filePath = filepath.Join(cwd, appName, "components", filename)
	err = writeTemplate(filePath, codegen.NavComponentHTML)
	if err != nil {
		return err
	}
	filename = "nav.go"
	filePath = filepath.Join(cwd, appName, "components", filename)
	err = writeTemplate(filePath, codegen.NavComponentGo)
	if err != nil {
		return err
	}
	filename = "Index.html"
	filePath = filepath.Join(cwd, appName, "routes", filename)
	err = writeTemplate(filePath, codegen.RoutesHTML)
	if err != nil {
		return err
	}
	filename = "index.go"
	filePath = filepath.Join(cwd, appName, "routes", filename)
	err = writeTemplate(filePath, codegen.RoutesGo)
	if err != nil {
		return err
	}
	filename = "main.go"
	filePath = filepath.Join(cwd, appName, "server", filename)
	serverGoMain, err := codegen.ServerGoMain(appPkg)
	if err != nil {
		return err
	}
	err = writeTemplate(filePath, serverGoMain)
	if err != nil {
		return err
	}
	filename = "Makefile"
	filePath = filepath.Join(cwd, appName, filename)
	err = writeTemplate(filePath, codegen.Makefile)
	if err != nil {
		return err
	}
	return err
}
func writeTemplate(filePath string, templateName string) error {
	gofile, err := os.Create(filePath)
	if err != nil {
		log.Println("ERROR", err)
		return err
	}
	defer gofile.Close()
	tpl := template.Must(template.New("component").Parse(templateName))
	// TODO - get the gopath of cwd and pass it in a context below to this call
	err = tpl.Execute(gofile, nil)
	if err != nil {
		return err
	}
	return nil
}
func makeDirectories(cwd string) error {
	for _, dir := range directories {
		err := os.MkdirAll(filepath.Join(cwd, appName, dir), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.Flags().StringP("name", "n", "", "Name of your new application")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
