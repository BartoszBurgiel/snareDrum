package concurrency

import (
	"bytes"
	"snareDrum/backend/interpreter"
)

// ProcessedMaterial contains the processed material
// and its index to help to put it back together
type ProcessedMaterial struct {
	Index   int
	Content *bytes.Buffer
}

// ProcessedRegister contains the register
// for the stack
type ProcessedRegister struct {
	Index   int
	Content []func(s *interpreter.Stack) int
}
