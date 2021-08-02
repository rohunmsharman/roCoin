package node
//meant for chechking chain validity and functions related to the chain
import(
  "fmt"
  "math/big"
)
func CreateChain(genBlock Block) []Block{
  return []Block{genBlock};
}


func IsChainValid(bc []Block) bool {
  for i := 0; i < len(bc); i++ { //deal with error of genesis block

    //check if previous block hash is equals the prev hash of current block
    if HashToString(bc[i].PrevHash) != HashToString(bc[i - 1].Hash){
      fmt.Println("previous hashes not equal")
      return false;
    }
    if HashToString(bc[i].Hash) != HashToString(CalculateHash(bc[i])){
      fmt.Println("current hashes not equal")
      return false;
    }

    //check if current block has been hashed
    if !PoWCheck(bc[i]) {
      fmt.Println("block: ", bc[i].Index, " hash: ", HashToString(bc[i].Hash), " is not mined")
      return false
    }

    //MISSING TXN VERIFICATION
  }

  fmt.Println("blockchain is valid")
  return true;
  }


func TxnsValid(txns []Txn) bool {
  for i := 1; i < len(txns); i++ {
    //checking signatures
    prevPubKey := &txns[i-1].RecipPubKey //pointer for verify function
    if !VerifyTxn(*prevPubKey, txns[i]){
      //sigature does not verify
      fmt.Println("txn signature invalid: ", ByteToString(txns[i].TxnID))
      return false
    }
    //check txns based on consensus (i.e. inputs and outputs)
  }
  return true;
}
func PoWCheck(b Block) bool {
  var hashInt big.Int
  hash := CalculateHash(b)
  hf := hashInt.SetBytes(hash[:])
  targetInt := big.NewInt(1)
  targetInt.Lsh(targetInt, uint(256 - b.Target))

  if hf.Cmp(targetInt) != -1 {
    //fmt.Println("block is not mined")
    //block hash does not meet recalculated hash w/ current nonce
    return false
  }
  //block hash meets target
  return true
}
