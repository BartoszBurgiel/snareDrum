package interpreter

// Loop represents a loop on a program
type Loop struct {
	Conditional *Cell

	// Lang for reference and code trimming
	Lang     Lang
	Register []func()
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

// trimCode and leave only loop code
// @param -> prog string
// represents the code snipplet
// from the beginning of the loop
// untill the end of the program
func (l *Loop) trimCode(prog *string) string {
	// Number of all loop openings -> all subloops
	n_opening := 0

	snipplet := []byte

	// Iterate over code
	for _, token := range prog {

		// Detect levels
		if token == l.Lang.Loop.Start {
			n_opening++
		} else if token == l.Lang.Loop.End {

			// If on highest level 
			if n_opening = 0 {
				
				// End
				return snipplet
			}
			n_opening--
		}

		snipplet = append(snipplet, token)
	}
}

func (l *Loop) Build() {
	l.Conditional = 
}