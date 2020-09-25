# ts2-api

Golang based api for ts2 (starting with acces control service)

## Process

This is info from the current TS2 Access control service

Receive UDP packet. If a **not new** action/connection (ie previous action from client not complete), send Ack (more later)

### ProcessAck 

Through the old connections object, 'ProcessAckPacket'
This method is valid through a few;

[AuthenticateEntryCommand.cs](https://github.com/TouchSafe3/touchsafe2/blob/master/TouchSafe.AccessControlService/EntryController/AuthenticateEntryCommand.cs#L106)

[DownloadAuthTableCommand.cs](https://github.com/TouchSafe3/touchsafe2/blob/master/TouchSafe.AccessControlService/EntryController/DownloadAuthTableCommand.cs#L119)

[RecieveUserAuthCommand.cs](https://github.com/TouchSafe3/touchsafe2/blob/master/TouchSafe.AccessControlService/EntryController/RecieveUserAuthCommand.cs#L109)

[SendAuthTableCommand.cs](https://github.com/TouchSafe3/touchsafe2/blob/master/TouchSafe.AccessControlService/EntryController/SendAuthTableCommand.cs#L127)

[SendBusAuthTableCommand.cs](https://github.com/TouchSafe3/touchsafe2/blob/master/TouchSafe.AccessControlService/EntryController/SendBusAuthTableCommand.cs#L111)

[SendOldAuthTableCommand.cs](https://github.com/TouchSafe3/touchsafe2/blob/master/TouchSafe.AccessControlService/EntryController/SendOldAuthTableCommand.cs#L106)



In this function it uses a switch case (with only one case and no default...) to check the authentication state is WaitAck for both AethenticateEntryCommand and RecieveUserAuthCommand. (authstate is apart of connection, not passed through udp data)

#### Authstate

These methods are near identical, if the requested command is AuthenticateAck, and the sequence number is correct, just remove the connection now.
Only difference is RecieveUserAuthCommand logs if it was vehicle or user, authentrycommand just logs asset id.

#### TransferState
All these are to do with the auth table and in reference to sending it somewhere.

(DownloadAuthTableCommand, SednAuthTableCommand, SendBusAuthTableCommand, SendOldAuthTableCommand)

The only difference is after sanity checking the header and checking the sequence number;

Reset retry count ( we are sending a new packet), increment the sequence numbers, send next packet of data (which is part of the authentication table)

These should all just be one method...

## Def

UDP

buf[0]= Board command


|BoardCommand|Process|Rec/Send|header len|
|---|---|---|---|
|AuthenticateSingle|RecieveUserAuthCommand|StartConnection|8|
|UploadAuthData|RecieveUploadAuthDataCommand|Recieve|10|
|UploadOfflineAuthData|RecieveUploadOfflineAuthDataCommand|Recieve|NA|
|ForceNewAuthTable/DownloadAuthenticate|SendAuthTableCommand|StartConnection|?|
|DownloadOldAuthenticate|SendOldAuthTableCommand|StartConnection|?|
|BusAuthenticate|SendBusAuthTableCOmmand|StartConnection|?|
|IsServiceOnline|IsServiceOnlineCommand|Recieve|NA|
|ALL OTHERS|???????????|??????????|?|

### AuthenticateSingle

Headerlength=8

|buf[0]|buf[1]|buf[2]|buf[3]|buf[4]|buf[5]|buf[6]|buf[7]|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|Request command|Request header size| Sequence number (uint16)||AP ID (uint16)||Request auth type|req data length|


### UploadAuthData

Headerlength=10

|buf[0]|buf[1]|buf[2]|buf[3]|buf[4]|buf[5]|buf[6]|buf[7]|buf[8]|buf[9]|
|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|
|Request command|location ID (uint16)||Tag # (uint32)||||Sequence # (uint16) **is big endian needs to swap to little** ||entity count|

**for each entity in entity #**

|data[0]|data[1]|data[2]|
|:-----:|:-----:|:-----:|
|auth type|userid (uint16) **big to little endian swap needed**||

### UploadOfflineAuthData

HeaderLength=13
|buf[0]|buf[1]|buf[2]|buf[3]|buf[4]|buf[5]|buf[6]|buf[7]|buf[8]|buf[9]|buf[10]|buf[11]|buf[12]|buf[13]|
|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:-----:|:-----:|:-----:|:-----:|
|Request command|Sequence # (uint16) || location ID (uint16)||Entity count (max is 100) uint8?|timeLong (assuming epoc) int64|||||||||

**for each entity in entity #**

|data[0]|data[1]|data[2]|data[3]|data[4]|data[5]|data[6]|data[7]|
|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|
|auth type|rawuserdata (needs to be byte reversed)|||Seconds relative to time uint32||||

**If transponder**

|data[8]|
|:-----:|
|Passenger count (uint8)|

**If passenger>0 NOTE:repeats for >1 obviously**

|data[9]|data[10]|data[11]|
|:-----:|:-----:|:-----:|
|auth type|user id (reverse bytes then uint16)||


**rawuserdata**

|Auth type|rawuserdata[0]|rawuserdata[1]|rawuserdata[2]|
|:-------:|:---------:|:---------:|:---------:|
|RFID|NOT USED|auth data (uint16)||
|Pin|Use all, expand to 4byte convert to uint32|||
|Transponder|Use all, put into array format [0111] then reverse and convert uint32|||

TODO: Check what pins actually come out as, we are copying 3 bytes into 4 byte array, just verify logically what the outcome is or should be(make C# and go programs and check they match)


`
var rawUserData = binaryReader.ReadBytes(3);
uint authData = 0;
Array.Reverse(rawUserData);
var tempArray = new byte[4];
Array.Copy(rawUserData, 0, tempArray, 0, rawUserData.Length);
`

## word swizzle

`
private static UInt16 WordSwizzle(UInt16 word)
		{
			return (UInt16)((word >> 8) | (word << 8));
		}
`