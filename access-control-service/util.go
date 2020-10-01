package access_control_service

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func AuthDataToString(authType AuthenticationType, authData []byte) (string, error) {
	switch AuthenticationType(authType) {
	case AuthenticationTypePin:
		//there is no way the ping is longer than 64 characters lol
		pin := ""
		for i := range authData {
			digit, err := ByteArrayToIntString(authData[i : i+1])
			if err != nil {
				return "", errors.New("Failed to decode pin: " + err.Error())
			}
			pin += digit
		}
		return pin, nil
	case AuthenticationTypeRFID:
		if len(authData) != 3 {
			return "", fmt.Errorf("authentication data for RFID should be 3 bytes, but is %d", len(authData))
		}
		//Swap the bytes 0,2 keep byte 1 (reverse order)
		authData[0], authData[2] = authData[2], authData[0]
		result, err := ByteArrayToIntString(authData)
		if err != nil {
			return "", errors.New("Failed to decode RFID: " + err.Error())
		}
		return result, nil
	case AuthenticationTypeTransponder:
		fallthrough
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
	got, n := binary.Uvarint(array)
	if n <= 0 {
		return "", errors.New("byte array to string failed to read the array")
	}
	return fmt.Sprintf("%d", int64(got)), nil
}
