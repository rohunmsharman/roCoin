package main

import(
"testing"
"roCoin/node"
)

//connect test w/ test TxnStream
/*
ctx := context.Background()
// h is host, ps is pubsub from gossibRouter, err is error
h, ps, err := networking.Setup(ctx)
if err != nil {
  panic(err)
}

//setup local mDNS discovery
err = networking.SetupDiscovery(ctx, h)
if err != nil {
  panic(err)
}

ts, err := networking.JoinTxnStream(ctx, ps, h.ID(), txnStrm)
if err != nil {
  panic(err)
}


//testing DO NOT DELETE
//the connect works, broadcasting the TXNs works, just need to create better handling
var in int
fmt.Println("peer 1 or 2?")
fmt.Scanln(&in)

tstTxns := []node.TestTxn{}
if in == 1 {
  fmt.Println("send test txns? type 1 again")
  fmt.Scanln(&in)
  if in == 1 {
    tstTxns = append(tstTxns, node.TestTxn{Sender: "bob", Recipient: "alice", Amount: 30})
    tstTxns = append(tstTxns, node.TestTxn{Sender: "alice", Recipient: "bob", Amount: 40})
    tstTxns = append(tstTxns, node.TestTxn{Sender: "bob", Recipient: "alice", Amount: 15})
    for _, txn := range tstTxns {
      ts.Publish(txn)
    }
  }

}

if in == 2 {

  //tstTxns = append(tstTxns, node.TestTxn{Sender: "michelle", Recipient: "britney", Amount: 10})
  //tstTxns = append(tstTxns, node.TestTxn{Sender: "britney", Recipient: "michelle", Amount: 20})
  //tstTxns = append(tstTxns, node.TestTxn{Sender: "michelle", Recipient: "britney", Amount: 90})

  for {
    go ts.HandleEvents()
  }

}
*/

/*
func GenTest(t *testing.T) {
  /*
  Should test txn, mining, etc...
  
  return
}
*/

func TxnMarshalTest(t *testing.T) {
  //txn marshalling test
  bob := node.NewWallet("bob", 500)
  alice := node.NewWallet("alice", 600)

  tx1 := node.CreateTxn(bob, alice, 300)
  bob.SignTxn(tx1)

  fmt.Println("marshalling txn")
  marshalTx1 := tx1.MarshalTx()
  unmarshalTx1 := node.UnmarshalTx(marshalTx1)

  fmt.Println(unmarshalTx1)
  if unmarshalTx1 != tx1 {
    t.Errorf("failed to marshal txn properly!")
    switch unmarshalTx1 {
    case: tx1.RecipPubKey != unmarshalTx1.RecipPubKey:
      t.Errorf("recipient public keys do not match")
    case: tx1.SenderPubKey != unmarshalTx1.SenderPubKey:
      t.Errorf("sender public keys do not match")
    case: tx1.Sig != unmarshalTx1.Sig:
      t.Errorf("signatures do not match")
    case: tx1.Amount != unmarshalTx1.Amount:
      t.Errorf("amounts do not match")
    case: tx1.TxnID != unmarshalTx1.TxID:
      t.Errorf("txn IDs do not match")
    }
  }
}
