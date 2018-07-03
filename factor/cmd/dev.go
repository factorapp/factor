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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bketelsen/factor/factor/component"
	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(cwd)
		fmt.Println("dev called")
		err = ensureFactor()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = processComponents(cwd)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func processComponents(base string) error {
	dir := filepath.Join(base, "components")

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c, _ := component.Parse(f, componentName(path))
			fmt.Printf("visited file: %q\n", path)
			gofile, err := os.Create(goFileName(componentName(path)))
			if err != nil {
				return err
			}
			defer gofile.Close()
			c.Transform(gofile)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
	return err
}
func goFileName(comp string) string {
	return "factortmp/components/" + strings.ToLower(comp) + ".go"
}
func componentName(path string) string {
	base := filepath.Base(path)
	return strings.Replace(base, filepath.Ext(path), "", -1)

}
func ensureFactor() error {

	_, err := os.Stat("factortmp")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("factor tmp directory does not exist, creating...")
			return os.MkdirAll("factortmp/components", 0755)
		} else {
			return err
		}
	}
	return nil
}
func init() {
	rootCmd.AddCommand(devCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// devCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// devCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
