package ws_cmd

type Cmd string

const (
	InstantMessage = "instant message"
	OpenRoom       = "open room"
	Ask            = "ask"
	Accept         = "accept"
	Reject         = "reject"
	JoinRoom       = "join room"
	CloseRoom      = "close room" // called by mentor
	LeaveRoom      = "leave room" // called by student
	AccidentLeave  = "accident leave"

	// class status informance
	Online  = "online"
	Offline = "offline"
	InClass = "in class"

	// start class procedure
	AskAcceptance = "ask acceptance"
	AcceptClass   = "accept class"
	ClockOn       = "clock on"
	PayFail       = "pay fail"
	PaySucess     = "pay success"
	StartClass    = "start class"
	FinishClass   = "finish class"
	UnexpectError = "unexpect error"

	// webRTC
	Offer     = "offer"
	Answer    = "answer"
	Candidate = "candidate"

	// heart beat
	Ping = "ping"
	Pong = "pong"
)
