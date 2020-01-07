package interpreter

import (
	"bytes"
	"fmt"
	"snareDrum/backend/ui"
	"strconv"
	"time"
)

// Debug stack and return log
func (s Stack) Debug() bytes.Buffer {
	log := bytes.Buffer{}
	log.WriteString("Number of the cells: ")
	log.WriteString(strconv.Itoa(s.Cells.Number) + "\n")

	// Overview over all cells
	log.WriteString("\nAll cells:\n\n")

	tempCells := [][]string{}
	// Convert all cells to data for table
	fmt.Println("\nGathering cell data...")
	operationTimestamp := time.Now()
	for i, cell := range s.Cells.Cells {
		tempCells = append(tempCells, cellToRow(cell, i))

		// Print the progress bar
		ui.PrintProgressBar(i, len(s.Cells.Cells)-1, operationTimestamp)
		operationTimestamp = time.Now()
	}

	// Get Table
	table := formatTable([]string{"Index", "Value", "act. Value"}, tempCells)

	// Write table
	log.WriteString(table.String())

	// Flush table
	table.Reset()
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
func formatTable(colName []string, data [][]string) bytes.Buffer {

	out := bytes.Buffer{}

	//Table head
	for i := 0; i < len(colName); i++ {

		// Add seperating bar -> not on first and last column
		if i > 0 && i < len(colName) {
			out.WriteString("|")
		}

		out.WriteString(fmt.Sprintf("%-10s", colName[i]))
	}
	out.WriteString("\n")

	// Add seperator
	for i := 0; i < len(colName)*10+len(colName)-1; i++ {
		out.WriteString("-")

	}
	out.WriteString("\n")

	fmt.Println("\nBuilding debug table...")
	operationTimestamp := time.Now()
	// Iterate over data (rows)
	for i := 0; i < len(data); i++ {

		// Write each string (columns)
		for j := 0; j < len(data[i]); j++ {
			// Add seperating bar -> not on first and last column
			if j > 0 && j < len(data[i]) {
				out.WriteString("|")
			}

			if data[i][j] == "\n" {
				out.WriteString(fmt.Sprintf("%-10s", "\\n"))
			} else {
				out.WriteString(fmt.Sprintf("%-10s", data[i][j]))
			}
		}

		// newline
		out.WriteString("\n")
		ui.PrintProgressBar(i, len(data)-1, operationTimestamp)
		operationTimestamp = time.Now()
	}

	return out
}
