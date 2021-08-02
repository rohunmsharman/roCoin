package main

import(
  //"fmt"
  //"encoding/json"
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

  //ADD goroutines to keep receiving txns and handling them

  cmd.Execute() //RUN EVERYTHING THROUGH CLI NOW



  //json marhsalling test?
  /*
  bob := node.NewWallet("bob", 500)
  alice := node.NewWallet("alice", 500)

  tx1 := node.CreateTxn(bob, alice, 200)
  fmt.Println(tx1)
  tx_Json := node.TxnToJson(tx1)
  fmt.Println(tx_Json)
  var txUnmarshal node.Txn
  err := json.Unmarshal(tx_Json, &txUnmarshal)
  if err != nil {
    panic(err)
  }
  */



  //figuring out why gob decoder fails to read wallet
  /*
  noKey := node.Wallet{Name: "bob", Coins: 300}
  withKey := node.NewWallet("alice", 600)

  fmt.Println("")

  node.SaveWallet(noKey)
  node.SaveWallet(withKey)

  fmt.Println("no key: ")
  nK := node.ReadWallet("bob_wallet.gob")
  fmt.Println(nK)

  fmt.Println("with key: ")
  wK := node.ReadWallet("alice_wallet.gob")
  fmt.Println(wK)
  */

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
