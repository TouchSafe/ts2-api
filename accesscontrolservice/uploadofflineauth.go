package accesscontrolservice

import (
	"net"
	"time"
)

//OfflineOldAuth not 100% sure of what it does but it appears to be only for cars and using transponer
type OfflineOldAuth struct {
	LocationID     uint16
	TagNumber      uint32
	SequenceNumber uint16
	Entities       []OldEntity
}

//OfflineOldEntity contains info for each entity (inside the car?) being sent
type OfflineOldEntity struct {
	//AuthType should always be transponder as they are in the car???
	AuthType AuthenticationType
	Time     time.Time
	UserID   uint16
}

//ProcessUploadOfflineAuthDataRequest recieves a request to log using the old authentication type?
// This then logs the accesss attempt for each user to the database (usually, not in this fn)
func ProcessUploadOfflineAuthDataRequest(buf []byte, addr net.Addr) (OfflineOldAuth, error) {
	//TODO: this is so messed up, authdata is parsed differently and ina  weird way. Pins should be incorrect. I don't know what is happening
	// oldAuth := OfflineOldAuth{}
	// expectedHeaderSize := 10
	// if len(buf) < expectedHeaderSize {
	// 	return oldAuth, ErrorDecodePacket{Source: "OfflineOldAuth",
	// 		Reason: "buffer does not even fit header"}
	// }
	// // var cmdRequest = binaryReader.ReadByte();
	// //boardCommand := uint8(buf[0])

	// oldAuth.SequenceNumber = binary.BigEndian.Uint16(buf[1:2])
	// oldAuth.LocationID = binary.LittleEndian.Uint16(buf[3:4])
	// entityCount := uint8(buf[5]) //max 100 don't check though cause the other one didn't check lol
	// timeEpoch := binary.LittleEndian.Uint64(buf[6:13])
	// time := time.Unix(int64(timeEpoch), 0)

	// if int(entityCount)*4 != len(buf[expectedHeaderSize:]) {
	// 	return oldAuth, ErrorDecodePacket{Source: "OfflineOldAuth", Reason: "rest of buffer does not match entity count"}
	// }
	// oldAuth.Entities = make([]OldEntity, entityCount)
	// for i := uint8(0); i < entityCount; i++ {
	// 	oldAuth.Entities[i].AuthType = AuthenticationType(uint8(buf[10+i*3]))      //10, 13, 16
	// 	oldAuth.Entities[i].UserID = binary.BigEndian.Uint16(buf[11+i*3 : 12+i*3]) //[11:12], [14:15], [17:18]
	// }
	return OfflineOldAuth{}, nil
}
