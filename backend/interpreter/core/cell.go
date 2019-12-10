package interpreter

// Cell defines a single cell in the turing machine
type Cell struct {
	Value int
	Add   func(int, *int)
	Sub   func(int, *int)
}
