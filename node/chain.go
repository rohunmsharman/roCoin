package node
//meant for chechking chain validity and functions related to the chain
import(
  "fmt"
  "os"
  "encoding/gob"
)
func CreateChain(genBlock Block) []Block{
  return []Block{genBlock};
}


func IsChainValid(bc []Block) bool {
  for i := 0; i < len(bc); i++ { //deal with error of genesis block
    //check if previous block hash is equals the prev hash of current block
    if bc[i].PrevHash != bc[i - 1].Hash{
      fmt.Println("previous hashes not equal")
      return false;
    }
    if bc[i].Hash != HashToString(CalculateHash(bc[i])){
      fmt.Println("current hashes not equal")
      return false;
    }

    //check if current block has been hashed
    /*
    if bc[i].Hash.substringRelatingToDifficulty != bc[i].Target{
      fmt.Println("block:", bc[i].Index, " has not been mined yet")
      return false;
    }
    */

    //MISSING TXN VERIFICATION
  }

  fmt.Println("blockchain is valid")
  return true;
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
