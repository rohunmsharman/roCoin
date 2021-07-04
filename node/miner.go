package node

import(
  "fmt"
  "strconv"
)

func MineBlock(block Block) bool{
  //implementation completely wrong: USE: https://justinmeiners.github.io/tiny-blockchain/
  t := make([]byte, block.Target)
  for i := 0; i < len(t); i++{
    t[i] = 0
  }

  fmt.Println("mining")
  for block.Hash[:block.Target] != string(t){
    block.Nonce++
    block.Hash = CalculateHash(block)
    fmt.Println("nonce: " + strconv.Itoa(block.Nonce) + " Hash: " + block.Hash)

  }
  return true; //should return true once the hash meets min difficulty
}
