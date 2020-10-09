package accesscontrolservice

import (
	"encoding/binary"
	"net"
	"strconv"
)

//OldEntity contains info for each entity (inside the car?) being sent
type OldEntity struct {
	//AuthType should always be transponder as they are in the car???
	AuthType AuthenticationType
	UserID   uint16
}

//OldAuth not 100% sure of what it does but it appears to be only for cars and using transponer
type OldAuth struct {
	LocationID     uint16
	TagNumber      uint32
	SequenceNumber uint16
	Entities       []OldEntity
}

//EncodeTS2 encodes the struct for TS2 network communication
func (oe OldEntity) EncodeTS2() []byte {
	expectedHeaderSize := 3
	buf := make([]byte, expectedHeaderSize)
	buf[0] = byte(oe.AuthType)
	binary.BigEndian.PutUint16(buf[1:3], oe.UserID)
	return buf
}

func (oe OldEntity) toString() string {
	return "OldEntity{" + oe.AuthType.String() + "," + strconv.Itoa(int(oe.UserID)) + "}"
}

//EncodeTS2 encodes the struct for TS2 network communication
func (oa OldAuth) EncodeTS2() []byte {
	expectedHeaderSize := 10
	buf := make([]byte, expectedHeaderSize)
	buf[0] = byte(BoardCommandUploadAuthData)
	binary.LittleEndian.PutUint16(buf[1:3], oa.LocationID)

	binary.LittleEndian.PutUint32(buf[3:7], oa.TagNumber)
	// var seqNumber = WordSwizzle(binaryReader.ReadUInt16());
	binary.BigEndian.PutUint16(buf[7:9], oa.SequenceNumber)
	// var entityCount = binaryReader.ReadByte();
	buf[9] = byte(len(oa.Entities))
	for _, e := range oa.Entities {
		buf = append(buf, e.EncodeTS2()...)
	}
	return buf
}

//ProcessUploadAuthDataRequest recieves a request to log using the old authentication type?
// This then logs the accesss attempt for each user to the database (usually, not in this fn)
func ProcessUploadAuthDataRequest(buf []byte, addr net.Addr) (OldAuth, error) {
	oa := OldAuth{}
	expectedHeaderSize := 10
	if len(buf) < expectedHeaderSize {
		return oa, ErrorDecodePacket{Source: "OldAuth",
			Reason: "buffer does not even fit header"}
	}
	// var cmdRequest = binaryReader.ReadByte();
	//	boardCommand := uint8(buf[0])
	// var locationId = binaryReader.ReadUInt16();
	oa.LocationID = binary.LittleEndian.Uint16(buf[1:3])
	// var tagNumber = binaryReader.ReadUInt32();
	oa.TagNumber = binary.LittleEndian.Uint32(buf[3:7])
	// var seqNumber = WordSwizzle(binaryReader.ReadUInt16());
	oa.SequenceNumber = binary.BigEndian.Uint16(buf[7:9])
	// var entityCount = binaryReader.ReadByte();
	entityCount := uint8(buf[9])
	if int(entityCount)*3 != len(buf[10:]) {
		return oa, ErrorDecodePacket{Source: "OldAuth", Reason: "rest of buffer does not match entity count"}
	}
	oa.Entities = make([]OldEntity, entityCount)
	for i := uint8(0); i < entityCount; i++ {
		oa.Entities[i].AuthType = AuthenticationType(uint8(buf[10+i*3]))      //10, 13, 16
		oa.Entities[i].UserID = binary.BigEndian.Uint16(buf[11+i*3 : 13+i*3]) //[11:13], [14:16], [17:19]
	}
	return oa, nil
}
