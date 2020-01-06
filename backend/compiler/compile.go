package compiler

import (
	"bytes"
	"fmt"
	"os"
	"snareDrum/backend/interpreter"
)

// Compile the register of the stack to binary
func Compile(s interpreter.Stack, progress *int) []byte {

	out := []byte{}

	// Iterate over stack's register
	for _, fun := range s.Register.Methods {
		*progress++
		out = append(out, FunctionToBinary(fun)...)
	}

	return out
}

// CompileAndWrite the binary to a file
func CompileAndWrite(s interpreter.Stack, path string, progress *int) error {

	// Open the file
	out := bytes.Buffer{}

	file, err := os.Create("GENERATED.sd")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	// Iterate over stack's register
	for _, fun := range s.Register.Methods {
		*progress++
		out.Write(FunctionToBinary(fun))
	}

	// Write to file
	file.Write(out.Bytes())

	// Flush the buffer
	out.Reset()

	return nil
}

// Execute the code of a given .sdexe file
func Execute(binary []byte) *bytes.Buffer {
	// tokenize the given binary code
	tokenizedCode := tokenize(binary)

	// Create stack (to allow debugging)
	stack := interpreter.Stack{}
	stack.New()

	// Iterate over tokens and build the register of the stack
	for _, token := range tokenizedCode {
		stack.Register.Add(BinaryToFunction(token))
	}

	// Execute functions
	output, _ := stack.Execute()
	return output
}

// divide the binary code into tokens
func tokenize(code []byte) [][]byte {
	out := [][]byte{}

	// Counter of the length of the token
	// purpousedly don't use the for loop's iterator
	// because it's easier
	ind := 0

	// Temporary token
	tempToken := []byte{}

	// Iterate over code
	for _, b := range code {

		tempToken = append(tempToken, b)
		ind++

		if ind == 4 {
			// reset index
			ind = 0

			// add to out
			out = append(out, tempToken)

			// reset temptoken
			tempToken = []byte{}
		}

	}
	return out
}
