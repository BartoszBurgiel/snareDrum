package interpreter

// Action interface holds every
// every state-changing method
type Action interface {
	Incr()
	Decr()
}
