package main

import(
  "fmt"
  //"crypto/sha256"
  //"math/big"
  //"encoding/hex"
  //"strconv"
  //"unsafe"
  "roCoin.com/roCoin/node" //the goroot issue was fixed by roCoin.com ....
)


func main() {


  genTxns := []node.Txn{}
  genesis := node.NewBlock(0, genTxns, "0", "nil", "nil", 0, 3)
  genesis.Hash = node.HashToString(node.CalculateHash(genesis))

  chain := node.CreateChain(genesis)
  fmt.Println(chain[0])

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

  mR := node.CalculateMerkleRoot(tstTxns)
  //fmt.Println(node.CalculateMerkleRoot(tstTxns))

  fmt.Println("alice")
  alice.Coins = tstTxns[0].Amount + alice.Coins
  fmt.Println(alice.Coins)

  fmt.Println("bob")
  bob.Coins = bob.Coins - tstTxns[0].Amount
  fmt.Println(bob.Coins)

  b1 := node.NewBlock(1, tstTxns, genesis.Hash, "nil", mR, 0, 15)
  b1.Hash = node.HashToString(node.CalculateHash(b1))
  fmt.Println(b1.Hash)
  chain = append(chain, b1)
  fmt.Println(chain[1])
  if chain[0].Hash != chain[1].PrevHash{
    fmt.Println("mismatched hashes")
  }
  
  node.MineBlock(b1)
  //fmt.Println("hash: ")
  //fmt.Println(b1.Hash[:b1.Target])
  //node.MineBlock(b1)



/*
  var hash big.Int
  h := sha256.Sum256([]byte("asdfasdfasdfoahweliufasbldiguaelwiugbalweoguas;d"))
  hf := hash.SetBytes(h[:])
  fmt.Println(hf) //outputs number
  //fmt.Println(h)
*/

}
