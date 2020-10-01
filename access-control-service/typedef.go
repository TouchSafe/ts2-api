package access_control_service

//BoardCommand is the enumeration of the options the board command can send
// though only a handful are actually used in the access control service (any others are just discarded...)
type BoardCommand uint

const (
	BoardCommandNone                        BoardCommand = iota // None
	BoardCommandUploadAuthData                                  // Entity data
	BoardCommandUploadAuthDataAck                               // Acknowledgement
	BoardCommandAuthenticateSingle                              // Authenticate the specified entity
	BoardCommandDownloadOldAuthenticate                         // Download the authentication table
	BoardCommandDownloadOldAuthenticateAck                      // Acknowledge data
	BoardCommandDownloadOldAuthenticateData                     // Data for download authenticate
	BoardCommandAuthenticateResponse
	BoardCommandAuthenticateAck
	BoardCommandUploadOfflineAuthData
	BoardCommandUploadOfflineAuthDataAck
	BoardCommandDownloadAuthenticate     // Download the authentication table
	BoardCommandDownloadAuthenticateAck  // Acknowledge data
	BoardCommandDownloadAuthenticateData // Data for download authenticate
	BoardCommandIsServiceOnline
	BoardCommandBoardCommandIsServiceOnlineAck
	BoardCommandNoNewAuthTable
	BoardCommandForceNewAuthTable
	BoardCommandBusAuthenticateAck
	BoardCommandBusAuthenticate
	BoardCommandBusAuthenticateData
)

//String converts the board command to the string of what it is
func (b BoardCommand) String() string {
	return [...]string{
		"None",
		"UploadAuthData",
		"UploadAuthDataAck",
		"AuthenticateSingle",
		"DownloadOldAuthenticate",
		"DownloadOldAuthenticateAck",
		"DownloadOldAuthenticateData",
		"AuthenticateResponse",
		"AuthenticateAck",
		"UploadOfflineAuthData",
		"UploadOfflineAuthDataAck",
		"DownloadAuthenticate",
		"DownloadAuthenticateAck",
		"DownloadAuthenticateData",
		"IsServiceOnline",
		"BoardCommandIsServiceOnlineAck",
		"NoNewAuthTable",
		"ForceNewAuthTable",
		"BusAuthenticateAck",
		"BusAuthenticate",
		"BusAuthenticateData",
	}[b]
}

//ActionState is the current state of the action, is there no state
// are we waiting for acknowledgement, or are we finishing the connection/process
type ActionState uint

const (
	ActionStateNone ActionState = iota
	ActionStateWaitAck
	ActionStateTerminateConnect
)

//String converts the board command to the string of what it is
func (a ActionState) String() string { return [...]string{"None", "WaitAck", "TerminateConnect"}[a] }

//UpdateState is what state the current connection is in (generally when looking to ack)
type UpdateState uint

const (
	UpdateStateNone UpdateState = iota
	UpdateStateCompleted
	UpdateStateProcessing
)

//String converts the board command to the string of what it is
func (u UpdateState) String() string { return [...]string{"None", "Completed", "Processing"}[u] }

//AuthenticationType specifies which type of authentication is being used while authenticating
type AuthenticationType uint

const (
	AuthenticationTypeNone AuthenticationType = iota
	AuthenticationTypePin
	AuthenticationTypeRFID
	AuthenticationTypeTransponder
	AuthenticationTypeErased
	AuthenticationTypeFingerprint
	AuthenticationTypePassword
	AuthenticationTypeManualOverride
)

//String converts the board command to the string of what it is
func (a AuthenticationType) String() string {
	return [...]string{"None", "Pin", "RFID", "Transponder", "Erased", "Fingerprint", "Password", "ManualOverride"}[a]
}
