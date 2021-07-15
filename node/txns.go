package node

import(
  "fmt"
  "crypto/sha256"
  //"crypto/ecdsa"
  //"time"
  "encoding/hex"
  "encoding/gob"
  "strconv"
)

type Txn struct{
  SenderPubKey string //sender pubkey used to verify signature
  RecipPubKey string 
  Amount int
  TxnID string
}
/*
IMPLEMENT LATER
type TxnIn struct{

}

type TxnOut struct{

}
*/
//converts a txn to byte array
func TxnToByte(txn Txn) []byte{
    str := txn.Sender + txn.Recipient + strconv.Itoa(txn.Amount) + txn.TxnID
    byteTxn := []byte(str)
    return byteTxn;
}

// fix encoding to string, used as Txn ID, rewrite to take txn in directly
func TxnHash(txn Txn) string{
  txId := sha256.Sum256([]byte(txn.RecipPubKey + strconv.Itoa(txn.Amount)))
  return (hex.EncodeToString(txId[:])); //check the [:]
}

//returns new txn
func CreateTxn(sender Wallet, recipient Wallet, amount int) Txn

  TXN := Txn{SenderPubKey: sender.Name, RecipPubKey: recipient.Name, Amount: amount}
  //TXN.TxnID = CalcTxId(TXN)
  return TXN;
}

//encode pubkey (recipient) to be hashed
func KeyEncoder(wallet Wallet) []byte{
  encoder := gob.NewEncoder()

  wallet.PrivKey
}

//def needs to be redone, can only take numbers whose divison by two always results in an even number
func CalculateMerkleRoot(txn []Txn) string {
  txnHashes := []string{}

  //converts []Txn into array of sha256 strings
  for _, tx := range txn{
    t := sha256.Sum256(TxnToByte(tx))
    txnHashes = append(txnHashes, hex.EncodeToString(t[:]))

  }
  //CHANGE NEEDED
  if len(txnHashes) % 2 != 0 {
    duplicate := txnHashes[len(txnHashes) - 1]
    txnHashes = append(txnHashes, duplicate) //duplicate last txn if size is odd
    fmt.Println("duplicated last txn")
  }

  for len(txnHashes) > 1{

    j := 0

    for i := 0; i < len(txnHashes); i+=2 {
      //hash of leaf (i and (i + 1)), parent hash
      h := sha256.Sum256(append([]byte(txnHashes[i]), []byte(txnHashes[i + 1])...))
      txnHashes[j] = hex.EncodeToString(h[:])
      j += 1

    }
    //removes leftover hashes
    txnHashes = txnHashes[:(len(txnHashes)/2)] //current fix, honestly idk how to loop this better

  }
  return txnHashes[0]; //outputs parents hash
}
