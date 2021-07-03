package node
import(
  "fmt"
  "crypto/sha256"
  "encoding/hex"
  //"encoding/json"
//  "time"
  "strconv"
)

type Block struct{
  Index int //index on the chain
  Txns []Txn //array of txns fixed length for the sake of the merkle tree implementation
  PrevHash string //prev block hash
  Hash string //current block hash
  MerkleRoot string //merkle root of txns
  //Timestamp string //to be updated with timestamping
  Nonce int //nonce for PoW
  Target int //...target
}


func CalculateHash(block Block) string {
  //time := time.Now().String()
  //hash missing time and nonce
                          //replace TxnsToString with merkle MerkleRoot
  h := sha256.Sum256([]byte( block.PrevHash + strconv.Itoa(block.Nonce))) //CalculateMerkleRoot(block.Txns) +
  return (hex.EncodeToString(h[:]));
}

//rewrite NewBlock to take in an array of txns and output block, pull everything else of valid chain state
func NewBlock(index int, txns []Txn, prevHash string, hash string, nonce int, target int) Block{
  return Block{Index: index, Txns: txns, PrevHash: prevHash, Hash: hash, MerkleRoot: "asdf", Nonce: nonce, Target: target};
}

func PrintBlock(block Block){
  fmt.Println("Block Index: ", block.Index)
  fmt.Println("Hash: ", block.Hash)
  fmt.Println("Previous Hash: ", block.PrevHash)
  fmt.Println("Merkle Root: ", block.MerkleRoot)
}
/*
func AddTxn(block *Block, tx Txn){
  block.txns = append(block.txns, tx)
}
*/



//implement merkle tree
/*
func MerkleRoot(block Block){
  //use []txns in block to calculate merkleroot
}
*/
