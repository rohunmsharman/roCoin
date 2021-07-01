package node

import(
  "fmt"
  "crypto/sha256"
  "encoding/base64"
  //"time"
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
func TxnsToString(txn []Txn) string{
  txStrTemp := []string{}
  for _, j := range txn{
    str := j.Sender + j.Recipient + strconv.Itoa(j.Amount)
    txStrTemp = append(txStrTemp, str)

  }
  txnsString := strings.Join(txStrTemp," ")
  return txnsString;
}

// fix encoding to string, used as Txn ID 
func CalcTxId(txn Txn) string{
  txId := base64.StdEncoding.EncodeToString(sha256.Sum256([]byte(txn.Sender + txn.Recipient + strconv.Itoa(txn.Amount))))
  return fmt.Sprintf("%x", txId);
}
//implement merkle tree
func CreateTxn(sender string, recipient string, amount int) Txn {
  TXN := Txn{Sender: sender, Recipient: recipient, Amount: amount}
  TXN.TxnID = CalcTxId(TXN)
  return TXN;
}
