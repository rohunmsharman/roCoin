package node

import(
  "encoding/json"
  "fmt"
)
/*
type Block_JSON struct{
  Index int `json:"block_index"`
  PrevHash string `json:"block_PrevHash"`
  Hash string `json:"block_Hash"`
  MerkleRoot string `json:"block_Hash"`
}
*/

func BlockToJson(block Block) []byte{
  block_Json, err := json.Marshal(block)
  if err != nil {
    fmt.Println(err)
  }
  return block_Json;
}

func TxnToJson(txn TXN) []byte{
  txn_json, err := json.Marshal(txn)
  if err != nil{
    fmt.Println(err)
  }
  return txn_json;
}
