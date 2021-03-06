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
	"roCoin/node"
	"github.com/spf13/cobra"
	//"strconv"
)

// newWalletCmd represents the newWallet command
var newWalletCmd = &cobra.Command{
	Use:   "newWallet", //ADD FLAGS!!
	Short: "generates new wallet",
	Long: `creates a new wallet, prints out wallet info, will save wallet info as gob file `,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("creating new wallet")
		//tInt, err := strconv.Atoi(args[1])

		node.SaveWallet(node.NewWallet(args[0]))

		//fmt.Println(w1)
	},
}

func init() {
	rootCmd.AddCommand(newWalletCmd)

	// Here you will define your flags and configuration settings.
	//ADD FLAGS
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newWalletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newWalletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
