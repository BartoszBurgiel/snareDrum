package interpreter

import "fmt"

// AddEmptyCell to the stack
func AddEmptyCell(s *Stack) int {
	s.Cells.Cells = append(s.Cells.Cells, Cell{0})
	s.Cells.Number++

	return 0
}

// AddToCell value of the current cell by one
func AddToCell(s *Stack) int {
	// If pointer didn't move
	if s.Cells.Number == 0 {
		AddEmptyCell(s)
	}

	s.Cells.Cells[s.MemoryPointer.Pos].Incr()
	return 0
}

// SubFromCell value of the current cell by one
func SubFromCell(s *Stack) int {
	s.Cells.Cells[s.MemoryPointer.Pos].Decr()
	return 0
}

// IncreaseMemoryPointer by one
// if no cells -> create an empty cell
func IncreaseMemoryPointer(s *Stack) int {
	s.MemoryPointer.Incr()

	// Check the cells number
	if s.MemoryPointer.Pos >= s.Cells.Number {
		AddEmptyCell(s)
	}
	return 0
}

// DecreaseMemoryPointer by one
func DecreaseMemoryPointer(s *Stack) int {
	s.MemoryPointer.Decr()
	// Check if illegal pointer
	if s.MemoryPointer.Pos < 0 {
		fmt.Println("Illegal memory pointer of", s.MemoryPointer.Pos)
		return 0
	}
	return 0
}

// PrintCell the value of the current cell
func PrintCell(s *Stack) int {
	return s.Cells.Cells[s.MemoryPointer.Pos].Value
}

// ReadToCell user input
func ReadToCell(s *Stack) int {
	return 0
}

// Get cell the memory pointer is on currently
func getCell(s *Stack) *Cell {
	return &s.Cells.Cells[s.MemoryPointer.Pos]
}
