package interpreter

import (
	"fmt"
	"unicode"
)

// Build stack based on the .sd file to
func (s *Stack) Build(program string, lang Lang) {
	// Determine in what way the program should be built
	if lang.SingleChard {
		reg := buildSingeChardRegister(program, lang)
		s.Register.Methods = append(s.Register.Methods, reg.Methods...)
	}
}

// build the program without formating
// from a given code (snipplet)
// only assemble the register -> don't run
// functions
func buildSingeChardRegister(program string, lang Lang) Register {

	out := Register{}

	// Iterate over the program code
	for _, token := range program {

		// if isWhiteSpace -> skip
		if !isWhiteSpace(token) {

			// Check lang
			switch token {
			case lang.Pointer.Up.(rune):

				// Add to register
				out.add(increaseMemoryPointer)
				break
			case lang.Pointer.Down.(rune):

				// Add to register
				out.add(decreaseMemoryPointer)
				break
			case lang.IO.In.(rune):
				break
			case lang.IO.Out.(rune):

				// Add to register
				out.add(printCell)
				break
			case lang.Cell.Add.(rune):

				// Add to register
				out.add(addToCell)
				break
			case lang.Cell.Sub.(rune):

				// Add to register
				out.add(subFromCell)
				break
			case lang.Loop.Start.(rune):
				break
			case lang.Loop.End.(rune):
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
