package interpreter

// Register struct represents the slice
// of functions that will be executed
// during the runtime of the program
type Register struct {
	Methods []func(s *Stack)
}

// Add a function to the register
func (r *Register) add(f func(s *Stack)) {
	r.Methods = append(r.Methods, f)
}

// merge two register structs
func (r *Register) merge(newR Register) {
	// iterate over newR
	for _, f := range newR.Methods {
		r.add(f)
	}
}
