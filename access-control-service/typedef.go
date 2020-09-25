package access_control_service

//BoardCommand is the enumeration of the options the board command can send
// though only a handful are actually used in the access control service (any others are just discarded...)
type BoardCommand int

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

//ActionState is the current state of the action, is there no state
// are we waiting for acknowledgement, or are we finishing the connection/process
type ActionState int

const (
	ActionStateNone ActionState = iota
	ActionStateWaitAck
	ActionStateTerminateConnect
)

//UpdateState is what state the current connection is in (generally when looking to ack)
type UpdateState int

const (
	UpdateStateNone UpdateState = iota
	UpdateStateCompleted
	UpdateStateProcessing
)
