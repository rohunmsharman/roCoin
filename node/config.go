package node
import (
  //"fmt"
)

type ConfigVars struct {
  TxStream string //txn stream
  //Port int //port to connect to
  //... settings
}

func (cv ConfigVars) SetTxStream(streamName string) { //rewrite to save as a txt or similar holding settings line by line
  cv.TxStream = streamName
  //eventually add ability to autosave to gob file
}
