package main

import (
	"fmt"
	"io/ioutil"
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

	stack := interpreter.Stack{}

	stack.New()

	code, err := ioutil.ReadFile("../../other/example/test.sd")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	stack.Build(string(code), out)

	output, _ := stack.Execute()

	// Print output
	fmt.Println(output)

	fmt.Println(stack)
}
