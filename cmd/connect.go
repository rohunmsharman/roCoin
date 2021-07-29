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
	"context"
	"roCoin/node"
	"roCoin/networking"
	//"strconv"
	"time"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connect = &cobra.Command{
	Use:   "connect",
	Short: "connects node to p2p network",
	Long: `asdf`,
	Args: cobra.MinimumNArgs(1), //port number and txnStream name
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("setting up connection")
		//first argument should just be the string
		//port, err := strconv.Atoi(args[0])
		//if err != nil {
			//panic(err)
		//}
		//take in name of txn stream as argument
		txnStrm := args[0]


		ctx := context.Background()
		// h is host, ps is pubsub from gossibRouter, err is error
		h, ps, err := networking.Setup(ctx)
		if err != nil {
			panic(err)
		}

		//setup local mDNS discovery
		err = networking.SetupDiscovery(ctx, h)
		if err != nil {
			panic(err)
		}

		ts, err := networking.JoinTxnStream(ctx, ps, h.ID(), txnStrm)
		if err != nil {
			panic(err)
		}




		//testing
		//the connect works, broadcasting the TXNs works, just need to create better handling
		var in int
		fmt.Println("peer 1 or 2?")
		fmt.Scanln(&in)

		tstTxns := []node.TestTxn{}
		if in == 1 {
			tstTxns = append(tstTxns, node.TestTxn{Sender: "bob", Recipient: "alice", Amount: 30})
			tstTxns = append(tstTxns, node.TestTxn{Sender: "alice", Recipient: "bob", Amount: 40})
			tstTxns = append(tstTxns, node.TestTxn{Sender: "bob", Recipient: "alice", Amount: 15})
		}

		if in == 2 {
			tstTxns = append(tstTxns, node.TestTxn{Sender: "michelle", Recipient: "britney", Amount: 10})
			tstTxns = append(tstTxns, node.TestTxn{Sender: "britney", Recipient: "michelle", Amount: 20})
			tstTxns = append(tstTxns, node.TestTxn{Sender: "michelle", Recipient: "britney", Amount: 90})
		}

		peers := ts.ListPeers()
		fmt.Println("peers: ")

		for len(ts.ListPeers()) < 1 {
			time.Sleep(3 * time.Second)
			peers = ts.ListPeers()
			for _, peer := range peers {
				fmt.Println(peer)
				fmt.Println("")
			}
			fmt.Println("peers: ")
		}

		go ts.HandleEvents()

		for _, txn := range tstTxns {
			ts.Publish(txn)
		}
	},

}

func init() {
	rootCmd.AddCommand(connect)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
