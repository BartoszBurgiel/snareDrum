package interpreter

import (
	"fmt"
	"unicode"
)

// Build stack based on the .sd file to
func (s *Stack) Build(program string, lang Lang) {
	// Determine in what way the program should be built
	if lang.SingleChard {
		buildSingeChardRegister(program, lang, s)
	}
}

// build the program without formating
// from a given code (snipplet)
// only assemble the register -> don't run
// functions
func buildSingeChardRegister(program string, lang Lang, s *Stack) Register {

	out := Register{}

	// Iterate over the program code
	for n, token := range program {

		// if isWhiteSpace -> skip
		if !isWhiteSpace(token) {

			// Convert type
			bToken := byte(token)

			// Check lang
			switch bToken {
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
				printCell(s)

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

				// trim the code ->
				// pass only the remaining code
				remainingCode := program[n:len(program)]

				// Innit loop
				loop := Loop{
					Lang: lang,
				}

				out.merge(loop.Build(s, &remainingCode))

				break
			default:
				fmt.Println("Unknown character", token)
				return Register{}
			}
		}
	}

	return out
}

// Check if given rune is whitespace
func isWhiteSpace(r rune) bool {
	if r == '\t' || r == '\n' || unicode.IsSpace(r) {
		return true
	}
	return false
}
