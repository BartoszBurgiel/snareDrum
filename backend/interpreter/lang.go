package interpreter

// Lang type representes all tokens
// needed to execute a program
type Lang struct {

	// If all tokens are single character
	SingleChard bool

	// Move pointer up and down
	Pointer struct {
		Up   string
		Down string
	}

	// Input and output
	IO struct {
		In  string
		Out string
	}

	// Single cell in a turing machine
	Cell struct {

		// Add n to the cell
		Add string

		// Subtract n from the cell
		Sub string
	}

	// Loop
	Loop struct {

		// Token to begin loop
		Start string

		// Token to end loop
		End string
	}
}
