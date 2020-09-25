package access_control_service

import (
	"errors"
	"net"
)

var (
	errEmptyBuf              = errors.New("buffer is empty")
	errIncorrectBoardCommand = errors.New("incorrect BoardCommand ")
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
	case BoardCommandNone:
	case BoardCommandUploadAuthData:
	case BoardCommandUploadAuthDataAck:
	case BoardCommandAuthenticateSingle:
		ReceiveUserAuthCommand(buf, addr)
	case BoardCommandDownloadOldAuthenticate:
	case BoardCommandDownloadOldAuthenticateAck:
	case BoardCommandDownloadOldAuthenticateData:
	case BoardCommandAuthenticateResponse:
	case BoardCommandAuthenticateAck:
	case BoardCommandUploadOfflineAuthData:
	case BoardCommandUploadOfflineAuthDataAck:
	case BoardCommandDownloadAuthenticate:
	case BoardCommandDownloadAuthenticateAck:
	case BoardCommandDownloadAuthenticateData:
	case BoardCommandIsServiceOnline:
	case BoardCommandBoardCommandIsServiceOnlineAck:
	case BoardCommandNoNewAuthTable:
	case BoardCommandForceNewAuthTable:
	case BoardCommandBusAuthenticateAck:
	case BoardCommandBusAuthenticate:
	case BoardCommandBusAuthenticateData:
	default:
		return errIncorrectBoardCommand

	}
	return nil
}

//TODO: Sep into other files here

func ReceiveUserAuthCommand(buf []byte, addr net.Addr) {

}
