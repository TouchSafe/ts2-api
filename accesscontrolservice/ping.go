package accesscontrolservice

import (
	"encoding/binary"
	"net"
)

//ProcessIsServiceOnlineRequest is just a ping pong check
func ProcessIsServiceOnlineRequest(buf []byte, addr net.Addr) (SequenceNumber uint16, err error) {
	//Note this is same as sendAuthTableCommand
	expectedHeaderLength := 3
	if len(buf) != expectedHeaderLength {
		return 0, ErrorDecodePacket{
			Source: "SendAuthTable",
			Reason: "buffer does not even fit header",
		}
	}
	return binary.LittleEndian.Uint16(buf[1:]), nil
}
