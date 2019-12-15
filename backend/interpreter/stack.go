package interpreter

import (
	"reflect"
	"runtime"
	"strings"
)

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

// New stack will be initialized
func (s *Stack) New() {

	// Build
	s = &Stack{
		Cells: struct {
			Number int
			Cells  []Cell
		}{
			Number: 0,
			Cells:  []Cell{},
		},
		MemoryPointer: Pointer{
			Pos: 0,
		},
		Register: Register{
			Methods: []func(s *Stack) int{},
		},
	}

	// // Add one empty cell
	// s.Register.add(addEmptyCell)
	// addEmptyCell(s)
}

// Dump all data from the stack execpt the
// register
func (s *Stack) Dump() {
	s.Cells.Cells = nil
	s.Cells.Number = 0
	s.MemoryPointer.Pos = 0
}

// Clear the whole stack -> delete all cells
// and functions
func (s *Stack) Clear() {
	s.Dump()
	s.Register.Methods = nil
}

// Build stack based on the .sd file to
func (s *Stack) Build(program string, lang Lang) {
	// Determine in what way the program should be built
	if lang.SingleChard {

		// Delete whitespace
		programNoWhitespace := removeWhitespace(program)

		s.Register = buildSingeChardRegister(programNoWhitespace, lang, s)
	} else {

		// Divide to keywords
		progSlice := sliceProgram(program)

		s.Register = buildRegister(progSlice, lang, s)
	}

	s.Dump()
}

// Execute and call all functions of the register
// And return the output of the program
// and calls -> names of all called functions
func (s *Stack) Execute() (output string, funcCalls []string) {

	for _, f := range s.Register.Methods {

		// Get the name of the function
		function := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()

		// Format name of the function
		functionFormatted := strings.SplitAfter(function, ".")

		// Print the name of the function
		funcCalls = append(funcCalls, functionFormatted[1])

		// Call function anf get the value
		v := f(s)

		// Add to output
		output += string(byte(v))
	}

	return output, funcCalls
}
