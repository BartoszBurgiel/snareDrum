package generator

import (
	"fmt"
	"snareDrum/backend/interpreter"
)

// Generate a program printing ani giving string
// in any given language
func Generate(lang interpreter.Lang, progOutput string) string {
	out := ""

	for i := 0; i < len(progOutput); i++ {
		// Optimazation variables
		nextCharSame := false

		// 21 so it's not detected by default by the algorithm
		charDiff := 21

		// Add value to cell without the loop
		manuallyAddToCell := false

		// Adjust value to cell by... -> true -> add to cell; false -> sub from cell
		addToCell := false

		// claculate walkerLoop value
		iterations, incrPI, leftover := calcWalkerLoop(getASCIIValue(progOutput[i]))

		// Write the walker loop to get to a certain value
		out += walkerLoopAdd(iterations, incrPI, lang)

		// If there's a leftover
		if leftover > 0 {

			// Manually add values to the cell untill
			// leftover excided
			for j := 0; j < leftover; j++ {
				out += lang.Cell.Add
				out += whiteSpace(lang)
			}
		}

		// if next char is same -> print again and incr i

		if i < len(progOutput)-1 {
			if progOutput[i] == progOutput[i+1] {
				nextCharSame = true
			} else {

				// Calculate the difference between current and next char
				charDiff = getASCIIValue(progOutput[i]) - getASCIIValue(progOutput[i+1])

				fmt.Println(charDiff, "->", progOutput[i], ":", progOutput[i+1])

				// Check if absolute difference < 15
				if abs(charDiff) < 15 {

					// Activate manual adjustment of the cell
					manuallyAddToCell = true

					// Calculate the "direction" of the adjustment
					if charDiff < 0 {
						addToCell = true
					}
				}
			}
		}

		// Print value
		out += lang.IO.Out
		out += whiteSpace(lang)

		// If next char is same -> print again
		if nextCharSame {
			// Print value
			out += lang.IO.Out
			out += whiteSpace(lang)

			// Jump to the next char
			i++
		} else if manuallyAddToCell {
			if addToCell {
				for j := 0; j < -charDiff; j++ {
					out += lang.Cell.Add
					out += whiteSpace(lang)
				}
			} else {
				for j := 0; j < charDiff; j++ {
					out += lang.Cell.Sub
					out += whiteSpace(lang)
				}
			}

			// Print value
			out += lang.IO.Out
			out += whiteSpace(lang)

			// Jump to the next char
			i++
		}

		// Move pointer up
		out += lang.Pointer.Up
		out += whiteSpace(lang)

	}

	return out
}

// Walker loop represents a loop that's purpose is
// to bring the value of a cell to a certain value
// by adding
func walkerLoopAdd(iterations, increasementPI int, lang interpreter.Lang) string {
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

// Walker loop represents a loop that's purpose is
// to bring the value of a cell to a certain value
// by sub
func walkerLoopSub(iterations, increasementPI int, lang interpreter.Lang) string {
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
		out += lang.Cell.Sub
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

// return new line character
func newLine() string {
	return "\n"
}

// Calculate the most optimal walker loop
// to get to the certail number
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
