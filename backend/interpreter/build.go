package interpreter

import (
	"fmt"
)

// Build stack based on the .sd file to
func (s *Stack) Build(program string, lang Lang) {
	// Determine in what way the program should be built
	if lang.SingleChard {

		// Delete whitespace
		programNoWhitespace := removeWhitespace(program)

		s.Register = buildSingeChardRegister(programNoWhitespace, lang, s)
	}

	s.Dump()
}

// build the program without formating
// from a given code (snipplet)
// only assemble the register -> don't run
// functions
func buildSingeChardRegister(program string, lang Lang, s *Stack) Register {

	out := Register{}

	// Iterate over the program code
	for i := 0; i < len(program); i++ {

		// if isWhiteSpace -> skip
		if !isWhiteSpace(program[i]) {

			fmt.Println(program[i], "->", string(program[i]), ":", i, "lenp:", len(program))

			// Check lang
			switch program[i] {
			case lang.Pointer.Up[0]:
				increaseMemoryPointer(s)

				// Add to register
				out.add(increaseMemoryPointer)
				break
			case lang.Pointer.Down[0]:
				decreaseMemoryPointer(s)

				// Add to register
				out.add(decreaseMemoryPointer)
				break
			case lang.IO.In[0]:
				break
			case lang.IO.Out[0]:
				// printCell(s)

				// Add to register
				out.add(printCell)
				break
			case lang.Cell.Add[0]:
				addToCell(s)

				// Add to register
				out.add(addToCell)
				break
			case lang.Cell.Sub[0]:
				subFromCell(s)

				// Add to register
				out.add(subFromCell)
				break
			case lang.Loop.Start[0]:

				// Find out the length of the loop
				length := loopLength(program[i:len(program)], lang.Loop.Start[0], lang.Loop.End[0])

				// Iterate untill > 1 -> last run will be
				// executed on the main function
				// -> no need to manipulate i
				for getCell(s).Value > 1 {
					out.merge(buildSingeChardRegister(program[i+1:i+length], lang, s))
				}

				break
			case lang.Loop.End[0]:
				break
			default:
				fmt.Println("Unknown character", string(program[i]))
				return Register{}
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

// delete all whitespace
// from program
func removeWhitespace(program string) string {
	out := ""

	for _, c := range program {
		// if not whitespace -> add to out
		if !isWhiteSpace(byte(c)) {
			out += string(c)
		}
	}

	return out
}
