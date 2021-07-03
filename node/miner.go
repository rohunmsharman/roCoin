package node

import(
  "fmt"
)

func MineBlock(block Block) bool{

  t := make([]byte, block.Target)
  for i := 0; i < len(t); i++{
    t[i] = 0
  }


  for block.Hash[:block.Target] != string(t){
    block.Nonce++
    block.Hash = CalculateHash(block)
    fmt.Println("nonce: ",  block.Nonce, "Hash",  ) //syntax wrong

  }
  return true; //should return true once the hash meets min difficulty
}
