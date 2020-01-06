package generator

import (
	"bytes"
	"snareDrum/backend/interpreter"
)

// Generate a program printing ani giving string
// in any given language
func Generate(lang interpreter.Lang, progOutput string, progress *int) *bytes.Buffer {
	out := bytes.Buffer{}

	// Optimizing variables
	var nextCharSame, manuallyAddToCell, addToCell bool
	var charDiff int

	for i := 0; i < len(progOutput); i++ {
		*progress = i

		// claculate walkerLoop value
		iterations, incrPI, leftover := calcWalkerLoop(getASCIIValue(progOutput[i]))

		// Write the walker loop to get to a certain value
		out.WriteString(walkerLoopAdd(iterations, incrPI, lang))
		out.WriteString(handleLeftoverAdd(leftover, lang))

		// If next char valid
		if i < len(progOutput)-1 {

			// Calculate optimation
			nextCharSame, manuallyAddToCell, addToCell, charDiff = isSameOrWithinRange(progOutput[i], progOutput[i+1], 20)
		}

		// Print value
		out.WriteString(reprint(lang))

		// If next char is same -> print again
		if nextCharSame {
			// Print value
			out.WriteString(reprint(lang))
			// Jump to the next char
			i++

		} else if manuallyAddToCell {
			out.WriteString(adjustCellValue(addToCell, charDiff, lang))
			out.WriteString(reprint(lang))
			i++
		}

		// Move pointer up
		out.WriteString(lang.Pointer.Up)
		out.WriteString(whiteSpace(lang))

	}

	return &out
}
