package concurrency

import (
	"bytes"
	"fmt"
	"os"
	"snareDrum/backend/interpreter"
	"snareDrum/backend/ui"
	"sync"
	"time"
)

// RunTranslate file concurrently
func RunTranslate(gen func(lang interpreter.Lang, progOutput string) *bytes.Buffer, lang interpreter.Lang, progOutput string) {
	// Divide to packages
	packages := divideToPackages([]byte(progOutput))
	length := len(packages)

	// Processed packages will be appended into this slice
	processedPackages := []ProcessedMaterial{}

	var wg sync.WaitGroup

	// Here will all prcessed packages land
	pPack := make(chan ProcessedMaterial, 10000)

	// Fire go routines
	for i, pack := range packages {

		wg.Add(1)
		go func(index int, material *bytes.Buffer, processedPackages chan ProcessedMaterial, wg *sync.WaitGroup) {

			// Dummy progress for gen function
			cont := gen(lang, material.String())

			// Build processed struct
			processed := ProcessedMaterial{
				Index:   index,
				Content: cont,
			}

			// Append
			processedPackages <- processed
			wg.Done()
			// *progress++
		}(i, pack, pPack, &wg)
	}

	// Progress bar
	go func() {
		operationTimestamp := time.Now()
		// tempLen := 0
		for {
			// only if size is reasonable to display progress
			if length-1 > 4 {

				// If len pPack changed
				ui.PrintProgressBar(len(pPack), length-1, operationTimestamp)
				// operationTimestamp = time.Now()

				// tempLen = len(pPack)

			}
			if len(pPack) == length {
				return
			}
		}
	}()
	wg.Wait()

	// Fetch processed Packages from channel
	for i := 0; i < len(pPack); i++ {
		processedPackages = append(processedPackages, <-pPack)
	}

	// Sort processed packages
	sort(processedPackages)

	// Write off to file
	file, err := os.Create("GENERATED.sd")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Println("\nWriting off...")

	length = len(processedPackages)
	operationTimestamp := time.Now()
	for i, pack := range processedPackages {
		// Write data to file
		_, err := file.Write(pack.Content.(*bytes.Buffer).Bytes())
		if err != nil {
			fmt.Println(err)
		}

		// Save
		err = file.Sync()
		if err != nil {
			fmt.Println(err)
		}
		ui.PrintProgressBar(i, length-1, operationTimestamp)

		// Define new operation time stamp after one operation
		operationTimestamp = time.Now()

	}
	fmt.Println("")
}
