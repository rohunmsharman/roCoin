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

//encode a test txn w/ JSON
/*
func TestTxnToJson(txn TestTxn) []byte {
  tstTxn_Json, err := json.Marshal(txn)
  if err != nil {
    panic(err)
  }

  return tstTxn_Json;
}
*/

func TxnToJson(txn Txn) []byte{
  txn_json, err := json.Marshal(txn)
  if err != nil{
    fmt.Println(err)
  }
  return txn_json;
}

func WalletToJson(wallet Wallet) []byte{
  wallet_json, err := json.Marshal(wallet)
  if err != nil{
    fmt.Println(err)
  }
  return wallet_json
}
