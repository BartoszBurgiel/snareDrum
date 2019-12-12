package interpreter

// Loop represents a loop on a program
type Loop struct {

	// Lang for reference and code trimming
	Lang Lang
}

// trimCode and leave only loop code
// @param -> prog string
// represents the code snipplet
// from the beginning of the loop
// untill the end of the program
func (l *Loop) trimCode(prog *string) string {
	// Number of all loop openings -> all subloops
	nOpening := 0

	snipplet := []byte{}

	// Iterate over code
	for _, token := range *prog {

		// Detect levels
		if byte(token) == l.Lang.Loop.Start[0] {
			nOpening++
		} else if byte(token) == l.Lang.Loop.End[0] {
			nOpening--

			// If on highest level
			if nOpening == 0 {
				// End
				return string(snipplet)
			}
		}

		snipplet = append(snipplet, byte(token))
	}

	return "fail"
}

// Build the loop code snipplet
func (l *Loop) Build(s *Stack, remCode *string) Register {

	out := Register{}

	// trim the code
	code := l.trimCode(remCode)

	// Produce functions untill condition value == 0
	condition := getCell(s).Value

	// Pointer of the cells' value
	initPointerPosition := s.MemoryPointer.Pos

	// Iterate
	for condition >= 0 {

		out.merge(buildSingeChardRegister(code, l.Lang, s))

		// Decr condition
		condition--

		// Get current pointer position
		currentPointerPosition := s.MemoryPointer.Pos

		out.merge(pointerRetrack(initPointerPosition, currentPointerPosition))
	}

	return out
}

// Assemble register that retracts the pointer to the initial
// poisiton
func pointerRetrack(initPos, currPos int) Register {
	out := Register{}

	if initPos == currPos {
		return out
	} else if currPos > initPos {
		for currPos > initPos {
			out.add(decreaseMemoryPointer)
			currPos--
		}
	} else if currPos < initPos {
		for currPos < initPos {
			out.add(increaseMemoryPointer)
			currPos++
		}
	}

	return out
}
