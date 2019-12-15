package generator

import (
	"snareDrum/backend/interpreter"
)

// Generate a program printing ani giving string
// in any given language
func Generate(lang interpreter.Lang, progOutput string) string {
	out := ""

	// Optimizing variables
	var nextCharSame, manuallyAddToCell, addToCell bool
	var charDiff int

	for i := 0; i < len(progOutput); i++ {

		// claculate walkerLoop value
		iterations, incrPI, leftover := calcWalkerLoop(getASCIIValue(progOutput[i]))

		// Write the walker loop to get to a certain value
		out += walkerLoopAdd(iterations, incrPI, lang)

		out += handleLeftoverAdd(leftover, lang)

		if i < len(progOutput)-1 {
			nextCharSame, manuallyAddToCell, addToCell, charDiff = isSameOrWithinRange(progOutput[i], progOutput[i+1], 20)
		}

		// Print value
		out += reprint(lang)

		// If next char is same -> print again
		if nextCharSame {
			// Print value
			out += reprint(lang)
			// Jump to the next char
			i++
		} else if manuallyAddToCell {

			out += adjustCellValue(addToCell, charDiff, lang)
			out += reprint(lang)

			// Jump to the next char
			i++
		}

		// Move pointer up
		out += lang.Pointer.Up
		out += whiteSpace(lang)

	}

	return out
}
