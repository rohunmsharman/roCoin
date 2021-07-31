package node

import (
  "fmt"
  "os"
  "encoding/gob"
  //"bufio"
  //"github.com/syndtr/goleveldb"
)

//func StoreUTXO()


func SaveWallet(w Wallet) {
  name := w.Name + "_wallet.gob"
  wFile, err := os.Create(name)

  if err != nil {
    fmt.Println("wallet failed to save locally")
    panic(err)
    os.Exit(1)
  } else {
    fmt.Println("wallet saved")
    fmt.Println(err)
  }

  enc := gob.NewEncoder(wFile)
  enc.Encode(&w) //was originally just w

}

//!! Returns EOF error when trying to read wallet from gob file
func ReadWallet(wName string) Wallet {
  var w Wallet
  wFile, err := os.Open(wName)

  if err != nil{
    fmt.Println("failed to open ", wFile)
    panic(err)
    os.Exit(1)
  }

  dec := gob.NewDecoder(wFile)
  err = dec.Decode(&w)

  if err != nil{
    fmt.Println(1)
    //fmt.Println(err)
    panic(err)
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
