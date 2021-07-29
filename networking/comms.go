package networking


import (
  //"bufio"
	"context"
	//"crypto/rand"
	//"flag"
	"fmt"
	//"io"
	//"log"
	//mrand "math/rand"
	//"os"
  //"roCoin/node"

	"github.com/libp2p/go-libp2p"
	//"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	//"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery"
  "time"
  pubsub "github.com/libp2p/go-libp2p-pubsub"

	//"github.com/multiformats/go-multiaddr"
)
//how often mDNS records are republished
const DiscoveryInterval = time.Hour

//to be use in mDNS advertisements to discover other chat peers
const DiscoveryServiceTag = "roCoin-net"

func Setup(ctx context.Context) (host.Host, *pubsub.PubSub, error) { //add argument for specific port
  //create new libp2p Host listening on a random tcp port
  h, err := libp2p.New(ctx, libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
  if err != nil {
    panic(err)
    return nil, nil, err
  }

  //create a new pubsub service
  ps, err := pubsub.NewGossipSub(ctx, h)
  if err != nil {
    panic(err)
    return nil, nil, err
  }
  return h, ps, nil;
}

//returns the last 8 chars of a base58-encoded peer id
func ShortID(p peer.ID) string {
  pretty := p.Pretty()
  return pretty[len(pretty)-8:]
}

//discoveryNotifee gets notified when we find a new peer via mDNS discovery
type DiscoveryNotifee struct {
  h host.Host
}

//connects to peers discovered via mDNS, once connected PubSub will begin interacting if they support
func (n *DiscoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
  fmt.Printf("discoverd new peer %s\n", pi.ID.Pretty())
  err := n.h.Connect(context.Background(), pi)
  if err != nil {
    fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
  }
  fmt.Printf("connected to peer %s: %s\n", pi.ID.Pretty())
}

//setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host
//allows automatic discovery of peers on same LAN and connection
func SetupDiscovery(ctx context.Context, h host.Host) error {
  //setup mDNS discovery to find local peers
  disc, err := discovery.NewMdnsService(ctx, h, DiscoveryInterval, DiscoveryServiceTag)
  if err != nil {
    return err
  }

  n := DiscoveryNotifee{h: h}
  disc.RegisterNotifee(&n)
  return nil
}
