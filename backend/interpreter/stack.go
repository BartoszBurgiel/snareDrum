package interpreter

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"snareDrum/backend/ui"
	"strings"
	"time"
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
func (s *Stack) Build(lang Lang, program string) {

	fmt.Println("Building...")
	progSlice := DivideProgram(lang, program)

	s.Register = BuildRegister(progSlice, lang, s)

	s.Dump()
}

// Execute and call all functions of the register
// And return the output of the program
// and calls -> names of all called functions
func (s *Stack) Execute() (output *bytes.Buffer, funcCalls []string) {
	output = &bytes.Buffer{}
	max := len(s.Register.Methods)

	//opTime := []int{}

	operationTimestamp := time.Now()
	fmt.Println("Executing...")
	for i, f := range s.Register.Methods {

		// Print the name of the function
		funcCalls = append(funcCalls, FormatFunctionName(f))

		// Call function anf get the value
		v := f(s)

		// Add to output
		output.WriteByte(byte(v))
		//opTime = append(opTime, int(time.Second/time.Now().Sub(operationTimestamp)))
		ui.PrintProgressBar(i, max-1, operationTimestamp)
		operationTimestamp = time.Now()

	}

	// min, max, avr := 999999999, 0, 0.0
	// for _, as := range opTime {
	// 	if as < min {
	// 		min = as
	// 	}
	// 	if as > max {
	// 		max = as
	// 	}
	// 	avr += float64(as) * float64(1.0/float64(len(opTime)))

	// }
	// fmt.Printf("\n\nmin:%d max:%d avr:%f", min, max, avr)

	return output, funcCalls
}

// FormatFunctionName and return it as string
func FormatFunctionName(f func(s *Stack) int) string {

	// Get and Format name of the function
	functionFormatted := strings.SplitAfter(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
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
