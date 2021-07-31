package node
import(
  "fmt"
  "crypto/ecdsa"
  "crypto/rand"
  "crypto/elliptic"
  //"encoding/gob"
  //"math/big"

)

type Wallet struct{
  Name string
  PrivKeys map[string]int //maps txnID and private key !TEMPORARY BC THIS DOESN'T ACTUALLY STORE TXN, DB MUST BE ADDED
  Coins int
  PrivKey *ecdsa.PrivateKey
  PubKey ecdsa.PublicKey
  //UTXOs

}

//txn to be taken in should have an empty sender pubKey field
func (w Wallet) SignTxn(txn Txn) {
  sig, err := ecdsa.SignASN1(rand.Reader, w.PrivKey, TxnHashByte(txn))
  if err != nil {
    fmt.Println("failed to sign txn")
    panic(err)
  }
  txn.Sig = sig
  fmt.Println("txn signed")
}

//verify txn
func VerifyTxn(sendW Wallet, tx Txn) bool {
  valid := ecdsa.VerifyASN1(&sendW.PubKey, tx.TxnID, tx.Sig) //&sendW.Pubkey = *ecdsa.PubKey
  return valid;
}

//prints wallet info
func (w *Wallet) Print(){
  fmt.Println("Name: ", w.Name)
  fmt.Println("Private Keys (UTXOS): ", w.PrivKeys)
  fmt.Println("Coins: ", w.Coins)
  fmt.Println("Private Key: ", w.PrivKey.D.String())
  fmt.Println("Public Key: ", w.PrivKey.PublicKey.X, " ", w.PrivKey.PublicKey.Y)
}

func NewWallet(name string, amount int) Wallet{
  wName := name
  pK := make(map[string]int)
  privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
  if err != nil{
    panic(err)
  }
  fmt.Println("Name: ", wName)
  fmt.Println("PrivKeys: ", pK)
  fmt.Println("Coins: ", amount)
  fmt.Println("PrivKey: ", privKey.D.String()) //something of about formatting but its fine
  fmt.Println("PubKey: ", *privKey.PublicKey.X, " ", *privKey.PublicKey.Y) //something of about formatting but its fine
  //local storage function located in store.go

  return Wallet{Name: wName, PrivKeys: pK, Coins: amount, PrivKey: privKey, PubKey: privKey.PublicKey}
}
