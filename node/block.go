package node
import(
  "fmt"
  "crypto/sha256"
  "encoding/base64"
  "time"
  "strconv"
)

type Block struct{
  MerkleRoot string
  txns []Txn
  Index int
  PrevHash string
  Hash string
  Time string
  Nonce int
  Target int
}


func CalculateHash(block Block) string {
  time := time.Now().String()
  //hash missing time and nonce
                          //replace TxnsToString with merkle MerkleRoot
  h := sha256.Sum256([]byte(TxnsToString(block.txns) + block.PrevHash + time + strconv.Itoa(block.Nonce)))
  hash := base64.StdEncoding.EncodeToString(h)
  return hash;
}

func NewBlock(txns []Txn, index int, prevHash string, hash string) Block{
  b := Block{txns: txns, Index: index, PrevHash: prevHash, Hash: hash}
  return b;
}

func AddTxn(block *Block, tx Txn){
  block.txns = append(block.txns, tx)
}

//implement merkle tree
/*
func MerkleRoot(block Block){
  //use []txns in block to calculate merkleroot
}
*/
