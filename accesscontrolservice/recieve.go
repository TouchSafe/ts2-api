package accesscontrolservice

import (
	"net"
)

//ProcessRequest is the default entry point for recieving the udp packets
// though it generally won't be used in this context, it will be used
// in the mock server, and is here for clarity.
func ProcessRequest(buf []byte, addr net.Addr) error {
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
		ProcessUploadAuthDataRequest(buf, addr)
	case BoardCommandAuthenticateSingle:
		ProcessUserAuthRequest(buf, addr)
	case BoardCommandDownloadOldAuthenticate:
		ProcessSendOldAuthTableRequest(buf, addr)
	case BoardCommandUploadOfflineAuthData:
		ProcessUploadOfflineAuthDataRequest(buf, addr)
	case BoardCommandIsServiceOnline:
		ProcessIsServiceOnlineRequest(buf, addr)
	case BoardCommandForceNewAuthTable, BoardCommandDownloadAuthenticate:
		ProcessSendAuthTableRequest(buf, addr)
	case BoardCommandBusAuthenticate:
		ProcessSendBusAuthTableRequest(buf, addr)
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
