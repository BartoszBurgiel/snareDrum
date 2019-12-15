package interpreter

import (
	"fmt"
	"strconv"
)

// Debug stack and return log
func (s Stack) Debug() string {
	log := "Number of the cells: "
	log += strconv.Itoa(s.Cells.Number) + "\n"

	// Overview over all cells
	log += "\nAll cells:\n\n"

	tempCells := [][]string{}
	// Convert all cells to data for table
	for i, cell := range s.Cells.Cells {
		tempCells = append(tempCells, cellToRow(cell, i))
	}

	// Print table
	log += formatTable([]string{"Index", "Value", "act. Value"}, tempCells)

	return log
}

// cell info to data struct for table
func cellToRow(c Cell, index int) []string {
	out := []string{}

	// Index of the cell
	out = append(out, strconv.Itoa(index))

	// Value of the cell
	out = append(out, strconv.Itoa(c.Value))

	// Actual character
	out = append(out, string(byte(c.Value)))

	return out
}

// Build a table from given column names and data
func formatTable(colName []string, data [][]string) string {

	out := ""

	//Table head
	for i := 0; i < len(colName); i++ {

		// Add seperating bar -> not on first and last column
		if i > 0 && i < len(colName) {
			out += "|"
		}

		out += fmt.Sprintf("%-10s", colName[i])
	}
	out += "\n"

	// Add seperator
	for i := 0; i < len(colName)*10+len(colName)-1; i++ {
		out += "-"

	}
	out += "\n"

	// Iterate over data (rows)
	for i := 0; i < len(data); i++ {

		// Write each string (columns)
		for j := 0; j < len(data[i]); j++ {
			// Add seperating bar -> not on first and last column
			if j > 0 && j < len(data[i]) {
				out += "|"
			}

			out += fmt.Sprintf("%-10s", data[i][j])
		}

		// newline
		out += "\n"
	}

	return out
}
