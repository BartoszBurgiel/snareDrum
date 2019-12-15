package main

import (
	"fmt"
	"io/ioutil"
	"snareDrum/backend/generator"
	"snareDrum/backend/interpreter"
	"snareDrum/backend/interpreter/reader"
)

func main() {

	hgichtHeader, err := ioutil.ReadFile("../../other/example/hgicht/header.json")
	if err != nil {
		fmt.Println(err)
	}

	hgicht := reader.ReadHeader(hgichtHeader)

	hgichtStack := interpreter.Stack{}
	hgichtStack.New()

	hgichtGenerated := generator.Generate(hgicht, "Ich habe ein Programm geschrieben, der Programme schreibt :D")

	err = ioutil.WriteFile("../../other/example/hgicht/hgichtGen.sd", []byte(hgichtGenerated), 0644)
	if err != nil {
		fmt.Println(err)
	}

	// test
	hgichtStack.Clear()

	hgichtStack.Build(hgichtGenerated, hgicht)

	outputAleksandra, _ := hgichtStack.Execute()

	debug := hgichtStack.Debug()

	fmt.Println(debug)
	fmt.Println(outputAleksandra)
}
