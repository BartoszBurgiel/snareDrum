package interpreter

import (
	"fmt"
	"reflect"
	"runtime"
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

// Execute and call all functions of the register
func (s *Stack) Execute() {
	output := []byte{}

	for _, f := range s.Register.Methods {
		// Print the name of the function
		fmt.Println(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())

		// Call function anf get the value
		v := f(s)

		// Add to output
		output = append(output, byte(v))
	}

	// Print output
	fmt.Println(string(output))

}
