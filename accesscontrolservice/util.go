package accesscontrolservice

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func authDataToString(authType AuthenticationType, authData []byte) (string, error) {
	switch AuthenticationType(authType) {
	case AuthenticationTypePin:
		pin := ""
		for i := range authData {
			digit, err := ByteArrayToIntString(authData[i : i+1])
			if err != nil {
				//This error will never be reached due to passing in one at a time
				return "", ErrorDecodePacket{Source: "RFID decode", Reason: err.Error()}
			}
			pin += digit
		}
		return pin, nil
	case AuthenticationTypeRFID:
		if len(authData) != 3 {
			return "", ErrorDecodePacket{Source: "RFID decode", Reason: fmt.Sprintf("authData len expected 3 bytes, but is %d", len(authData))}
		}
		//Swap the bytes 0,2 keep byte 1 (reverse order)
		authData[0], authData[2] = authData[2], authData[0]
		result, err := ByteArrayToIntString(authData)
		if err != nil {
			//This error will never be reached as len is checked before passing though
			// and the function only fails on lenths too large
			return "", ErrorDecodePacket{Source: "RFID decode", Reason: err.Error()}
		}
		return result, nil
	case AuthenticationTypeTransponder:
		if len(authData) != 2 {
			return "", ErrorDecodePacket{Source: "Transponder decode", Reason: fmt.Sprintf("authData len expected 2 bytes, but is %d", len(authData))}
		}
		//Swap the bytes 0,2 keep byte 1 (reverse order)
		authData[0], authData[1] = authData[1], authData[0]
		result, err := ByteArrayToIntString(authData)
		if err != nil {
			//This will never be reached, as it will only get here with authdata len of 2
			//The function only errors on authdata len >8
			return "", ErrorDecodePacket{Source: "Transponder decode", Reason: err.Error()}
		}
		return result, nil
	case AuthenticationTypeErased:
		fallthrough
	case AuthenticationTypeFingerprint:
		fallthrough
	case AuthenticationTypePassword:
		fallthrough
	case AuthenticationTypeManualOverride:
		return "", errNotImplemented
	default:
		return "", errIncorrectAuthenticationType
		//unexpected auth type
	}
}

//ByteArrayToIntString returns the string representation of the int
// stored in the byte array
// (00000001) = "1"
// (00010000) = "16"
// (1111000000001111) = "4027514880"
func ByteArrayToIntString(array []byte) (string, error) {
	if len(array) > 8 {
		return "0", errors.New("array too large, max len 8")
	}
	array = append(make([]byte, 8-len(array)), array...)
	got := binary.BigEndian.Uint64(array)
	return fmt.Sprintf("%d", int64(got)), nil
}
