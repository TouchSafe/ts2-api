package access_control_service

import (
	"errors"
	"net"
)

var (
	errEmptyBuf                    = errors.New("buffer is empty")
	errIncorrectBoardCommand       = errors.New("incorrect BoardCommand")
	errIncorrectAuthenticationType = errors.New("incorrect AuthenticationType")
	errNotImplemented              = errors.New("not yet implemented")
)

//Receive is the default entry point for recieving the udp packets
// though it generally won't be used in this context, it will be used
// in the mock server, and is here for clarity.
func Receive(buf []byte, addr net.Addr) error {
	if len(buf) == 0 {
		return errEmptyBuf
	}
	//After error checking the original source sends and ack if the connection is still there.
	//This I believe generally shouldn't be happening as why would they be sending data again if
	// if the application is still processing it? Unless it's sending the same data.. In which
	// case we should be checking the status of it and make sure it hasn't timed out
	//TODO: Make a timeout
	//TODO: Put ack here, this is implementation detail not API?

	switch BoardCommand(buf[0]) {
	case BoardCommandUploadAuthData:
		RecieveUploadAuthDataCommand(buf, addr)
	case BoardCommandAuthenticateSingle:
		ReceiveUserAuthCommand(buf, addr)
	case BoardCommandDownloadOldAuthenticate:
		SendOldAuthTableCommand(buf, addr)
	case BoardCommandUploadOfflineAuthData:
		RecieveUploadOfflineAuthDataCommand(buf, addr)
	case BoardCommandIsServiceOnline:
		IsServiceOnlineCommand(buf, addr)
	case BoardCommandForceNewAuthTable, BoardCommandDownloadAuthenticate:
		SendAuthTableCommand(buf, addr)
	case BoardCommandBusAuthenticate:
		SendBusAuthTableCommand(buf, addr)
	case BoardCommandNone:
		fallthrough
	case BoardCommandUploadAuthDataAck:
		fallthrough
	case BoardCommandDownloadOldAuthenticateAck:
		fallthrough
	case BoardCommandDownloadOldAuthenticateData:
		fallthrough
	case BoardCommandAuthenticateResponse:
		fallthrough
	case BoardCommandAuthenticateAck:
		fallthrough
	case BoardCommandUploadOfflineAuthDataAck:
		fallthrough
	case BoardCommandDownloadAuthenticateAck:
		fallthrough
	case BoardCommandDownloadAuthenticateData:
		fallthrough
	case BoardCommandBoardCommandIsServiceOnlineAck:
		fallthrough
	case BoardCommandNoNewAuthTable:
		fallthrough
	case BoardCommandBusAuthenticateAck:
		fallthrough
	case BoardCommandBusAuthenticateData:
		return errNotImplemented
	default:
		return errIncorrectBoardCommand
	}
	return nil
}

//TODO: Sep into other files here

func RecieveUploadAuthDataCommand(buf []byte, addr net.Addr)        {}
func RecieveUploadOfflineAuthDataCommand(buf []byte, addr net.Addr) {}
func SendAuthTableCommand(buf []byte, addr net.Addr)                {}
func SendOldAuthTableCommand(buf []byte, addr net.Addr)             {}
func SendBusAuthTableCommand(buf []byte, addr net.Addr)             {}
func IsServiceOnlineCommand(buf []byte, addr net.Addr)              {}
