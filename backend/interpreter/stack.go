package interpreter

import (
	"bytes"
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

// Build stack based on the .sd file
func (s *Stack) Build(lang Lang, program string, progress *int) {

	progSlice := DivideProgram(lang, program)

	s.Register = BuildRegister(progSlice, lang, s, progress)

	s.Dump()
}

// Execute and call all functions of the register
// And return the output of the program
// and calls -> names of all called functions
func (s *Stack) Execute() (output *bytes.Buffer, funcCalls []string) {

	for _, f := range s.Register.Methods {

		// Print the name of the function
		funcCalls = append(funcCalls, FormatFunctionName(f))

		// Call function anf get the value
		v := f(s)

		// Add to output
		output.WriteByte(byte(v))
	}

	return output, funcCalls
}

// FormatFunctionName and return it as string
func FormatFunctionName(f func(s *Stack) int) string {
	// Get the name of the function
	function := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()

	// Format name of the function
	functionFormatted := strings.SplitAfter(function, ".")
	return functionFormatted[1]
}

// DivideProgram to keywords or single chars
func DivideProgram(lang Lang, program string) (progSlice []string) {
	// Determine in what way the program should be built
	if lang.SingleChard {
		progSlice = stringToStringSlice(program)
	} else {
		// Divide to keywords
		progSlice = sliceProgram(program)
	}
	return progSlice
}
