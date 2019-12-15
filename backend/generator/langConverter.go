package generator

import (
	"snareDrum/backend/interpreter"
)

// Generate a program printing ani giving string
// in any given language
func Generate(lang interpreter.Lang, progOutput string) string {
	out := ""

	for i, c := range progOutput {

		// claculate walkerLoop value
		iterations, incrPI := calcWalkerLoop(getASCIIValue(byte(c)))
		out += walkerLoop(iterations, incrPI, lang)

		// Print value
		out += lang.IO.Out
		out += whiteSpace(lang)

		// Move pointer up
		out += lang.Pointer.Up
		out += whiteSpace(lang)
	}

	return out
}

// Walker loop represents a loop that's purpose is
// to bring the value of a cell to a certain value
func walkerLoop(iterations, increasementPI int, lang interpreter.Lang) string {
	out := ""

	// Add iterations
	for i := 0; i < iterations; i++ {
		out += lang.Cell.Add
		out += whiteSpace(lang)
	}

	// Initialize loop
	out += lang.Loop.Start
	out += whiteSpace(lang)

	// Fill next cell
	out += lang.Pointer.Up
	out += whiteSpace(lang)
	for i := 0; i < increasementPI; i++ {
		out += lang.Cell.Add
		out += whiteSpace(lang)
	}

	// Decr condition
	out += lang.Pointer.Down
	out += whiteSpace(lang)
	out += lang.Cell.Sub
	out += whiteSpace(lang)

	// Close loop
	out += lang.Loop.End
	out += whiteSpace(lang)

	// Move to the cell that has been filled
	out += lang.Pointer.Up
	out += whiteSpace(lang)

	return out
}

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

// Calculate the most optimal walker loop
// to get to the certail number
func calcWalkerLoop(n int) (iterations, incrPI int) {

	// check all dividers below 10
	for i := 3; i <= 20; i++ {
		if n%i == 0 {
			return i, n / i
		}
	}

	// If not able to split efficiently -> 1->n
	return 1, n
}
