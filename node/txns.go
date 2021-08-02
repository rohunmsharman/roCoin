package node

import(
  "fmt"
  "crypto/sha256"
  "crypto/ecdsa"
  //"crypto/rand"
  //"math/big"
  //"io"
  "encoding/hex"
  "encoding/gob"
  //"bufio"
  "bytes"
  //"strconv"
)

type Txn struct{
  RecipPubKey ecdsa.PublicKey //!!should be stored as actual ecdsa PubKey
  //SenderPubKey ecdsa.PublicKey //!!should be stored as actual ecdsa PubKey

  Sig []byte
  //Amount int
  TxnID []byte
}

//for hashing purposes, alterantively could use byte arrays
type PubKStr struct {
  X  string //big.Int changed to string to support gob encoding
  Y  string //big.Int changed to string to support gob encoding
  XY []byte
}

//also for hashing purposes
func TxnStrDigest(key ecdsa.PublicKey) PubKStr {
  var out PubKStr
  out.X = key.X.String() //key X value
  out.Y = key.Y.String() //key Y value
  out.XY = append(key.X.Bytes(), key.Y.Bytes()...)
  return out;
}
//encodes txn as gob to be broadcasted
func (tx Txn) MarshalTx() []byte{
  var buf bytes.Buffer

  enc := gob.NewEncoder(&buf)
  err := enc.Encode(tx) //switched from just tx, not sure yet
  if err != nil {
    panic(err)
  }
  fmt.Println("txn marshalled ")
  return buf.Bytes(); //returns slice of unread buffer
}

func UnmarshalTx(inTx []byte) Txn {
  txn := Txn{}
  buf := bytes.NewBuffer(inTx)
  dec := gob.NewDecoder(buf)

  err := dec.Decode(&txn)
  if err != nil {
    panic(err)
    return txn
  }
  return txn;
}


func TxnHashByte(txn Txn, utxo Txn) []byte { //need to add prev txn
  txStrTemp := TxnStrDigest(txn.RecipPubKey) //creates temp tx of type PubKStr
  t := append(txStrTemp.XY[:], utxo.TxnID...)
  txId := sha256.Sum256(t)
  return txId[:];
}

//!!Change in arguments needed to support cli and sending
func (w Wallet) NewTxn(recipient ecdsa.PublicKey, prevTxn Txn) Txn {
  tx := Txn{RecipPubKey: recipient} //initialize tx

  txId := TxnHashByte(tx, prevTxn) //get tx hash
  tx.TxnID = txId //add hash to tx

  w.SignTxn(tx, prevTxn) //sender wallet signs the txn

  return tx;
}

//verify txn
func VerifyTxn(prevTxPubKey ecdsa.PublicKey, tx Txn) bool { //prevTxPubKey = sender pubkey
  valid := ecdsa.VerifyASN1(&prevTxPubKey, tx.TxnID, tx.Sig) //&sendW.Pubkey = *ecdsa.PubKey
  return valid;
}

//def needs to be redone, can only take numbers whose divison by two always results in an even number
func CalculateMerkleRoot(txns []Txn) []byte {
  //this needs to be rewritten without strings, using only byte arrays
  txnHashes := []string{}

  //converts []Txn into array of sha256 strings
  for _, tx := range txns{ //change to txns from tx
    t := sha256.Sum256(TxnHashByte(tx)) //NEED TO HAVE PREV TXN PASSSED IN
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
  return []byte(txnHashes[0]); //outputs parents hash string as byte array
}

/*
func NilTxn() Txn {

  //returns an empty txn, no field is nil to accomodate encoding, but all fields are meaningless
}
*/
