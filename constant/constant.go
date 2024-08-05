package constant

type State int

const (
	StateUnknown State = iota
	StateUnplanned
	StatePlanned
	StateDoing
	StateCheck
	StateDone
)
