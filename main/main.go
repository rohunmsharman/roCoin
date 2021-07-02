package main

import(
  //fmt"
  //"crypto/sha256"
  //"encoding/hex"
  //"strconv"
  "roCoin.com/roCoin/node" //the goroot issue was fixed by roCoin.com ....
)

func main() {
  genTxns := []node.Txn{}
  genesis := node.NewBlock(genTxns, 0, "0", "nil")
  genesis.Hash = node.CalculateHash(genesis)

  chain := node.CreateChain(genesis)

  bob := node.NewWallet("bob", 100)
  alice := node.NewWallet("alice", 50)
  tstTxns := []node.Txn{node.CreateTxn(bob, alice, 50)}
  b1 := node.NewBlock(tstTxns, 1, genesis.Hash, "nil")
  b1.Hash = node.CalculateHash(b1)
  chain = append(chain, b1)

  /*
  LAST RUN
  output:
   sender coin amount: 2
   recip coin amount: d

  probably has something to do with pointers
*/

}
