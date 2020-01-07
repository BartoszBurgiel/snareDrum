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
	for i, fun := range s.Register.Methods {
		*progress++
		out.Write(FunctionToBinary(fun))

		// Add whitespace
		if i > 0 && i%20 == 0 {
			out.WriteByte('\n')
		}
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

	// Temporary token
	tempToken := []byte{}

	// Iterate over code
	for i, b := range code {

		tempToken = append(tempToken, b)

		if i%4 == 0 && i > 0 {
			out = append(out, tempToken)

			// reset temptoken
			tempToken = []byte{}
		}

	}
	return out
}
