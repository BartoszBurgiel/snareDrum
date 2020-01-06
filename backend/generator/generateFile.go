package generator

import (
	"bytes"
	"fmt"
	"os"
	"snareDrum/backend/interpreter"
)

// write single data to file
func writeLineToFile(data []byte, file *os.File) error {

	// Write data to file
	_, err := file.Write(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Save
	err = file.Sync()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// GenerateFile containing the source code in given lang
// of the program that prints fiven progOutput
// returns an empty string in order to be implemented in
// GenerateWithProgressBar
func GenerateFile(lang interpreter.Lang, progOutput []byte, progress *int) string {
	out := bytes.Buffer{}

	file, err := os.Create("GENERATED.sd")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

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
		out.WriteString(reprint(lang))

		// Write out and flush
		err := writeLineToFile(out.Bytes(), file)
		if err != nil {
			fmt.Println(err)
		}
		out.Reset()
	}
	return ""
}
