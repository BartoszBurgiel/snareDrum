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

	// Here will all prcessed packages land
	pPack := make(chan ProcessedMaterial, 10000)

	// Fire go routines
	for i, pack := range packages {

		wg.Add(1)
		go func(index int, material *bytes.Buffer, processedPackages chan ProcessedMaterial, wg *sync.WaitGroup) {

			// Dummy progress for gen function
			dummyProgress := 0
			cont := gen(lang, material.String(), &dummyProgress)

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
		for {
			ui.PrintProgressBar(len(pPack), length-1)
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
		ui.PrintProgressBar(i, length-1)
	}
	fmt.Println("")
}
