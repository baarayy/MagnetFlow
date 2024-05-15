package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

func Unmarshal(peersBin []byte) ([]Peer, error) {
	const peerSize = 6 // 4 bytes for IP, 2 bytes for port
	numPeers := len(peersBin) / peerSize
	if len(peersBin)%peerSize != 0 {
		return nil, fmt.Errorf("Received malformed peers")
	}
	peers := make([]Peer, numPeers)

	for i := 0; i < numPeers; i++ {
		offSet := i * peerSize
		peers[i].IP = net.IP(peersBin[offSet : offSet+4])
		peers[i].Port = binary.BigEndian.Uint16([]byte(peersBin[offSet+4 : offSet+6]))
	}
	return peers, nil
}

func (p Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}
