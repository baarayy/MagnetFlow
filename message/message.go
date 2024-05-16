package message

import "encoding/binary"

type messageID uint8

const (
	MsgChoke         messageID = 0
	MsgUnchoke       messageID = 1
	MsgInterested    messageID = 2
	MsgNotInterested messageID = 3
	MsgHave          messageID = 4
	MsgBitfield      messageID = 5
	MsgRequest       messageID = 6
	MsgPiece         messageID = 7
	MsgCancel        messageID = 8
)

type Message struct {
	ID      messageID
	payload []byte
}

func (m *Message) Serialize() []byte {
	if m == nil {
		return make([]byte, 0)
	}
	length := uint32(len(m.payload) + 1) // 1 byte for ID
	buf := make([]byte, length+4)        // 4 bytes for length
	binary.BigEndian.PutUint32(buf[0:4], length)
	buf[4] = byte(m.ID)
	copy(buf[5:], m.payload)
	return buf
}
