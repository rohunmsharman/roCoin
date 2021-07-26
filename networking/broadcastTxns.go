package networking

import (
  "fmt"
  "context"
  "roCoin/node"
  "github.com/libp2p/go-libp2p-core/peer"
  pubsub "github.com/libp2p/go-libp2p-pubsub"

)
const TxnStreamBufSize = 128
//TxnStream is a subscription to a single PubSub topic, txns can be published to the topic.
//txns are pushed to the Txns channel
type TxnStream struct{
  //Transactions is a channel of txns received from peers
  Transactions chan *node.Txn

  ctx context.Context
  ps *pubsub.PubSub
  topic *pubsub.topicsub *pubsub.Subscription

  TxnStreamName string
  self peer.ID
}

func JoinTxnStream(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, streamName string) (*TxnStream, error) {
  //join the pubsub topic
  topic, err := ps.Join(topicName(streamName))
  if err != nil {
    return nil, err
  }

  //subscribe to the topic
  sub, err := topic.Subscibe()
  if err != nil {
    return nil, err
  }

  ts := &TxnStream {
    ctx: ctx,
    ps: ps,
    topic: topic,
    sub: sub,
    self: selfID,
    streamName: streamName,
    Transactions: make(chan *node.Txns, TxnStreamBufSize),
  }

  //start reading messages from the subscription in a loop
  go ts.ReadLoop()
  return ts, nil;
}

//Publish sends txn to PubSub topic
func (ts *TxnStream) Publish(txnToPub node.Txn) error {
  //TxnToJson outputs byte array of txn
  txBytes := node.TxnToJson(txnToPub)
  return ts.topic.Publish(ts.ctx, txBytes)
}

//loop pulls txns from the pubsub topic and pushes them onto the Transactions channel
func (ts *TxnStream) ReadLoop() {
  for {
    txMsg, err := ts.sub.Next(ts.ctx) //Next() returns type message from PubSub
    if err != nil {
      close(ts.Transactions)
      return
    }
    //only forward txns delivered by others
    if txMsg.ReceivedFrom == ts.self {
      continue
    }

    tx := new(node.Txn)
    err = json.Unmarshal(txMsg.Data, tx)
    if err != nil {
      continue
    }
    //send unmarshaled valid messages onto the Transactions channel
    ts.Transactions <- tx
  }
}


func (ts *TxnStream) ListPeers() []peer.ID {
  return ts.ps.ListPeers(TopicName(ts.TxnStreamName))
}

func TopicName(txnStreamName string) string {
  return "txn-stream:" + txnStreamName
}
