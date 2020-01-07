package concurrency

import (
	"bytes"
	"fmt"
	"snareDrum/backend/ui"
	"time"
)

// DivideToPackages given material into packages each containing 1000 bytes
func DivideToPackages(material []byte) (out []*bytes.Buffer) {
	length := len(material)

	fmt.Println("\nDivide to packages...")
	operationTimestamp := time.Now()
	// Iterate over material and create packages
	for i := 0; i <= int(length/10000); i++ {
		// If there's a leftover
		if (i+1)*10000 < length {
			out = append(out, bytes.NewBuffer(material[i*10000:(i+1)*10000]))
		} else {
			out = append(out, bytes.NewBuffer(material[i*10000:length]))
		}
		ui.PrintProgressBar(i, int(length/10000), operationTimestamp)
		operationTimestamp = time.Now()
	}
	fmt.Printf("\n")
	return out
}

// If a struct satisfies this interface it can be sorted
type sortable interface {
	getValue() int
}

// Sort data with bubblesort
func sort(data []ProcessedMaterial) {
	// Iterate over data
	for i := 0; i < len(data); i++ {
		for j := len(data) - 1; j >= i+1; j-- {
			if (data)[j].Index < (data)[j-1].Index {
				(data)[j], (data)[j-1] = (data)[j-1], (data)[j]
			}
		}
	}
}
