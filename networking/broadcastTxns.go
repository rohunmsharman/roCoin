package networking

import (
  "fmt"
  "time"
  "context"
  "roCoin/node"
  "github.com/libp2p/go-libp2p-core/peer"
  "os/signal"
  pubsub "github.com/libp2p/go-libp2p-pubsub"
)

/*
this implementation of p2p pubsub is probably not the best, i basically rewrote github.com/libp2p/go-libp2p/tree/master/examples/pubsub/chat
w/ my own gui. using libp2p seems like it'll be easier than hardcoding a networking protocol, but this whole thing is tbd
*/
const TxnStreamBufSize = 128
//TxnStream is a subscription to a single PubSub topic, txns can be published to the topic.
//txns are pushed to the Txns channel
type TxnStream struct{
  //Transactions is a channel of txns received from peers
  Transactions chan *node.Txn

  ctx context.Context
  ps *pubsub.PubSub
  topic *pubsub.Topic
  sub *pubsub.Subscription

  TxnStreamName string
  self peer.ID
}

func JoinTxnStream(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, streamName string) (*TxnStream, error) {
  //join the pubsub topic
  topic, err := ps.Join(TopicName(streamName))
  if err != nil {
    return nil, err
  }

  //subscribe to the topic
  sub, err := topic.Subscribe()
  if err != nil {
    return nil, err
  }

  ts := &TxnStream {
    ctx: ctx,
    ps: ps,
    topic: topic,
    sub: sub,
    self: selfID,
    TxnStreamName: streamName,
    Transactions: make(chan *node.Txn, TxnStreamBufSize), //change back to normal txn after testing
  }

  //start reading messages from the subscription in a loop
  go ts.ReadLoop()
  return ts, nil;
}

//Publish sends txn to PubSub topic
func (ts *TxnStream) Publish(txnToPub node.Txn) error {
  txBytes := txnToPub.MarshalTx()
  fmt.Println("publishing txn: ", txnToPub.TxnID)

  return ts.topic.Publish(ts.ctx, txBytes);
}

/*
//PUBLISH FOR TEST TXN
func (ts *TxnStream) Publish(txnToPub node.TestTxn) error {
  //TxnToJson outputs byte array of txn
  //txBytes := node.TxnToJson(txnToPub) //commented out for testing
  txBytes := node.TxnToJson(txnToPub)
  fmt.Println("publishing txn")
  return ts.topic.Publish(ts.ctx, txBytes)
}
*/
//loop pulls txns from the pubsub topic and pushes them onto the Transactions channel
func (ts *TxnStream) ReadLoop() {
  for {
    txMsg, err := ts.sub.Next(ts.ctx) //Next() returns txn as type ipubsub.Message
    if err != nil {
      close(ts.Transactions)
      return
    }
    //only forward txns delivered by others
    if txMsg.ReceivedFrom == ts.self {
      continue
    }

    tx := node.UnmarshalTx(txMsg.Data) //txMsg.Data from ipubusub.Message struct
    //err = json.Unmarshal(txMsg.Data, tx)

    //send unmarshaled valid messages onto the Transactions channel
    ts.Transactions <- &tx //used to be just tx, im only kinda sure why & makes it a *Txn
  }
}
//handles what to do on Transactions channel
func (ts *TxnStream) HandleEvents() txn {
  c := make(chan os.Signal, 1) //channel to handle os Signal
  refreshTicker := time.NewTicker(time.Second)
  defer refreshTicker.Stop()
  //InTxns := []node.Txn{}
  for {
    //maybe add a sleep?
    //this handler needs to be edited!
    select{
      case inTx := <- ts.Transactions:
        fmt.Println("new txn received")
        //InTxns = append(InTxns, *inTx)
        return txn
      }
    }
    //returns array of all txns from ts
   return InTxns;
}


func (ts *TxnStream) ListPeers() []peer.ID {
  return ts.ps.ListPeers(TopicName(ts.TxnStreamName))
}

func TopicName(txnStreamName string) string {
  return "txn-stream:" + txnStreamName
}
