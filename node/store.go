package node

import (
  "fmt"
  "os"
  "encoding/gob"
  //"github.com/syndtr/goleveldb"
)

//func StoreUTXO()

func SaveWallet(w Wallet) {
  wFile, err := os.Create("localWallet.gob")

  if err != nil {
    fmt.Println("wallet failed to save locally")
    fmt.Println(err)
    os.Exit(1)
  } else {
    fmt.Println("wallet saved")
  }

  enc := gob.NewEncoder(wFile)
  enc.Encode(w)


}

func ReadWallet() Wallet {
  var w Wallet
  wFile, err := os.Open("localWallet.gob")

  if err != nil{
    fmt.Println("failed to open localWallet.gob")
    fmt.Println(err)
    os.Exit(1)
  }

  dec := gob.NewDecoder(wFile)
  err = dec.Decode(&w)

  if err != nil{
    fmt.Println(err)
    os.Exit(1)
  }
  wFile.Close()
  return w;

}

func SaveChain(bc []Block){
  chainFile, err := os.Create("savedChain.gob")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  dataEncoder := gob.NewEncoder(chainFile)
  dataEncoder.Encode(bc)
  fmt.Println("current chain saved under 'savedChain.gob' ")
}

func ReadChain() []Block{
  var bc []Block
  chainFile, err := os.Open("savedChain.gob")

  if err != nil{
    fmt.Println(err)
    os.Exit(1)
  }

  dataDecoder := gob.NewDecoder(chainFile)
  err = dataDecoder.Decode(&bc)

  if err != nil{
    fmt.Println(err)
    os.Exit(1)
  }

  chainFile.Close()
  return bc;
}
