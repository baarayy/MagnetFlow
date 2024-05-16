package client

import (
	"magnetflow/bitfield"
	"magnetflow/peers"
	"net"
)

type Client struct {
	conn     net.Conn
	choked   bool
	Bitfield bitfield.Bitfield
	peer     peers.Peer
	infoHash [20]byte
	PeerID   [20]byte
}
