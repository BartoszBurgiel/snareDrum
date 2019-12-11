package interpreter

import "fmt"

// Cell defines a single cell in the turing machine
type Cell struct {
	Value int
}

// Incr the cells value by one
func (c *Cell) Incr() {
	c.Value++
}

// Decr the cells value by one
func (c *Cell) Decr() {
	c.Value--
}

// Print the value of the cell
func (c Cell) Print() {
	fmt.Println(c.Value)
}
