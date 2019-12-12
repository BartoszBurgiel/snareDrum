package main

import (
	"fmt"
	"snareDrum/backend/interpreter"
	"snareDrum/backend/interpreter/reader"
)

func main() {
	out := reader.ReadHeader([]byte(`{
		"_comment": "Brainfuck lang setup",
		"SingleChard": true,
		"Pointer": {
			"Up": ">",
			"Down": "<"
		},
		"IO": {
			"In": ",",
			"Out": "."
		},
		"Cell": {
			"Add": "+",
			"Sub": "-"
		},
		"Loop": {
			"Start": "[",
			"End": "]"
		}
	}`))

	fmt.Println(out)

	stack := interpreter.Stack{
		Cells: struct {
			Number int
			Cells  []interpreter.Cell
		}{
			Number: 0,
			Cells:  []interpreter.Cell{},
		},
		MemoryPointer: interpreter.Pointer{
			Pos: 0,
		},
		Register: interpreter.Register{
			Methods: []func(s *interpreter.Stack){},
		},
	}

	stack.Build(`+++++ +++++ +++++ +++++
	+++++ +++++ +++++ +++++
	+++++ +++++ +++++ +++++
	+++++ .`, out)

	fmt.Println(stack)

	stack.Execute()

	fmt.Println(stack)
}
