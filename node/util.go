package node
/*
All utility functions:
 - persist to disc
 - read from disc
 - json encoding
 - marshalling
 - merkle root
*/
import(
  "fmt"
  "os"
  "bytes"
  "math/big"
  "encoding/gob"
  "encoding/json"
  "encoding/hex"
  "crypto/sha256"
  "crypto/ecdsa"
  "crypto/elliptic"
  "strconv"
)

//used generally
func HashToString(hash []byte) string{
  return (hex.EncodeToString(hash[:]))
}

//used generally
func ByteToString(hash []byte) string {
  return (hex.EncodeToString(hash[:]))
}

func TxnToJson(txn Txn) []byte{
  txn_json, err := json.Marshal(txn)
  if err != nil{
    panic(err)
  }
  return txn_json;
}

//also for hashing purposes
func TxnStrDigest(key ecdsa.PublicKey) PubKStr {
  var out PubKStr
  out.X = key.X.String() //key X value
  out.Y = key.Y.String() //key Y value
  out.XY = append(key.X.Bytes(), key.Y.Bytes()...)
  return out;
}

func TxnHashByte(txn Txn, utxo Txn) []byte { //need to add prev txn
  txStrTemp := TxnStrDigest(txn.RecipPubKey) //creates temp tx of type PubKStr
  t := append(txStrTemp.XY[:], utxo.TxnID...)
  txId := sha256.Sum256(t)
  return txId[:];
}

//encodes txn as gob to be broadcasted
func (tx Txn) MarshalTx() []byte{
  var buf bytes.Buffer

  enc := gob.NewEncoder(&buf)
  err := enc.Encode(tx) //switched from just tx, not sure yet
  if err != nil {
    panic(err)
  }
  fmt.Println("txn marshalled ")
  return buf.Bytes(); //returns slice of unread buffer
}

func UnmarshalTx(inTx []byte) Txn {
  txn := Txn{}
  buf := bytes.NewBuffer(inTx)
  dec := gob.NewDecoder(buf)

  err := dec.Decode(&txn)
  if err != nil {
    panic(err)
    return txn
  }
  return txn;
}

func BlockToJson(block Block) []byte{
  block_Json, err := json.Marshal(block)
  if err != nil {
    panic(err)
  }
  return block_Json;
}

func WalletToJson(w Wallet) []byte {
  w_Json, err := json.Marshal(w)
  if err != nil {
    panic(err)
  }
  return w_Json;
}

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


//used in miner.go
func IntToHex(n int64) []byte{
  return []byte(strconv.FormatInt(n, 16))
}

//def needs to be redone, can only take numbers whose divison by two always results in an even number
func CalculateMerkleRoot(txns []Txn) []byte {
  //this needs to be rewritten without strings, using only byte arrays
  txnHashes := []string{}
  nilTx := NilTxn()
  t := [32]byte{} //just becase sum256 outputs this size
  //converts []Txn into array of sha256 strings
  for k, tx := range txns{ //change to txns from tx
    if k == 0 {
      t = sha256.Sum256(TxnHashByte(tx, nilTx)) //add nilTx
      txnHashes = append(txnHashes, hex.EncodeToString(t[:]))
    }else {
      t = sha256.Sum256(TxnHashByte(tx, txns[k-1]))
      txnHashes = append(txnHashes, hex.EncodeToString(t[:]))
    }
  }
  //CHANGE NEEDED
  if len(txnHashes) % 2 != 0 {
    duplicate := txnHashes[len(txnHashes) - 1]
    txnHashes = append(txnHashes, duplicate) //duplicate last txn if size is odd
    fmt.Println("duplicated last txn")
  }

  for len(txnHashes) > 1{

    j := 0

    for i := 0; i < len(txnHashes); i+=2 {
      //hash of leaf (i and (i + 1)), parent hash
      h := sha256.Sum256(append([]byte(txnHashes[i]), []byte(txnHashes[i + 1])...))
      txnHashes[j] = hex.EncodeToString(h[:])
      j += 1

    }
    //removes leftover hashes
    txnHashes = txnHashes[:(len(txnHashes)/2)] //current fix, honestly idk how to loop this better

  }
  return []byte(txnHashes[0]); //outputs parents hash string as byte array
}

//returns a pubkey that can be encoded with gob
func NilPubKey() ecdsa.PublicKey {
  nilKey := new(ecdsa.PublicKey)
  c := elliptic.P256()
  nilKey.Curve = c
  nilKey.X, nilKey.Y = big.NewInt(10000000000), big.NewInt(10000000000)

  return *nilKey;
}
