package generator

import (
	"snareDrum/backend/interpreter"
)

// Detect if a character is equal to the following
// or at least is in the certain range
func isSameOrWithinRange(currChar, nextChar byte, cRange int) (nextCharSame, manuallyAddToCell, addToCell bool, charDiff int) {
	if currChar == nextChar {
		return true, false, false, 0
	}

	// Calculate the difference between current and next char
	charDiff = getASCIIValue(currChar) - getASCIIValue(nextChar)

	// Check if absolute difference < 15
	if abs(charDiff) < cRange {

		// Activate manual adjustment of the cell
		manuallyAddToCell = true

		// Calculate the "direction" of the adjustment
		if charDiff < 0 {
			addToCell = true
		}
	}
	return false, manuallyAddToCell, addToCell, charDiff
}

// Print the value of current cell one again
func reprint(lang interpreter.Lang) string {
	out := lang.IO.Out
	out += whiteSpace(lang)
	return out
}

// Add or substract cell's value by given value
func adjustCellValue(addToCell bool, charDiff int, lang interpreter.Lang) string {
	out := ""

	if addToCell {
		// Add manually
		for j := 0; j < -charDiff; j++ {
			out += lang.Cell.Add
			out += whiteSpace(lang)
		}

	} else {

		// Add manually
		for j := 0; j < charDiff; j++ {
			out += lang.Cell.Sub
			out += whiteSpace(lang)
		}

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

// handle leftover from a loop -> add
// leftover 'add' tokens
func handleLeftoverAdd(leftover int, lang interpreter.Lang) string {
	out := ""

	// if no leftover
	if leftover == -1 {
		return ""
	}

	// Manually add values to the cell untill
	// leftover excided
	for j := 0; j < leftover; j++ {
		out += lang.Cell.Add
		out += whiteSpace(lang)
	}

	return out
}

// handle leftover from a loop -> sub
// leftover 'sub' tokens
func handleLeftoverSub(leftover int, lang interpreter.Lang) string {
	out := ""

	// if no leftover
	if leftover == -1 {
		return ""
	}

	// Manually add values to the cell untill
	// leftover excided
	for j := 0; j < leftover; j++ {
		out += lang.Cell.Sub
		out += whiteSpace(lang)
	}

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
