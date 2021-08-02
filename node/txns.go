package node

import(
  "crypto/ecdsa"
)

type Txn struct{
  RecipPubKey ecdsa.PublicKey //!!should be stored as actual ecdsa PubKey
  //SenderPubKey ecdsa.PublicKey //!!should be stored as actual ecdsa PubKey
  Sig []byte
  //Amount int
  TxnID []byte
}

//for hashing purposes, alterantively could use byte arrays
type PubKStr struct {
  X  string //big.Int changed to string to support gob encoding
  Y  string //big.Int changed to string to support gob encoding
  XY []byte
}


//!!Change in arguments needed to support cli and sending
func (w Wallet) NewTxn(recipient ecdsa.PublicKey, prevTxn Txn) Txn {
  tx := Txn{RecipPubKey: recipient} //initialize tx

  txId := TxnHashByte(tx, prevTxn) //get tx hash
  tx.TxnID = txId //add hash to tx

  w.SignTxn(tx, prevTxn) //sender wallet signs the txn

  return tx;
}

//verify txn
func VerifyTxn(prevTxPubKey ecdsa.PublicKey, tx Txn) bool { //prevTxPubKey = sender pubkey
  valid := ecdsa.VerifyASN1(&prevTxPubKey, tx.TxnID, tx.Sig) //&sendW.Pubkey = *ecdsa.PubKey
  return valid;
}



//returns nil txn that can be encoded with gob (no nil interfaces)
func NilTxn() Txn {
  nilKey := NilPubKey()
  nilSig := []byte("nil sig")
  nilID := []byte("nil ID")
  nilTxn := Txn{RecipPubKey: nilKey, Sig: nilSig, TxnID: nilID}

  return nilTxn;

}
