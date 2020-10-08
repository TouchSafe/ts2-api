package accesscontrolservice

import (
	"encoding/binary"
	"net"
)

//ProcessSendAuthTableRequest gets a new auth table
func ProcessSendAuthTableRequest(buf []byte, addr net.Addr) (SequenceNumber uint16, err error) {
	return processAuthTableRequest(buf, addr)
}

//ProcessSendOldAuthTableRequest sends the old authentication table
func ProcessSendOldAuthTableRequest(buf []byte, addr net.Addr) (SequenceNumber uint16, err error) {
	return processAuthTableRequest(buf, addr)
}

//ProcessSendBusAuthTableRequest sends the bus auth table??? just a new layout
func ProcessSendBusAuthTableRequest(buf []byte, addr net.Addr) (SequenceNumber uint16, err error) {
	return processAuthTableRequest(buf, addr)
}

func processAuthTableRequest(buf []byte, addr net.Addr) (SequenceNumber uint16, err error) {
	expectedHeaderLength := 3
	if len(buf) != expectedHeaderLength {
		return 0, ErrorDecodePacket{
			Source: "SendAuthTable",
			Reason: "buffer does not even fit header",
		}
	}
	return binary.LittleEndian.Uint16(buf[1:]), nil
}
