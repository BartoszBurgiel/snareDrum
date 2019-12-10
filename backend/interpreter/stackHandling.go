package interpreter

import "fmt"

// Add one empty cell to the stack
func (s *Stack) addEmptyCell() {
	s.Cells.Cells = append(s.Cells.Cells, Cell{0})
}

// Print value of the cell on which the pointer is
func (s *Stack) printCell() {
	fmt.Println(s.Cells.Cells[s.MemoryPointer])
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
func (s *Stack) addAction(a Action) {
	s.Actions = append(s.Actions, a)
}
