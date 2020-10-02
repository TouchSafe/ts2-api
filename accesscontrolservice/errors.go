package accesscontrolservice

import (
	"errors"
	"fmt"
)

var (
	errEmptyBuf                    = errors.New("buffer is empty")
	errIncorrectBoardCommand       = errors.New("incorrect BoardCommand")
	errIncorrectAuthenticationType = errors.New("incorrect AuthenticationType")
	errNotImplemented              = errors.New("not yet implemented")
)

//ErrorDecodePacket is a wrapper for errors in decoding packets
type ErrorDecodePacket struct {
	Source string
	Reason string
}

func (e ErrorDecodePacket) Error() string {
	return fmt.Sprintf("failed to decode \"%s\" packet, reason:%s", e.Source, e.Reason)
}
