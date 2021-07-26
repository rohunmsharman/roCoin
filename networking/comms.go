package networking

import (
  "context"
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "github.com/libp2p/go-libp2p"
  peerstore "github.com/libp2p/go-libp2p-core/peer"
  "github.com/libp2p/go-libp2p/p2p/protocol/ping"
  multiaddr "github.com/multiformats/go-multiaddr"
)

func GenComms() {
  //add array of string Addr s to listen to
  ctx := context.Background()

  cNode, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))

  //add connection manager for controlling how many connections a peer has
}
