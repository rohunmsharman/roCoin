package node
import(
//  "fmt"
  "crypto/sha256"
  "encoding/hex"
//  "time"
//  "strconv"
)

type Block struct{
  //MerkleRoots string //merkle root of txns
  txns []Txn //array of txns
  Index int //index on the chain
  PrevHash string //prev block hash
  Hash string //current block hash
  //Timestamp string //to be updated with timestamping
  //Nonce int //nonce for PoW
  //Target int //...target
}


func CalculateHash(block Block) string {
  //time := time.Now().String()
  //hash missing time and nonce
                          //replace TxnsToString with merkle MerkleRoot
  h := sha256.Sum256([]byte(TxnsToString(block.txns) + block.PrevHash)) //strconv.Itoa(block.Nonce)
  return (hex.EncodeToString(h[:]));
}

//rewrite NewBlock to take in an array of txns and output block, pull everything else of valid chain state
func NewBlock(txns []Txn, index int, prevHash string, hash string) Block{
  return Block{txns: txns, Index: index, PrevHash: prevHash, Hash: hash};
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
