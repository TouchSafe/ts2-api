package accesscontrolservice

import (
	"encoding/binary"
	"net"
)

//OldAuth not 100% sure of what it does but it appears to be only for cars and using transponer
type OldAuth struct {
	LocationID     uint16
	TagNumber      uint32
	SequenceNumber uint16
	Entities       []OldEntity
}

//OldEntity contains info for each entity (inside the car?) being sent
type OldEntity struct {
	//AuthType should always be transponder as they are in the car???
	AuthType AuthenticationType
	UserID   uint16
}

//ProcessUploadAuthDataRequest recieves a request to log using the old authentication type?
// This then logs the accesss attempt for each user to the database (usually, not in this fn)
func ProcessUploadAuthDataRequest(buf []byte, addr net.Addr) (OldAuth, error) {
	oldAuth := OldAuth{}
	expectedHeaderSize := 10
	if len(buf) < expectedHeaderSize {
		return oldAuth, ErrorDecodePacket{Source: "OldAuth",
			Reason: "buffer does not even fit header"}
	}
	// var cmdRequest = binaryReader.ReadByte();
	//	boardCommand := uint8(buf[0])
	// var locationId = binaryReader.ReadUInt16();
	oldAuth.LocationID = binary.LittleEndian.Uint16(buf[1:2])
	// var tagNumber = binaryReader.ReadUInt32();
	oldAuth.TagNumber = binary.LittleEndian.Uint32(buf[3:6])
	// var seqNumber = WordSwizzle(binaryReader.ReadUInt16());
	oldAuth.SequenceNumber = binary.BigEndian.Uint16(buf[7:8])
	// var entityCount = binaryReader.ReadByte();
	entityCount := uint8(buf[9])
	if int(entityCount)*3 != len(buf[10:]) {
		return oldAuth, ErrorDecodePacket{Source: "OldAuth", Reason: "rest of buffer does not match entity count"}
	}
	oldAuth.Entities = make([]OldEntity, entityCount)
	for i := uint8(0); i < entityCount; i++ {
		oldAuth.Entities[i].AuthType = AuthenticationType(uint8(buf[10+i*3]))      //10, 13, 16
		oldAuth.Entities[i].UserID = binary.BigEndian.Uint16(buf[11+i*3 : 12+i*3]) //[11:12], [14:15], [17:18]
	}
	return oldAuth, nil
}
