package accesscontrolservice

import (
	"errors"
	"strings"
	"testing"
)

func TestAuthDataToStringCorrect(t *testing.T) {
	//THIS HAD RESULTS CHECKED MANUALLY AGAINST C# CODE
	inAuthType := []AuthenticationType{
		AuthenticationTypePin,
		AuthenticationTypeRFID,
	}
	want := []string{
		"123",
		"197121",
	}
	for i := range inAuthType {
		got, err := authDataToString(inAuthType[i], []byte{1, 2, 3})
		if got != want[i] || err != nil {
			errorFail(t, "auth to string, 3 bytes correct", i, want[i], nil, got, err)
		}
	}

	//THIS HAD RESULTS CHECKED MANUALLY AGAINST C# CODE
	inAuthType = []AuthenticationType{
		AuthenticationTypePin,
		AuthenticationTypeTransponder,
	}
	want = []string{
		"12",
		"258",
	}
	for i := range inAuthType {
		got, err := authDataToString(inAuthType[i], []byte{1, 2})
		if got != want[i] || err != nil {
			errorFail(t, "auth to string, 2 bytes correct", i, want[i], nil, got, err)
		}
	}
}

func TestAuthDataToStringErrors(t *testing.T) {
	//todo then just capitalise first letter
	inAuthType := []AuthenticationType{
		//AuthenticationTypePin, Pin has no limit, so should fail on bytearraytointstring
		// but it also breaks it up byte by byte.. So will never fail.
		AuthenticationTypeRFID,
		AuthenticationTypeTransponder,
	}
	in := make([]byte, 9)
	for i := range inAuthType {
		got, err := authDataToString(inAuthType[i], in)
		if err == nil || !strings.Contains(err.Error(), "authData len expected") {
			errorFail(t, "AuthData string error", i, "", errors.New("authData len expected"), got, err)
		}
	}
}

func testAuthDataToString(t *testing.T) {
	//todo then just capitalise first letter
	inAuthType := []AuthenticationType{
		AuthenticationTypePin,
		AuthenticationTypeRFID,
		//AuthenticationTypeTransponder,
	}
	inAuthData := [][]byte{
		{1, 2, 9, 8},
		{1, 2, 3},
		//{},
	}
	want := []string{
		"1298",
		"321",
	}
	wantErr := []error{
		nil,
		nil,
	}
	for i := range inAuthType {
		got, err := authDataToString(inAuthType[i], inAuthData[i])
		if got != want[i] || (err != nil && err.Error() != wantErr[i].Error()) {
			errorFail(t, "not yet implemented", i, want[i], wantErr[i], got, err)
		}
	}
}

func TestByteArrayToIntString(t *testing.T) {
	input := [][]byte{
		{0x00, 0x0c}, //12
		{1, 1},
		{0, 0, 0, 0},
		{},
		make([]byte, 9),
	}
	want := []string{
		"12",
		"257",
		"0",
		"0",
		"0",
	}
	wantErr := []error{
		nil,
		nil,
		nil,
		nil,
		errors.New("array too large, max len 8"),
	}
	for i := range input {
		got, err := ByteArrayToIntString(input[i])
		if got != want[i] || (err != nil && err.Error() != wantErr[i].Error()) {
			errorFail(t, "bytearray", i, want[i], wantErr[i], got, err)
		}
	}
}

func errorFail(t *testing.T, test string, testIndex int, want string, wantErr error, got string, err error) {
	t.Errorf("In %s, test %d failed\nwant:\t%v\nwanterr:\t%v\ngot:\t%v\ngoterr: \t%v", test, testIndex, want, wantErr, got, err)
}
