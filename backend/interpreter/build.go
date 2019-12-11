package interpreter

// Build stack based on the .sd file to 
func (s *Stack) Build(program string, lang Lang) {
	// Determine in what way the program should be built
	if lang.SingleChard {
		s.buildSingeChard(program, lang)
	}
}

// build the program without formating
func (s *Stack) buildSingeChard(program string, lang Lang) {
	// Iterate over the program code
	for _, token := range program {

		// if isWhiteSpace -> skip
		if !isWhiteSpace(token) {

			// Check lang
			switch token {
				case lang.Pointer.Up:
					s.increaseMemoryPointer()
					
					// Add to register
					s.addAction(s.increaseMemoryPointer)
					break
				case lang.Pointer.Down:
					s.decreaseMemoryCell()

					// Add to register
					s.addAction(s.decreaseMemoryPointer)
					break
				case lang.IO.In:
					break
				case lang.IO.Out:
					s.printCell()

					// Add to register
					s.addAction(s.printCell)
					break
				case lang.Cell.Add:
					s.addToCell()

					// Add to register
					s.addAction(s.addToCell)
					break
				case lang.Cell.Sub:
					s.subFromCell()

					// Add to register
					s.addAction(s.subFromCell)
					break
				case lang.Loop.Start:
					break
				case lang.Loop.End:
					break
				default: 
					fmt.Println("Unknown character", token)
					return
				} 
			}
		}
	}
}

// Check if given rune is whitespace
func isWhiteSpace(r rune) bool {
	if r == '\t' || r == '\n' || unicode.IsSpace(r) {
		return true
	}
	return false
}
