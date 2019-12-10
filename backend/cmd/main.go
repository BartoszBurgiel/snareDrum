package main

import (
	"fmt"
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
}
