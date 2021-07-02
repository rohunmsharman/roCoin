package main

import(
  "fmt"
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
  fmt.Println(chain[0])
  bob := node.NewWallet("bob", 100)
  //fmt.Println(bob.Name, " ", bob.Coins)
  alice := node.NewWallet("alice", 50)

  //add txns works
  tstTxns := []node.Txn{node.CreateTxn(bob, alice, 50)}
  /*
  fmt.Println("alice")
  alice.Coins = tstTxns[0].Amount + alice.Coins
  fmt.Println(alice.Coins)

  fmt.Println("bob")
  bob.Coins = bob.Coins - tstTxns[0].Amount
  fmt.Println(bob.Coins)
  */

  b1 := node.NewBlock(tstTxns, 1, genesis.Hash, "nil")
  b1.Hash = node.CalculateHash(b1)
  chain = append(chain, b1)
  fmt.Println(chain[1])
  if chain[0].Hash != chain[1].PrevHash{
    fmt.Println("mismatched hashes")
  }
  //fmt.Println(node.IsChainValid(chain))
  /*
  LAST RUN
  output:
   sender coin amount: 2
   recip coin amount: d

  probably has something to do with pointers
*/

}
