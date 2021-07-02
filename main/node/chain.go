package node
//meant for chechking chain validity and functions related to the chain
import(
  "fmt"
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
    if bc[i].Hash != CalculateHash(bc[i]){
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
