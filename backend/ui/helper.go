package ui

import (
	"fmt"
	"os/exec"
)

// ClearLine from the terminal output
func ClearLine() {
	cmd := exec.Command("echo", "-en", "\"\\r\"ss")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}
