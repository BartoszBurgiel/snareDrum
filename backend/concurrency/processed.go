package concurrency

// ProcessedMaterial contains the processed material
// and its index to help to put it back together
type ProcessedMaterial struct {
	Index   int
	Content interface{}
}
