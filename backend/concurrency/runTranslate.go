package concurrency

import (
	"bytes"
	"fmt"
	"snareDrum/backend/interpreter"
	"sync"
)

// RunTranslate file concurrently
func RunTranslate(gen func(lang interpreter.Lang, progOutput string, progress *int) *bytes.Buffer, lang interpreter.Lang, progOutput string, progress *int) {
	// Divide to packages
	packages := divideToPackages([]byte(progOutput))

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

	wg.Wait()
	//fmt.Println(processedPackages)
	fmt.Println(len(processedPackages))
}
