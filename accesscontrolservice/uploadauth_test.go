package accesscontrolservice

import (
	"net"
	"testing"
)

func getOAList() []OldEntity {
	return []OldEntity{
		{AuthenticationTypePin, 1},
		{AuthenticationTypePin, 258}, //test little endied; 00000001 00000010
		{AuthenticationTypePassword, 1},
		{AuthenticationTypePassword, 258},
		{AuthenticationTypeRFID, 1},
		{AuthenticationTypeRFID, 258},
	}
}
func TestOldAuthEncodeTS2(t *testing.T) {
	in := getOAList()
	//header 10
	//10th is length of OA
	buf := []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, byte(len(in))}
	for _, oa := range in {
		buf = append(buf, oa.EncodeTS2()...)
	}
	interfaces, err := net.Interfaces()
	if err != nil {
		t.Error("idk test OldAuthEncode fialed getting netowrk address that isn't REALLY necessary")
	}
	addr, err := interfaces[0].Addrs()
	if err != nil {
		t.Error("idk test OldAuthEncode fialed getting netowrk address that isn't REALLY necessary")
	}
	result, err := ProcessUploadAuthDataRequest(buf, addr[0])
	if err != nil {
		t.Error("idk test OldAuthEncode fialed getting netowrk address that isn't REALLY necessary")
	}
	for i := range result.Entities {
		if in[i].AuthType != result.Entities[i].AuthType || in[i].UserID != result.Entities[i].UserID {
			t.Errorf("OldAuth test %d do not match\nwant:\t%s\n got:\t%s", i, in[i].toString(), result.Entities[i].toString())
		}
	}
}
