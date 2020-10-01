package access_control_service

import (
	"testing"
)

//ByteArrayToIntString returns the string representation of the int
// stored in the byte array
// (00000001) = "1"
// (00010000) = "16"
// (00010000) = "16"
// (1111000000001111) = "4027514880"

func TestByteArrayToIntString(t *testing.T) {
	input := [][]byte{
		{0, 12},
		{1, 1},
		{0, 0, 0, 0},
	}
	want := []string{
		"12",
		"257",
		"0",
	}
	wanterr := []error{
		nil,
		nil,
		nil, //need more
	}
	for i := range input {
		got, err := ByteArrayToIntString(input[i])
		if got != want[i] || err != wanterr[i] {
			t.Errorf("\nwant:%v\nwanterr:%v\ngot:%v\ngoterr:%v", want[i], wanterr[i], got, err)
		}

	}
}
