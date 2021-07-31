package node

import(
  "fmt"
  "crypto/sha256"
  //"crypto/ecdsa"
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
  RecipPubKey PublicKey //!!should be stored as actual ecdsa PubKey
  SenderPubKey PublicKey //!!should be stored as actual ecdsa PubKey
  Sig []byte
  Amount int
  TxnID []byte
}

//for hashing purposes, will depracate when networking begins
type PublicKey struct {
  X  string //big.Int changed to string to support gob encoding
  Y  string //big.Int changed to string to support gob encoding
  XY []byte
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


//needs to be depracated
func TxnHash(txn Txn) string {
  txId := sha256.Sum256(txn.RecipPubKey.XY[:]) //!! need to add prev txn
  return (hex.EncodeToString(txId[:]));
}

func TxnHashByte(txn Txn) []byte {
  txId := sha256.Sum256(txn.RecipPubKey.XY[:])
  return txId[:];
}

//returns new txn
func CreateTxn(sender Wallet, recipient Wallet, amount int) Txn{
  rX := recipient.PubKey.X.String() //recipient pubKey X
  rY := recipient.PubKey.Y.String() //recipient pubKey Y
  sX := sender.PubKey.X.String() //sender pubKey X
  sY := sender.PubKey.Y.String() //sender pubKey Y

  recipKey := PublicKey{X: rX, Y: rY}
  recipKey.XY = append(recipient.PubKey.X.Bytes(), recipient.PubKey.Y.Bytes()...)

  sendKey := PublicKey{X: sX, Y: sY}
  sendKey.XY = append(sender.PubKey.X.Bytes(), sender.PubKey.Y.Bytes()...)

  TXN := Txn{RecipPubKey: recipKey, SenderPubKey: sendKey, Amount: amount}
  TXN.TxnID = TxnHashByte(TXN)
  //TXN.SenderSig = SignTxn(sender, TXN)
  return TXN;
}

/* SignTxn(wallet Wallet, txn Txn) []byte{
  sig, err := ecdsa.SignASN1(rand.Reader, &wallet.PrivKey, []byte(txn.TxnID))
  if err != nil{
    panic(err)
  }
  return sig;
}
*/

/*
func VerifySig(wallet Wallet, txn Txn) bool {
  fmt.Println("signature: %x\n", txn.SenderSig) //formatting for signature
  return ecdsa.VerifyASN1(&wallet.PrivKey.PublicKey, []byte(txn.TxnID), txn.SenderSig[:]);
}
*/


//def needs to be redone, can only take numbers whose divison by two always results in an even number
func CalculateMerkleRoot(txns []Txn) string {
  txnHashes := []string{}

  //converts []Txn into array of sha256 strings
  for _, tx := range txns{ //change to txns from tx
    t := sha256.Sum256(TxnHashByte(tx))
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

//converts a txn to byte array
//!! TO BE DEPRACATED
/*
func TxnToByte(txn Txn) []byte{
    str := strconv.Itoa(txn.Amount) + txn.TxnID //REWRITE
    byteTxn := []byte(str)
    return byteTxn;

}
*/
