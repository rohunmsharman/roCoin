package main

import(
  "fmt"
  //"crypto/ecdsa"
  //"math/big"
  //"encoding/hex"
  //"strconv"
  //"unsafe"
  "roCoin/node"
  "roCoin/cmd"
  //"roCoin.com/roCoin/node" //the goroot issue was fixed by roCoin.com ....
)


func main() {

  //txn marshalling test
  bob := node.NewWallet("bob", 500)
  alice := node.NewWallet("alice", 600)

  tx1 := node.CreateTxn(bob, alice, 300)
  bob.SignTxn(tx1)

  fmt.Println("marshalling txn")
  marshalTx1 := tx1.MarshalTx()
  unmarshalTx1 := node.UnmarshalTx(marshalTx1)

  fmt.Println(unmarshalTx1)

  cmd.Execute() //RUN EVERYTHING THROUGH CLI NOW

/*  genTxn := node.Txn{RecipPubKey: nil, SenderPubKey: nil, }
  w1 := node.NewWallet("bob")
  w2 := node.NewWallet("alice")
  //fmt.Println(w1.PrivKey.PublicKey)

  tx1 := node.CreateTxn(w1, w2, 50)
  //fmt.Println(tx1.SenderPubKey.XY)

  fmt.Println(node.VerifySig(w1, tx1))
  */
  //fmt.Println(tx1)

  //fmt.Println(node.CreateTxn(w1, w2, 50))


  //GEN TEST
  /*
  genTxns := []node.Txn{}
  genesis := node.NewBlock(0, genTxns, "0", "nil", "nil", 0, 3)
  genesis.Hash = node.HashToString(node.CalculateHash(genesis))

  chain := node.CreateChain(genesis)
  fmt.Println(chain[0])

  bob := node.NewWallet("bob", 100)
  alice := node.NewWallet("alice", 50)


  //add txns
  tstTxns := []node.Txn{}
  tx1 := node.CreateTxn(bob, alice, 20)


  mR := node.CalculateMerkleRoot(tstTxns)
  //fmt.Println(node.CalculateMerkleRoot(tstTxns))


  alice.Coins = tstTxns[0].Amount + alice.Coins

  bob.Coins = bob.Coins - tstTxns[0].Amount


  b1 := node.NewBlock(1, tstTxns, genesis.Hash, "nil", mR, 0, 15)
  b1.Hash = node.HashToString(node.CalculateHash(b1))
  fmt.Println(b1.Hash)
  chain = append(chain, b1)
  fmt.Println(chain[1])
  if chain[0].Hash != chain[1].PrevHash{
    fmt.Println("mismatched hashes")
  }

  node.MineBlock(b1)

  node.SaveChain(chain)

  fmt.Println(len(node.ReadChain()))
  fmt.Println("")

  fmt.Println(string(node.BlockToJson(chain[1])))
  */
  //fmt.Println("hash: ")
  //fmt.Println(b1.Hash[:b1.Target])
  //node.MineBlock(b1)



}
