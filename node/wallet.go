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
  //PrivKeys map[string]int //maps txnID and private key !TEMPORARY BC THIS DOESN'T ACTUALLY STORE TXN, DB MUST BE ADDED
  Coins int
  PrivKey *ecdsa.PrivateKey
  PubKey ecdsa.PublicKey
  UTXOs []Txn
  //UTXOs

}

//txn to be taken in should have an empty sender pubKey field
func (w Wallet) SignTxn(txn Txn, prevTxn Txn) {
  sig, err := ecdsa.SignASN1(rand.Reader, w.PrivKey, TxnHashByte(txn, prevTxn))
  if err != nil {
    fmt.Println("failed to sign txn")
    panic(err)
  }
  txn.Sig = sig
  fmt.Println("txn signed")
}



//prints wallet info
func (w *Wallet) Print(){
  fmt.Println("Name: ", w.Name)
  fmt.Println("Private Keys (UTXOS): ", len(w.UTXOs))
  fmt.Println("Coins: ", w.Coins)
  fmt.Println("Private Key: ", w.PrivKey.D.String())
  fmt.Println("Public Key: ", w.PrivKey.PublicKey.X, " ", w.PrivKey.PublicKey.Y)
}

func NewWallet(name string) Wallet{
  wName := name
  //var utxo []Txn
  //append(utxo, Txn{})
  pK := make(map[string]int)
  pK["holder"] = 0 //to allow gob to encode it
  //fmt.Println("privKey map: ", pK)
  privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
  if err != nil{
    panic(err)
  }
  fmt.Println("Name: ", wName)
  fmt.Println("PrivKeys: ", pK)
  //fmt.Println("UTXOs: ", len(UTXOs))
  //fmt.Println("Coins: ", amount)
  fmt.Println("PrivKey: ", privKey.D.String()) //something of about formatting but its fine
  fmt.Println("PubKey: ", *privKey.PublicKey.X, " ", *privKey.PublicKey.Y) //something of about formatting but its fine
  //local storage function located in store.go

  //coin amount should be determined by UTXOs
  return Wallet{Name: wName, Coins: 0, PrivKey: privKey, PubKey: privKey.PublicKey}
}
