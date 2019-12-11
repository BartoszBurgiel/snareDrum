package interpreter

import "fmt"

// Add one empty cell to the stack
func (s *Stack) addEmptyCell() {
	s.Cells.Cells = append(s.Cells.Cells, Cell{0})
}

// Increase value of the current cell by one
func (s *Stack) addToCell() {
	s.Cells.Cells[s.MemoryPointer]++
}

// Decrease value of the current cell by one
func (s *Stack) subFromCell() {
	s.Cells.Cells[s.MemoryPointer]--
}

// Add action to the stack register
func (s *Stack) addAction(f func()) {
	s.Actions = append(s.Actions, a)
}

// Increase the memory pointer by one
// if no cells -> create an empty cell
func (s *Stack) increaseMemoryPointer() {
	s.MemoryPointer++

	// Check the cells number
	if s.MemoryPointer > s.Cells.Number {
		s.Cells.Number++
		s.addEmptyCell()
	}
}

func (s *Stack) decreaseMemoryPointer() {
	s.MemoryPointer--

	// Check if illegal pointer
	if s.MemoryPointer < 0 {
		fmt.Println("Illegal memory pointer of", 0)
		return
	}
}

// Print the value of the current cell
func (s *Stack) printCell() {
	s.Cells.Cells[s.MemoryPointer].Print()
}
