package ui

import (
	"fmt"
)

// PrintProgressBar returns string with a progress bar in ascii with the progress in %
func PrintProgressBar(curr, max int) {
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

	fmt.Printf("%s\r", out)

}
