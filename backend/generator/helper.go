package generator

import "snareDrum/backend/interpreter"

// Return the ascii value of any given byte
func getASCIIValue(b byte) int {
	return int(b)
}

// Add whitespace based on lang.SingleChard
func whiteSpace(lang interpreter.Lang) string {
	if lang.SingleChard {
		return ""
	}
	return " "
}

// return new line character
func newLine() string {
	return "\n"
}

// Calculate the most optimal walker loop
// to get to the certain number
func calcWalkerLoop(n int) (iterations, incrPI, leftover int) {

	// check all dividers below 10
	for i := 20; i > 2; i-- {
		if n%i == 0 {
			return i, n / i, -1
		}
	}

	// If not able to split efficiently
	return 10, int(n / 10), n % 10
}

// return absolute value
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
