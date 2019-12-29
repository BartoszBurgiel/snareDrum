package main

import (
	"fmt"
	"io/ioutil"
	"snareDrum/backend/compiler"
	"snareDrum/backend/interpreter"
	"snareDrum/backend/interpreter/reader"
)

func main() {

	bfHeader, err := ioutil.ReadFile("../../other/example/hgicht/header.json")
	if err != nil {
		fmt.Println(err)
	}

	bfLang := reader.ReadHeader(bfHeader)

	bfStack := interpreter.Stack{}
	bfStack.New()

	bfCode, err := ioutil.ReadFile("../../other/example/hgicht/hgichtGen.sd")
	if err != nil {
		fmt.Println(err)
	}

	bfStack.Build(string(bfCode), bfLang)

	compiled := compiler.Compile(bfStack)

	err = ioutil.WriteFile("compiled.sdexe", compiled, 0644)
	if err != nil {
		fmt.Println(err)
	}

	log := compiler.Execute(compiled)

	fmt.Println(log)

}
