package main

import(
  "fmt"
  //"crypto/sha256"
  //"encoding/hex"
  //"strconv"
  "roCoin.com/roCoin/node" //the goroot issue was fixed by roCoin.com ....
)

func main() {
  /*
  genTxns := []node.Txn{}
  genesis := node.NewBlock(0, genTxns, "0", "nil", 0, 3)
  genesis.Hash = node.CalculateHash(genesis)

  chain := node.CreateChain(genesis)
  fmt.Println(chain[0])
  */
  bob := node.NewWallet("bob", 100)
  //fmt.Println(bob.Name, " ", bob.Coins)
  alice := node.NewWallet("alice", 50)


  //add txns works
  tstTxns := []node.Txn{node.CreateTxn(bob, alice, 50)}
  tstTxns = append(tstTxns, node.CreateTxn(bob, alice, 30))
  tstTxns = append(tstTxns, node.CreateTxn(bob, alice, 40))
  tstTxns = append(tstTxns, node.CreateTxn(bob, alice, 20))
  tstTxns = append(tstTxns, node.CreateTxn(bob, alice, 10))
  tstTxns = append(tstTxns, node.CreateTxn(bob, alice, 60))
  tstTxns = append(tstTxns, node.CreateTxn(bob, alice, 15))
  tstTxns = append(tstTxns, node.CreateTxn(bob, alice, 12))
  fmt.Println(node.CalculateMerkleRoot(tstTxns))
  /*
  fmt.Println("alice")
  alice.Coins = tstTxns[0].Amount + alice.Coins
  fmt.Println(alice.Coins)

  fmt.Println("bob")
  bob.Coins = bob.Coins - tstTxns[0].Amount
  fmt.Println(bob.Coins)
  */
  /*
  b1 := node.NewBlock(1, tstTxns, genesis.Hash, "nil", 0, 3)
  b1.Hash = node.CalculateHash(b1)
  chain = append(chain, b1)
  fmt.Println(chain[1])
  if chain[0].Hash != chain[1].PrevHash{
    fmt.Println("mismatched hashes")
  }
  */

  //node.MineBlock(b1)

}
