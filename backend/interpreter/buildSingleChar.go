package interpreter

// build the program without formating
// from a given code (snipplet)
// only assemble the register -> don't run
// functions
func buildSingeChardRegister(program string, lang Lang, s *Stack, progress *int) Register {
	out := Register{}

	// Iterate over the program code
	for i := 0; i < len(program); i++ {
		*progress++
		// if isWhiteSpace -> skip
		if !isWhiteSpace(program[i]) {

			// Check lang
			switch program[i] {
			case lang.Pointer.Up[0]:
				IncreaseMemoryPointer(s)

				// Add to register
				out.add(IncreaseMemoryPointer)
				break
			case lang.Pointer.Down[0]:
				DecreaseMemoryPointer(s)

				// Add to register
				out.add(DecreaseMemoryPointer)
				break
			case lang.IO.In[0]:
				break
			case lang.IO.Out[0]:
				// printCell(s)

				// Add to register
				out.add(PrintCell)
				break
			case lang.Cell.Add[0]:
				AddToCell(s)

				// Add to register
				out.add(AddToCell)
				break
			case lang.Cell.Sub[0]:
				SubFromCell(s)

				// Add to register
				out.add(SubFromCell)
				break
			case lang.Loop.Start[0]:

				// Find out the length of the loop
				length := loopLength(program[i:len(program)], lang.Loop.Start[0], lang.Loop.End[0])

				// Dummy progress for not overrunning the actual progress
				dummyProg := 0
				// Iterate untill > 1 -> last run will be
				// executed on the main function
				// -> no need to manipulate i
				for getCell(s).Value > 1 {
					out.merge(buildSingeChardRegister(program[i+1:i+length], lang, s, &dummyProg))
				}

				break
			case lang.Loop.End[0]:
				break
			default:
				break
			}
		}
	}
	return out
}

// Check if given rune is whitespace
func isWhiteSpace(r byte) bool {
	if r == '\t' || r == '\n' || r == ' ' {
		return true
	}
	return false
}

// loopLength returns the length
// of the loop -> from init [ to closing ]
func loopLength(prog string, lStart, lEnd byte) int {
	// Number of all loop openings -> all subloops
	nOpening := 0

	// Iterate over code
	for i, token := range prog {

		// Detect levels
		if byte(token) == lStart {
			nOpening++
		} else if byte(token) == lEnd {
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
