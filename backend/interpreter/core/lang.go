package interpreter

// Lang type representes all tokens
// needed to execute a program
type Lang struct {

	// Move pointer up and down
	Pointer struct {
		Up   string
		Down string
	}

	// Jump prefix is a prefix that comes before
	// the pointer to a memory address
	JumpPrefix string

	// Input and output
	IO struct {
		In  string
		Out string
	}

	// Single cell in a turing machine
	Cell struct {

		// Read the value from address
		Read string

		// Write the value to the cell
		Write string

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
