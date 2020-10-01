# ts2-api

Golang based api for ts2 (starting with acces control service)

there is a help section for the API at the bottom

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


**Authentication type**
None = 0,
Pin = 1,
Rfid = 2,
Transponder = 3,
Erased = 4,
Fingerprint = 5,
Password = 6,
ManualOverride = 7       
 
### AuthenticateSingle

See [AuthDataToString](https://github.com/TouchSafe3/touchsafe2/blob/master/TouchSafe.AccessControlService/EntryController/RecieveUserAuthCommand.cs#L274-L310)

**Header**

Headerlength=8

|buf[0]|buf[1]|buf[2]|buf[3]|buf[4]|buf[5]|buf[6]|buf[7]|buf[8:]|
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|Request command|Request header size| Sequence number (uint16)||AP ID (uint16)||Request auth type|req data length|AUTHENTICATION DATA|

**Authentication data**

PIN

|auth[0]|auth[1]|auth[2]|auth[3]|
|:-----:|:-----:|:-----:|:-----:|
|Pin#1|Pin#2|Pin#3|Pin#4|

Note there seems to be no limit on length of pin at the moment...

TRANSPONDER

Reverse bytes then cast to int ie; 
123 (00000001 00000010 00000011)
would turn into
197121 (00000011 00000010 00000001)

RFID

Reverse pins then just string it

#### Response

NOTIFY AUTH STATUS
Packet size is 9

|buf[0]|buf[1]|buf[2]|buf[3]|buf[4]|buf[5]|buf[6]|buf[7]|buf[8]|
|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|:----:|
|BoardCommand.AuthResponse|packetsize(9)|Sequencenumber[0]|SequenceNumber[1]|authType(rfid/pin etc)|authLength??? check me|userId[0]|userId[1]|Success (bool)|

Keep retrying if not recieved

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
var rawUserData = binaryReader. ReadBytes(3); 
uint authData = 0; 
Array. Reverse(rawUserData); 
var tempArray = new byte[4]; 
Array. Copy(rawUserData, 0, tempArray, 0, rawUserData. Length); 
`

### ForceNewAuthTable/DownloadAuthenticate

|buf[0]|buf[1]|buf[2]|
|:----:|:----:|:----:|
|Request command|sequence number (uint16)||

Sequence number is where we are starting from
Check for new database
If request is forcenewauthtable or there just is a new one for the 'endpoint (request IP)' build
Otherwise init transfer

#### Responses

#### no new auth table

If there is no new auth table, tell them; 

|buf[0]|buf[1]|buf[2]|
|:----:|:----:|:----:|
|NoNewAuthTable board command|SeqNumber(uint16)[1]|SeqNumber(uint16)[0]|

#### new auth table

##### unack'd (first packet, or resending same packet)

uint16 Packetsize=min(authtable.length - authtableseek, max packet size(500))

Note that the packet size limit is 500, but the header for this data is size 6, so max size is actually 506??

|buf[0]|buf[1]|buf[2]|buf[3]|buf[4]|buf[5]|
|:----:|:----:|:----:|:----:|:----:|:----:|
|Request command (DownloadAuthenticateData)|SeqNumber(uint16)[1]|SeqNumber(uint16)[0]|packetsize(uint16)[1]|packetsize(uint16)[0]|islastpacket?|

islastpacket is 1 (0000 0001) if datapacket size + auth table seek (which should be 0 if calling this function)== the auth table size, 0 (0000 0000) otherwise

data packet is just stored _authtable

##### processackpacket (recieved an acknowledgement)

|buf[0]|buf[1]|buf[2]|
|:----:|:----:|:----:|
|Request command (DownloadAuthenticateAck)|ReqSeqNumber(uint16)|

ReqSeqNumber should match seq number; then iterate sequence number and continue wiht unack'd packes before

### DownloadOldAuthenticate

|buf[0]|buf[1]|buf[2]|
|:----:|:----:|:----:|
|Request command |sequence number (uint16)||

build old auth table? idk just implement

#### Response

SEE AUTH TABLE, DETAILS ARE EXACT SAME AFTER BUILDING OLD AUTH TABLE?

### BusAuthenticate

|buf[0]|buf[1]|buf[2]|
|:----:|:----:|:----:|
|Request command |sequence number (uint16)||

build bus auth table? idk just implement

#### Response

SEE AUTH TABLE, DETAILS ARE EXACT SAME AFTER BUILDING bus AUTH TABLE?

### IsServiceOnline

|buf[0]|buf[1]|buf[2]|
|:----:|:----:|:----:|
|Request command(IsServiceOnlineAck) |SeqNumber(uint16)[1]|SeqNumber(uint16)[0]|

## word swizzle

``` 
private static UInt16 WordSwizzle(UInt16 word)
		{
			return (UInt16)((word >> 8) | (word << 8));
		}
```

## HELP

Just some QoL notes

``` 
dataPacket[1] = (Byte)(_seqNumber & 0xff);
dataPacket[2] = (Byte)(_seqNumber >> 8);
```

All this does is store _seqNumber as a uint16 with swapped bytes, when recieving you would just cast combine and cast Uint16(dataPacker[1:2]) (not correct syntax but you get it)
Would make more sense to just convert _seqNumber to byte array then go; 

``` 
dataPacket[1]=_seqNumber[1]
dataPacket[2]=_seqNumber[0]
```

HELP FOR UNDERSTANDING WHY THEY ARE SHIFTING BITS SO MUCH!?!?!

``` 
using System;

namespace OperatorsAppl {

   class Program {
   
      static void Main(string[] args) {
         UInt16 a = 60;            /* 60 = 0011 1100 */ 
          UInt16 b = 13;            /* 13 = 0000 1101 */
          UInt16 c = 0; 
         Byte d;
         
         c = (UInt16)(a & b);             /* 12 = 0000 1100 */ 
         Console.WriteLine("Line 1 - Value of c is {0}", c );
         
         c = (UInt16)(a | b);             /* 61 = 0011 1101 */
         Console.WriteLine("Line 2 - Value of c is {0}", c);
         
         c = (UInt16)(a ^ b);             /* 49 = 0011 0001 */
         Console.WriteLine("Line 3 - Value of c is {0}", c);
         
         c = (UInt16)(~a);                /*-61 = 1100 0011 */
         Console.WriteLine("Line 4 - Value of c is {0}", c);
         
         c = (UInt16)(a << 2);      /* 240 = 1111 0000 */
         Console.WriteLine("Line 5 - Value of c is {0}", c);
         
         c = (UInt16)(a >> 4);      /* 15 = 0000 1111 */
         Console.WriteLine("Line 6 - Value of c is {0}", c);
        c = (UInt16)(a >> 6);      /* 15 = 0000 1111 */
         Console.WriteLine("Line 7 - Value of c is {0}", c);
        c = (UInt16)(a >> 8);      /* 15 = 0000 1111 */
         Console.WriteLine("Line 8 - Value of c is {0}", c);
        c = (UInt16)(UInt16.MaxValue >> 8);      /* 15 = 0000 1111 */
         Console.WriteLine("Line 9 - Value of c is {0}", c);
         
         
        d = (Byte)(UInt16.MaxValue >> 8);      /* 15 = 0000 1111 */
         Console.WriteLine("Line 10 - Value of d is {0}", d);
         d = (Byte)(UInt16.MaxValue & 0x00ff);      /* 15 = 0000 1111 */
         Console.WriteLine("Line 11 - Value of d is {0}", d);
         
         //max uint16= 65535
         //max uint8 = 
         Console.ReadLine();
      }
   }
}
```
