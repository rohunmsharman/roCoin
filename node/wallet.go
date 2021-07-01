package node
import(
  "fmt"
)

type Wallet struct{
  Name string
  //privKeys map[string]string //maps txnID and private key
  Coins int
}

func SendCoin(send Wallet, recip Wallet, amount int) Txn{
  send.Coins = send.Coins - amount
  recip.Coins = recip.Coins + amount
  fmt.Println("sender coin amount: " + string(send.Coins))
  fmt.Println("recip coin amount: " + string(recip.Coins))
  return CreateTxn(send.Name, recip.Name, amount)
}
