package node
import(
  //"fmt"
  //"crypto/ecdsa"
  //"crypto/rand"
//  "crypto/sha256"

)

type Wallet struct{
  Name string
  //privKeys map[string]string //maps txnID and private key
  Coins int
  //privKey
  //PubKey PublicKey
}

func SendCoin(send Wallet, recip Wallet, amount int){
  send.Coins = send.Coins - amount
  recip.Coins = recip.Coins + amount

}

func NewWallet(name string, value int) Wallet{
  return Wallet{Name: name, Coins: value}
}

//right now this just returns public key
/*
func genKeys() PublicKey{
  privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
  if err != nil{
    panic(err)
  }
  return privKey.PublicKey;
}
*/
