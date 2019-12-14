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

	//fmt.Println(out)

	stack := interpreter.Stack{}

	stack.New()

	code, err := ioutil.ReadFile("../../other/example/bartosz.sd")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	stack.Build(string(code), out)

	//output, _ := stack.Execute()

	// Print output
	//fmt.Println(output)

	hgichtHeader, err := ioutil.ReadFile("../../other/example/hgicht/header.json")
	if err != nil {
		fmt.Println(err)
	}

	hgicht := reader.ReadHeader(hgichtHeader)

	hgichtStack := interpreter.Stack{}
	hgichtStack.New()

	hgichtCode, err := ioutil.ReadFile("../../other/example/hgicht/hgicht.sd")
	if err != nil {
		fmt.Println(err)
	}

	hgichtStack.Build(string(hgichtCode), hgicht)

	hgichtOut, _ := hgichtStack.Execute()

	fmt.Println(hgichtOut)
}
