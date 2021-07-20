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
  PrivKey ecdsa.PrivateKey
  PubKey ecdsa.PublicKey

}

func AddUTXO(wallet Wallet, txn Txn){
  //add check to make sure UTXO belongs to the wallet
  wallet.PrivKeys[txn.TxnID] = txn.Amount
}

func GetAmount(wallet Wallet) int { //to be deprecated with DB
  sum := 0
  for _, amount := range wallet.PrivKeys {
    sum = sum + amount
  }
  return sum;
}



func SendCoin(send Wallet, recip Wallet, amount int){
  for txnID, amt := range send.PrivKeys{
    if amt == amount{
      CreateTxn(send, recip, amount)
      recip.PrivKeys[txnID] = amt
      delete(send.PrivKeys, txnID)
    }
  }

}


func NewWallet(name string, amount int) Wallet{
  wName := name + "_wallet"
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

  return Wallet{Name: wName, PrivKeys: pK, Coins: amount, PrivKey: *privKey, PubKey: privKey.PublicKey}
}
