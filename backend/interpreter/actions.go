package interpreter

// Loop represents a loop on a program
type Loop struct {
	Conditional Cell
	Actions     []Action
}

// Done returns if loop is done looping -> if conditional's value == 0
func (l *Loop) Done() bool {
	if l.Conditional.Value > 0 {
		return false
	}
	return true
}

// Incr loops conditional by one
func (l *Loop) Incr() {
	l.Conditional.Incr()
}

// Decr loops conditional by one
func (l *Loop) Decr() {
	l.Conditional.Decr()
}

// Action interface holds every
// every state-changing method
type Action interface {
	Incr()
	Decr()
}
