package node

import(
  //"fmt"
  "crypto/sha256"
  //"time"
  "encoding/hex"
  "strings"
  "strconv"
)

type Txn struct{
  Sender string
  Recipient string
  Amount int
  TxnID string
}

//converts array for transactions to string
//TO BE DEPRACATED
func TxnsToString(txn []Txn) string{
  txStrTemp := []string{}
  for _, j := range txn{
    str := j.Sender + j.Recipient + strconv.Itoa(j.Amount)
    txStrTemp = append(txStrTemp, str)

  }
  txnsString := strings.Join(txStrTemp," ")
  return txnsString;
}

// fix encoding to string, used as Txn ID, rewrite to take txn in directly
func TxnHash(txn Txn) string{
  txId := sha256.Sum256([]byte(txn.Sender + txn.Recipient + strconv.Itoa(txn.Amount)))
  return (hex.EncodeToString(txId[:])); //check the [:]
}


func CreateTxn(sender Wallet, recipient Wallet, amount int) Txn {
  SendCoin(sender, recipient, amount)
  TXN := Txn{Sender: sender.Name, Recipient: recipient.Name, Amount: amount}
  //TXN.TxnID = CalcTxId(TXN)
  return TXN;
}

//implement merkle tree
/*
func CreateMerkle(txn []Txn) string {
  //creates merkle tree of txns and returns merkleroot
}
*/
