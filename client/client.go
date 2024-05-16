package client

import (
	"magnetflow/bitfield"
	"magnetflow/message"
	"magnetflow/peers"
	"net"
)

type Client struct {
	Conn     net.Conn
	Choked   bool
	Bitfield bitfield.Bitfield
	peer     peers.Peer
	infoHash [20]byte
	PeerID   [20]byte
}

func (c *Client) Read() (*message.Message, error) {
	msg, err := message.Read(c.Conn)
	return msg, err
}
