/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// inspectChainCmd represents the inspectChain command
var inspectChainCmd = &cobra.Command{
	Use:   "inspectChain",
	Short: "displays current max verified chain",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inspectChain called")
	},
}

func init() {
	rootCmd.AddCommand(inspectChainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectChainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectChainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
