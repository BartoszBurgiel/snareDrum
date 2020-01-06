package concurrency

import (
	"bytes"
	"fmt"
	"os"
	"snareDrum/backend/interpreter"
	"snareDrum/backend/ui"
	"sync"
)

// RunTranslate file concurrently
func RunTranslate(gen func(lang interpreter.Lang, progOutput string, progress *int) *bytes.Buffer, lang interpreter.Lang, progOutput string, progress *int) {
	// Divide to packages
	packages := divideToPackages([]byte(progOutput))
	length := len(packages)

	// Processed packages will be appended into this slice
	processedPackages := []ProcessedMaterial{}

	var wg sync.WaitGroup

	// Fire go routines
	for i, pack := range packages {

		wg.Add(1)
		go func(index int, material *bytes.Buffer, processedPackages *[]ProcessedMaterial, wg *sync.WaitGroup) {

			// Dummy progress for gen function
			dummyProgress := 0
			cont := gen(lang, material.String(), &dummyProgress)

			// Build processed struct
			processed := ProcessedMaterial{
				Index:   index,
				Content: cont,
			}

			// Append
			*processedPackages = append(*processedPackages, processed)
			wg.Done()
			// *progress++
		}(i, pack, &processedPackages, &wg)
	}

	// Progress bar
	go func() {
		for {
			ui.PrintProgressBar(len(processedPackages), length-1)
			if len(processedPackages) == length {
				return
			}
		}
	}()
	wg.Wait()
	fmt.Println("\nTranslating done")
	// Sort processed packages
	sort(processedPackages)

	// Write off to file
	file, err := os.Create("GENERATED.sd")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Println("Writing off...")

	for i, pack := range processedPackages {
		ui.PrintProgressBar(i, length)
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
	}
}
