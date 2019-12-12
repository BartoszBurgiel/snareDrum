package interpreter

import "fmt"

// Add one empty cell to the stack
func addEmptyCell(s *Stack) {
	s.Cells.Cells = append(s.Cells.Cells, Cell{0})
	s.Cells.Number++
}

// Increase value of the current cell by one
func addToCell(s *Stack) {
	// If pointer didn't move
	if s.Cells.Number == 0 {
		addEmptyCell(s)
	}

	s.Cells.Cells[s.MemoryPointer.Pos].Value++
}

// Decrease value of the current cell by one
func subFromCell(s *Stack) {
	s.Cells.Cells[s.MemoryPointer.Pos].Value--
}

// Increase the memory pointer by one
// if no cells -> create an empty cell
func increaseMemoryPointer(s *Stack) {
	s.MemoryPointer.Pos++

	// Check the cells number
	if s.MemoryPointer.Pos > s.Cells.Number {
		addEmptyCell(s)
	}
}

func decreaseMemoryPointer(s *Stack) {
	s.MemoryPointer.Pos--

	// Check if illegal pointer
	if s.MemoryPointer.Pos < 0 {
		fmt.Println("Illegal memory pointer of", 0)
		return
	}
}

// Print the value of the current cell
func printCell(s *Stack) {
	s.Cells.Cells[s.MemoryPointer.Pos].Print()
}
