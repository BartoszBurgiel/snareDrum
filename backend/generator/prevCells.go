package generator

import "snareDrum/backend/interpreter"

// prevCells represents the last n previous cells
// with their values and their distance from the current counter
type prevCells struct {
	cells []byte
}

// Add new cell to the list and remove the last one
func (p *prevCells) add(value byte) {

	// Add new cell
	p.cells = append(p.cells, value)

	// If more than 5 elements
	if len(p.cells) > 5 {
		// remove first cell
		p.cells = p.cells[1:len(p.cells)]
	}
}

// Search for a given byte in the list
// and return its index
func (p prevCells) find(c byte) (index int) {
	for index, cL := range p.cells {
		if cL == c {
			return index
		}
	}

	return -1
}

// Retract the pointer by a given value, print the cell's value
// and return the pointer to initial position
func retractPrintReturn(moves int, lang interpreter.Lang) string {
	out := ""
	for i := 0; i <= moves; i++ {
		out += lang.Pointer.Down
	}

	// Print
	out += lang.IO.Out

	// Return pointer
	for i := 0; i <= moves; i++ {
		out += lang.Pointer.Up
	}

	return out
}
