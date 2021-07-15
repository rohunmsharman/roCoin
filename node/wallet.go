package node
import(
  //"fmt"
  "crypto/ecdsa"
  "crypto/rand"
  "crypto/elliptic"
  //"crypto"
  //"math/big"

)

type Wallet struct{
  Name string
  //privKeys map[string]string //maps txnID and private key
  Coins int
  PrivKey ecdsa.PrivateKey //i don't think this workss
  PubKey ecdsa.PublicKey


}



/*
func SendCoin(send Wallet, recip Wallet, amount int){
  send.Coins = send.Coins - amount
  recip.Coins = recip.Coins + amount

}
*/

func NewWallet(name string, value int) Wallet{
  privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
  if err != nil{
    panic(err)
  }
  return Wallet{Name: name, Coins: value, PrivKey: *privKey, PubKey: privKey.PublicKey}
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
