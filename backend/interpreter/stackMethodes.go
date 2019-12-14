package interpreter

import "fmt"

// Add one empty cell to the stack
func addEmptyCell(s *Stack) int {
	s.Cells.Cells = append(s.Cells.Cells, Cell{0})
	s.Cells.Number++

	return 0
}

// Increase value of the current cell by one
func addToCell(s *Stack) int {
	// If pointer didn't move
	if s.Cells.Number == 0 {
		addEmptyCell(s)
	}

	s.Cells.Cells[s.MemoryPointer.Pos].Incr()
	return 0
}

// Decrease value of the current cell by one
func subFromCell(s *Stack) int {
	s.Cells.Cells[s.MemoryPointer.Pos].Decr()
	return 0
}

// Increase the memory pointer by one
// if no cells -> create an empty cell
func increaseMemoryPointer(s *Stack) int {
	s.MemoryPointer.Incr()

	// Check the cells number
	if s.MemoryPointer.Pos >= s.Cells.Number {
		addEmptyCell(s)
	}
	return 0
}

func decreaseMemoryPointer(s *Stack) int {
	s.MemoryPointer.Decr()
	// Check if illegal pointer
	if s.MemoryPointer.Pos < 0 {
		fmt.Println("Illegal memory pointer of", s.MemoryPointer.Pos)
		return 0
	}
	return 0
}

// Print the value of the current cell
func printCell(s *Stack) int {
	return s.Cells.Cells[s.MemoryPointer.Pos].Value
}

// Get cell the memory pointer is on currently
func getCell(s *Stack) *Cell {
	return &s.Cells.Cells[s.MemoryPointer.Pos]
}
