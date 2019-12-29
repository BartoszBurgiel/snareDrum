package compiler

import (
	"snareDrum/backend/interpreter"
)

// Compile the register of the stack to binary
func Compile(s interpreter.Stack) []byte {

	out := []byte{}

	// Iterate over stack's register
	for _, fun := range s.Register.Methods {
		out = append(out, FunctionToBinary(fun)...)
	}

	return out
}

// Execute the code of a given .sdexe file
func Execute(binary []byte) string {
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
		if ind == 3 {
			// reset index
			ind = 0

			// add to out
			out = append(out, tempToken)

			// reset temptoken
			tempToken = []byte{}

		} else {
			tempToken = append(tempToken, b)
			ind++
		}

	}
	return out
}
