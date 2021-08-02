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
)

// viewWalletCmd represents the viewWallet command
var viewWalletCmd = &cobra.Command{
	Use:   "viewWallet",
	Short: "pulls up wallet under given name and displays wallet information",
	Long: `adf`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pulling up wallet...")
		walletName := args[0] + "_wallet.gob"
		w1 := node.ReadWallet(walletName)
		w1.Print()


	},
}

func init() {
	rootCmd.AddCommand(viewWalletCmd)

	// Here you will define your flags and configuration settings.
	//walletName := flag.String("name", "", "name associated with wallet")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewWalletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewWalletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
