package node

import (
  "fmt"
  "os"
  "encoding/gob"
  "encoding/json"
  //"bytes"
  //"bufio"
  //"github.com/syndtr/goleveldb"
)

//func StoreUTXO()
//save configVars locally
func SaveConfig(config ConfigVars) {
  configFile, err := os.Create("configVars.gob")
  if err != nil {
    fmt.Println("failed to create config file")
    panic(err)
    os.Exit(1) //not sure if this needs to b here
  }

  enc := gob.NewEncoder(configFile)
  enc.Encode(&config)
  fmt.Println("config vars saved")
}

//open and return config vars
func ReadConfig() ConfigVars{
  var configVars ConfigVars
  configFile, err := os.Open("configVars.gob")
  if err != nil {
    panic(err)
    os.Exit(1)
  }

  dec := gob.NewDecoder(configFile)
  err = dec.Decode(&configVars)
  if err != nil {
    panic(err)
    os.Exit(1)
  }

  configFile.Close()
  return configVars
}

func SaveWallet(w Wallet) {
  name := w.Name + "_wallet.gob"
  wFile, err := os.Create(name)
  wBytes := WalletToJson(w) //wrap wallet in json to deal with big int problem
  //idk if the json wrapping is actually necessary, but im leaving it

  if err != nil {
    fmt.Println("wallet failed to save locally")
    panic(err)
    os.Exit(1)
  } else {
    fmt.Println("wallet saved")
    fmt.Println("errors: ")
    fmt.Println(err)
  }

  enc := gob.NewEncoder(wFile)
  enc.Encode(&wBytes) //was originally just w

}

func ReadWallet(wName string) Wallet {
  var w Wallet
  var wBytes []byte
  wFile, err := os.Open(wName)

  if err != nil{
    fmt.Println("failed to open ", wFile)
    panic(err)
    os.Exit(1)
  }

  dec := gob.NewDecoder(wFile)
  err = dec.Decode(&wBytes)
  //err = dec.Decode(&w)
  if err != nil{
    fmt.Println(1)
    //fmt.Println(err)
    panic(err)
    os.Exit(1)
  }
  err = json.Unmarshal(wBytes, &w)
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
