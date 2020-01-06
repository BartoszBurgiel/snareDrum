package ui

import (
	"fmt"
)

// printProgressBar returns string with a progress bar in ascii with the progress in %
func printProgressBar(curr, max int) string {
	out := "["

	// Map ratio of curr to max to width
	barLength := int(20 * float64(curr) / float64(max))
	for i := 0; i < 20; i++ {
		if i < barLength {
			out += "#"
		} else {
			out += "_"
		}
	}
	out += "]"
	out += fmt.Sprintf("%d%% (%d/%d)", int(100*(float64(curr)/float64(max))), curr, max)

	return out
}

// ProgressBar runs the progress bar function
// and generator from snaredrum concurrently
func ProgressBar(progress *int, max int, keyword string) {

	// Prepare wait group
	fmt.Println("\n", keyword, "...")

	// Fire goroutine for the progress bar

	tempProgress := 1

	fmt.Printf("%s\r", printProgressBar(0, max))

	// Listen to the pointer
	for {
		// Check if progress changed
		if tempProgress != *progress {
			// Reassign progress
			tempProgress = *progress

			// Print progress bar
			fmt.Printf("%s\r", printProgressBar(*progress, max))
		}
	}
}
