package node

import(
  "fmt"
  "crypto/sha256"
  "crypto/ecdsa"
  "crypto/rand"
  "math/big"
  //"io/ioutil"
  "encoding/hex"
  //"encoding/gob"
  "strconv"
)

type Txn struct{
  RecipPubKey PublicKey
  SenderPubKey PublicKey
  SenderSig []byte //sender pubkey used to verify signature
  Amount int
  TxnID string
}

//for hashing purposes, will depracate when networking begins
type PublicKey struct{
  X big.Int
  Y big.Int
  XY []byte
}

// fix encoding to string, used as Txn ID, rewrite to take txn in directly
func TxnHash(txn Txn) string{
  txId := sha256.Sum256(txn.RecipPubKey.XY[:])
  return (hex.EncodeToString(txId[:])); //check the [:]
}

//returns new txn
func CreateTxn(sender Wallet, recipient Wallet, amount int) Txn{

  recipKey := PublicKey{X: *recipient.PubKey.X, Y: *recipient.PubKey.Y}
  recipKey.XY = append(recipKey.X.Bytes(), recipKey.Y.Bytes()...)
  sendKey := PublicKey{X: *sender.PubKey.X, Y: *sender.PubKey.Y}
  sendKey.XY = append(sendKey.X.Bytes(), recipKey.Y.Bytes()...)
  TXN := Txn{RecipPubKey: recipKey, SenderPubKey: sendKey, Amount: amount}
  TXN.TxnID = TxnHash(TXN)
  TXN.SenderSig = SignTxn(sender, TXN)
  return TXN;
}

func SignTxn(wallet Wallet, txn Txn) []byte{
  sig, err := ecdsa.SignASN1(rand.Reader, &wallet.PrivKey, []byte(txn.TxnID))
  if err != nil{
    panic(err)
  }
  return sig;
}

func VerifySig(wallet Wallet, txn Txn) bool {
  fmt.Println("signature: %x\n", txn.SenderSig) //formatting for signature
  return ecdsa.VerifyASN1(&wallet.PrivKey.PublicKey, []byte(txn.TxnID), txn.SenderSig[:]);
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

//converts a txn to byte array
func TxnToByte(txn Txn) []byte{
    str := strconv.Itoa(txn.Amount) + txn.TxnID //REWRITE
    byteTxn := []byte(str)
    return byteTxn;
}
