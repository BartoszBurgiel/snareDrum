package concurrency

// // RunBuild function more efficently by dividing it into
// // packages and processing them concurrently with goroutines
// func RunBuild(build func(lang interpreter.Lang, program string, progress *int), prepare func(lang interpreter.Lang, program string) []string,lang interpreter.Lang, program string, progress *int) {
// 	packages := divideToPackages([]byte(program))

// 	// Here will all processed packages land
// 	processedPackages := []ProcessedMaterial{}

// 	// Iterate over packages
// 	for i, pac := range packages {

// 		// Fire a goroutine
// 		go func(index int, pm *[]processedPackages) {
// 			// Process code
// 			processedCode := prepare(lang, program)

// 			// ProcessedRegister
// 			pR := ProcessedRegister{
// 				Index: index,
// 				Content: bu
// 			}
// 		}
// 	}

// }
