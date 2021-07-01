package main

import(
  //fmt"
  "crypto/sha256"
  //"encoding/hex"
  "strconv"
  "roCoin/main/node"
)

func main() {
  //UTXOs = make(map[string]string)
  walletA := wallet{name: "bob", coins: 100}
  walletB := wallet{name: "alice", coins: 50}
  chain := []block
  genesisBlock := block{index: 0, prevHash: "0"}
  gHash := sha256.Sum256([]byte(strconv.Itoa(genesisBlock.index) + genesisBlock.prevHash))
  genesisBlock.hash =  gHash
  chain = append(chain, genesisBlock)
  //txnTemp := []txn
  block1 := block{index: 1, prevHash: gHash}
  addTxn(block1, sendCoin(walletA, walletB, 25))
  block1.hash = calculateHash(block1)
  //fmt.println()
}
