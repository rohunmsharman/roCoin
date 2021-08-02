package node

import(
  "fmt"
  "strconv"
  "math/big"
)

func MineBlock(block Block) bool{
  //big.Int.SetBytes returns hash and target (byte arrays) as the bytes of a big-endian unsigned integer
  var hashInt big.Int


  hash := CalculateHash(block)
  hf := hashInt.SetBytes(hash[:]) //hf is used bc SetBytes returns size but sets hashInt to byte array
  targetInt := big.NewInt(1)
  targetInt.Lsh(targetInt, uint(256 - block.Target)) //shift the big.Int left by 256 - Target


  fmt.Println("mining: ")

  for hf.Cmp(targetInt) != -1{

    block.Nonce++
    block.Hash = CalculateHash(block)
    hash = CalculateHash(block)
    hf = hashInt.SetBytes(hash[:])

    /*
    fmt.Println("hash size: ", hf)
    fmt.Println()
    fmt.Println("target size: ", targetInt)
    */
  }
  fmt.Println("block mined! ")
  fmt.Println("hash size: ", hf)
  fmt.Println("")
  fmt.Println("target size: ", targetInt)
  fmt.Println("nonce: " + strconv.Itoa(block.Nonce) + " Hash: " + HashToString(block.Hash))
  return true; //should return true once the hash meets min size
}

func IntToHex(n int64) []byte{
  return []byte(strconv.FormatInt(n, 16))
}
