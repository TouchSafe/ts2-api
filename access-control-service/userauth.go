package access_control_service

import (
	"encoding/binary"
	"errors"
	"net"
)

type UserAuth struct {
	BoardCommand       BoardCommand
	SequenceNumber     uint16
	AccessPointID      uint16
	AuthenticationData Auth
}

type Auth struct {
	AuthenticationType AuthenticationType
	AuthenticationData string
}

func ReceiveUserAuthCommand(buf []byte, addr net.Addr) {

}

func DecodeUserAuthPacket(buf []byte) (UserAuth, error) {
	expectedHeaderSize := uint(8)
	boardCommand := uint(buf[0])
	headerSize := uint(buf[1])
	if BoardCommand(boardCommand) != BoardCommandAuthenticateSingle || headerSize != expectedHeaderSize || uint(len(buf)) <= expectedHeaderSize {
		return UserAuth{}, errors.New("failure to decode user auth packet, packet size or specifications incorrect")
	}
	sequenceNumber := binary.LittleEndian.Uint16(buf[2:3])
	accessPointID := binary.LittleEndian.Uint16(buf[4:5])
	requestAuthType := uint(buf[6])
	requestDataLength := uint(buf[7])
	authenticationData := buf[8:]
	if requestDataLength != uint(len(authenticationData)) {
		return UserAuth{}, errors.New("failure to decode user auth packet, specified data length does not match data length recieved")
	}
	userAuth := UserAuth{
		BoardCommand:   BoardCommand(boardCommand),
		SequenceNumber: sequenceNumber,
		AccessPointID:  accessPointID,
		AuthenticationData: Auth{
			AuthenticationType: AuthenticationType(requestAuthType),
		},
	}

	return userAuth, nil
}
