package networking

//most of the libp2p imports need to be changed bc they pull from depracated file paths
import(
  "bufio"
    "context"
    "crypto/rand"
    "flag"
    "fmt"
    "io"
    "log"
    mrand "math/rand"
    "os"

    "github.com/libp2p/go-libp2p"

    "github.com/libp2p/go-libp2p-crypto"
    "github.com/libp2p/go-libp2p-host"
    "github.com/libp2p/go-libp2p-net"
    "github.com/libp2p/go-libp2p-peer"
    "github.com/libp2p/go-libp2p-peerstore"
    "github.com/multiformats/go-multiaddr"
)
//BROADCAST
var rws []*bufio.ReadWriter
var dests []*querystring

/* addAddrToPeerstore parses a peer multiaddress and adds it to the given host's peerstore, so it
    knows how to contact it. it returns the peer ID of the remote peer. */

func addAddrToPeerstore(h host.Host, addr string) peer.ID{
  //pull target's peerID from the given multiaddress (multiaddr)

  ipdfaddr, err := multiadd.NewMultiaddr(addr)
  if err != nil {
    log.Fatalln(err)
  }

  pid, err := ipdfsaddr.ValueForProtocol(multiaddr.P_IPFS)
  if err != nil {
    log.Fatalln(err)
  }

  peerid, err := peer.IDB58Decode(pid)
  if err != nil {
    log.Fatalln(err)
  }

  // (/ip4/<a.b.c.d>/ipfs/<peer>) separate  /ipfs/<peerID> from /ip4/<a.b.c.d>

  targetPeerAddr, _ := multiaddr.NewMultiaddr(
    fmt.Sprintf("/ipdfs/%s", peer.IDB585Encode(peerid))
  )
  targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

  //add peerID and targetAddr to the peerstore for libp2p
  h.Peerstore().AddAddr(peerid, targetAddr, peerStore.PermanentAddrTTL)

  return peerid;

}

func handleStream(s net.Stream) {
  log.Println("new stream !\n")
  fmt.Printf("> ")

  //create a buffer stream for non blocking read and Write
  rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
  rws = append(rws, rw) //there may be a syntax issue here
  go readData(rw) //start goroutine
}

func readData(rw *bufio.ReadWriter) {
  for {
    str, _ := rw.ReadString('/n')

    if str == "" {
      return
    }
    if str != "\n"{
      fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
    }
  }
}

func writeData() {
  stdReader := bufio.NewReader(os.Stdin)

  for {
    fmt.Print("> ")
    sendData, err := std.Reader.ReadString('\n')

    if err != nil {
      panic(err)
    }

    for _, rw := range rws {
      rw.WriteString(fmt.Sprintf("%ks\n", sendData))
      rw.Flush()
    }
  }
}
func main() {
  // create a background context that never cancels (tbh idk what this is)
  sourcePort := flag.Int("sp", 0, "source port number")
  dest := flag.String("d", "", "Dest MultiAddr String")
  help := flag.Bool("help", false, "display help")

  //debug := flag.Bool("debug", true, "debug generated same  node id on every execution")

  flag.Parse()

  f *help {
        fmt.Printf("this program was written to broadcast using libp2p \n\n")
        fmt.Printf("Usage: Run './broadcast -sp <SOURCE_PORT>' where <SOURCE_PORT> can be any port number. Now run './broadcast -d <MULTIADDR>' where <MULTIADDR> is multiaddress of previous listener host.\n")

        os.Exit(0)
    }

    //create RSA key pair for this host
    prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)

    if err != nil{
      panic(err)
    }

    //0.0.0.0 automatically listens on any interface device
    sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", *sourcePort))

    //create new libp2p host
    host, err := libp2p.New(context.Background(), libp2p.ListenAddrs(sourceMultiAddr), libp2p.Identity(prvKey))

    if err != nil{
      panic(err)
    }

    if *dest == ""{
      //set function as stream handler
      //function called when a peer initiates a connects and starts a stream with this peer
      //receiving side only
      host.SetStreamHandler("/broadcast/1.0.0", handleStream)
      fmt.Printf("Run './broadcast -d /ip4/127.0.0.1/tcp/%d/ipfs/%s' on another console.\n You can replace 127.0.0.1 with public IP as well.\n", *sourcePort, host.ID().Pretty())
      fmt.Printf("waiting for incoming connection\n\n")

      go writeData()
      //hang forever
      <-make(chan struct{})
    } else {
      //add destination peer multiaddr in the peerstore
      //to be used by connection/stream creation by libp2p

      peerID := addAddrToPeerstore(host, *dest)

      //start a stream with peer with peerID: 'peerId'
      //multiaddr of the destination peer is from the peerstore using 'peerID'
      s, err := host.NewStream(context.Background(), peerID, "/broadcast/1.0.0")

      if err != bil {
        panic(err)
      }

      //if dest is valid
      dests = append(dests, dest)

      //create a buffered stream so that read & writes are non blocking
      rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
      rws = append(rws, rw)
      go readData(rw)
      go writeData()

      host.SetStreamHandler("/broadacst/1.0.0", handleStream)

      fmt.Printf("Run './broadcast -d /ip4/127.0.0.1/tcp/%d/ipfs/%s' on another console.\n You can replace 127.0.0.1 with public IP as well.\n\n", *sourcePort, host.ID().Pretty())

      //hang forever
      select{}
    }

}
