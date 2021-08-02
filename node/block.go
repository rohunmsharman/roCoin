package node
import(
  "fmt"
  "crypto/sha256"
  "encoding/hex"
  //"encoding/json"
  "time"
  "strconv"
)

type Block struct{
  Index int //index on the chain
  Txns []Txn //array of txns fixed length for the sake of the merkle tree implementation
  PrevHash []byte //prev block hash
  Hash []byte //current block hash
  MerkleRoot []byte //merkle root of txns
  Timestamp []byte //use TBD, timestamping server should be implemented
  Nonce int //nonce for PoW
  Target int //...target
}


func CalculateHash(block Block) []byte {
  //time := time.Now().String()
  //hash missing time and nonce
  squish := append(block.MerkleRoot[:], block.PrevHash[:]...)
  squish = append(squish, []byte(strconv.Itoa(block.Nonce))...)
  squish = append(squish, block.Timestamp...)
  squish = append(squish, []byte(strconv.Itoa(block.Target))...)
  h := sha256.Sum256(squish)
  stretch := []byte{}
  stretch = append(stretch[:], h[:]...) //this is because h is type [32]byte
  return stretch;
}

func HashToString(hash []byte) string{
  return (hex.EncodeToString(hash[:]))
}

func ByteToString(hash []byte) string {
  return (hex.EncodeToString(hash[:]))
}

//rewrite NewBlock to take in an array of txns and output block, pull everything else of valid chain state
func NewBlock(prevBlock Block, txns []Txn, hash string, target int) Block{
  newBlock := Block{}
  newBlock.Index = prevBlock.Index + 1
  newBlock.Txns = txns
  newBlock.PrevHash = prevBlock.Hash
  newBlock.MerkleRoot = CalculateMerkleRoot(txns)
  newBlock.Timestamp = []byte(time.Now().String()) //time at production of block (use tbd)
  newBlock.Nonce = 0 //unless we decide to start from elsewhere
  newBlock.Target = target
  newBlock.Hash = CalculateHash(newBlock)
  return newBlock;
}

/*
func GenesisBlock() Block {
  genesis := Block{Index: 0, Txns: }
  return genesis;
}
*/

func PrintBlock(block Block){
  fmt.Println("Block Index: ", block.Index)
  fmt.Println("Hash: ", block.Hash)
  fmt.Println("Previous Hash: ", block.PrevHash)
  fmt.Println("Merkle Root: ", block.MerkleRoot)
}
