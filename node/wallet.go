package node
import(
  //"fmt"
  //"crypto/ecdsa"
  //"crypto/rand"
//  "crypto/elliptic"
  //"crypto"
  //"math/big"

)

type Wallet struct{
  Name string
  //privKeys map[string]string //maps txnID and private key
  Coins int
  //PrivKey *ecdsa.PrivateKey //i don't think this workss
  //PubKey ecdsa.PublicKey


}
//signing doesn't work yet
/*
func GenKeyPair(wallet Wallet) {
  var err error
  privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) //pubKey is a part of privKey
  wallet.PrivKey = new(ecdsa.PrivateKey)
  wallet.PrivKey = privKey

  fmt.Println("priavte key: ", wallet.PrivKey.D)

  //wallet.PubKey = wallet.PrivKey.PublicKey


  //fmt.Println("public key: ", wallet.PubKey.X, " ", wallet.PubKey.Y)


  if err != nil{
    panic(err)
  }
}
*/

/*
func SendCoin(send Wallet, recip Wallet, amount int){
  send.Coins = send.Coins - amount
  recip.Coins = recip.Coins + amount

}
*/

func NewWallet(name string, value int) Wallet{
  return Wallet{Name: name, Coins: value}
}

/*
func PrintWallet(wallet Wallet) {
  fmt.Println("Name: ", wallet.Name)
  fmt.Println("Coins: ", wallet.Coins)
  fmt.Println("Private Key: ", wallet.PrivKeyStr)
  fmt.Println("Public Key(s): X: ", wallet.PubKeyStr)
}
*/

//right now this just returns public key
