package concurrency

import (
	"bytes"
)

// Divide given material into packages each containing 1000 bytes
func divideToPackages(material []byte) (out []*bytes.Buffer) {
	length := len(material)

	// Iterate over material and create packages
	for i := 0; i <= int(length/1000); i++ {
		// If there's a leftover
		if (i+1)*1000 < length {
			out = append(out, bytes.NewBuffer(material[i*1000:(i+1)*1000]))
		} else {
			out = append(out, bytes.NewBuffer(material[i*1000:length]))
		}
	}
	return out
}
