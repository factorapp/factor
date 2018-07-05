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
	"os"
	"path/filepath"

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
	Short: "A brief description of your command",
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
		err = makeDirectories(cwd)
		if err != nil {
			fmt.Println("error making directories:", err)
			return
		}
	},
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
	initCmd.Flags().StringP("name", "n", "", "Name of your new application")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
