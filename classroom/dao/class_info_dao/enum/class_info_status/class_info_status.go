package class_info_status

type Status int

const (
	Init Status = iota
	Accept
	Paid
	Start
	Finish
	Fail
)

func (c Status) String() string {
	return [...]string{"Init", "Accept", "Paid", "Start", "Finish", "Fail"}[c]
}
