package interpreter

// build the program without formating
// from a given code (snipplet)
// only assemble the register -> don't run
// functions
func buildRegister(program []string, lang Lang, s *Stack) Register {

	out := Register{}

	// Iterate over the program code
	for i := 0; i < len(program); i++ {

		// Check lang
		switch program[i] {
		case lang.Pointer.Up:
			IncreaseMemoryPointer(s)

			// Add to register
			out.add(IncreaseMemoryPointer)
			break
		case lang.Pointer.Down:
			DecreaseMemoryPointer(s)

			// Add to register
			out.add(DecreaseMemoryPointer)
			break
		case lang.IO.In:
			break
		case lang.IO.Out:
			// printCell(s)

			// Add to register
			out.add(PrintCell)
			break
		case lang.Cell.Add:
			AddToCell(s)

			// Add to register
			out.add(AddToCell)
			break
		case lang.Cell.Sub:
			SubFromCell(s)

			// Add to register
			out.add(SubFromCell)
			break
		case lang.Loop.Start:

			// Find out the length of the loop
			length := loopLengthKey(program[i:len(program)], lang.Loop.Start, lang.Loop.End)

			// Iterate untill > 1 -> last run will be
			// executed on the main function
			// -> no need to manipulate i
			for getCell(s).Value > 1 {
				out.merge(buildRegister(program[i+1:i+length], lang, s))
			}

			break
		case lang.Loop.End:
			break
		default:
			break
		}
	}

	return out
}

// loopLength returns the length
// of the loop -> from init [ to closing ]
func loopLengthKey(prog []string, lStart, lEnd string) int {
	// Number of all loop openings -> all subloops
	nOpening := 0

	// Iterate over code
	for i, token := range prog {

		// Detect levels
		if token == lStart {
			nOpening++
		} else if token == lEnd {
			nOpening--

			// If on highest level
			if nOpening == 0 {
				// End
				return i
			}
		}
	}

	return 0
}

// Divide program to smaller strings
// containing keywords -> tokenize
func sliceProgram(prog string) []string {
	out := []string{}

	currToken := ""
	// Iterate over program
	for _, c := range prog {

		// Check if whitespace
		if isWhiteSpace(byte(c)) {

			// if currToken isnot empty
			if currToken != "" {
				out = append(out, currToken)
				currToken = ""
			}
		} else {
			// Add to currToken
			currToken += string(byte(c))
		}
	}
	return out
}
