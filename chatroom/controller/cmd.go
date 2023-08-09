package controller

type Cmd int

const (
	SingleChat Cmd = iota
	SingleChatReaded
	GroupChat
	GroupChatReaded
	ClassroomInfo
	Heartbeat
)
