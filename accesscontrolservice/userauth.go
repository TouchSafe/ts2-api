package accesscontrolservice

import (
	"encoding/binary"
	"net"
)

//UserAuth contains data for the current requested authentication
type UserAuth struct {
	BoardCommand       BoardCommand
	SequenceNumber     uint16
	AccessPointID      uint16
	AuthenticationData Auth
}

//Auth is a basic authentication struct
type Auth struct {
	AuthenticationType AuthenticationType
	AuthenticationData string
}

//ProcessUserAuthRequest handles recieving the SINGLE user auth packet, this may not be 100% necessary
func ProcessUserAuthRequest(buf []byte, addr net.Addr) {

}

//DecodeUserAuthPacket decodes the SINGLE user authentication packet recieved
func DecodeUserAuthPacket(buf []byte) (UserAuth, error) {
	var err error
	expectedHeaderSize := uint(8)
	boardCommand := uint(buf[0])
	headerSize := uint(buf[1])
	if BoardCommand(boardCommand) != BoardCommandAuthenticateSingle || headerSize != expectedHeaderSize || uint(len(buf)) <= expectedHeaderSize {
		return UserAuth{}, ErrorDecodePacket{Source: "SingleUserAuth",
			Reason: "header or board command issue"}
	}
	sequenceNumber := binary.LittleEndian.Uint16(buf[2:3])
	accessPointID := binary.LittleEndian.Uint16(buf[4:5])
	requestAuthType := uint(buf[6])
	requestDataLength := uint(buf[7])
	authenticationData := buf[8:]
	if requestDataLength != uint(len(authenticationData)) {
		return UserAuth{}, ErrorDecodePacket{Source: "SingleUserAuth",
			Reason: "specified data length does not match data length recieved"}
	}
	userAuth := UserAuth{
		BoardCommand:   BoardCommand(boardCommand),
		SequenceNumber: sequenceNumber,
		AccessPointID:  accessPointID,
		AuthenticationData: Auth{
			AuthenticationType: AuthenticationType(requestAuthType),
		},
	}
	userAuth.AuthenticationData.AuthenticationData, err = authDataToString(AuthenticationType(requestAuthType), authenticationData)
	return userAuth, err
}
