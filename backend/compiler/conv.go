package compiler

import "snareDrum/backend/interpreter"

// FunctionToBinary converts the given function (name)
// into a binary code
func FunctionToBinary(f func(s *interpreter.Stack) int) []byte {

	switch interpreter.FormatFunctionName(f) {
	case "AddToCell":
		return []byte("0000")
	case "IncreaseMemoryPointer":
		return []byte("0001")
	case "DecreaseMemoryPointer":
		return []byte("0010")
	case "SubFromCell":
		return []byte("0011")
	case "PrintCell":
		return []byte("0100")
	case "ReadToCell":
		return []byte("0101")
	default:
		return []byte("1111")
	}
}

// BinaryToFunction returns the stack function based on the given
// binary code
func BinaryToFunction(code []byte) func(s *interpreter.Stack) int {
	// Turn to string in order for switch to work
	switch string(code) {
	case "0000":
		return interpreter.AddToCell
	case "0001":
		return interpreter.IncreaseMemoryPointer
	case "0010":
		return interpreter.DecreaseMemoryPointer
	case "0011":
		return interpreter.SubFromCell
	case "0100":
		return interpreter.PrintCell
	case "0101":
		return interpreter.ReadToCell
	default:
		return (func(s *interpreter.Stack) int { return 0 })
	}
}
