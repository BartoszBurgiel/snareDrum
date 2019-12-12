package interpreter

// Lang type representes all tokens
// needed to execute a program
type Lang struct {

	// If all tokens are single character
	SingleChard bool

	// Move pointer up and down
	Pointer struct {
		Up   interface{}
		Down interface{}
	}

	// Input and output
	IO struct {
		In  interface{}
		Out interface{}
	}

	// Single cell in a turing machine
	Cell struct {

		// Add n to the cell
		Add interface{}

		// Subtract n from the cell
		Sub interface{}
	}

	// Loop
	Loop struct {

		// Token to begin loop
		Start interface{}

		// Token to end loop
		End interface{}
	}
}
