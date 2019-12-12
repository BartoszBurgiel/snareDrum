package interpreter

// Stack struct represents a stack with all actions of a program and their order
type Stack struct {
	// All cells
	Cells struct {
		// Number of all cells
		Number int

		// Acutal cells
		Cells []Cell
	}

	// The current memory pointer
	MemoryPointer Pointer

	// All called functions
	Register Register
}

// Execute and call all functions of the register
func (s Stack) Execute() {
	for _, f := range s.Register.Methods {
		f(&s)
	}
}
