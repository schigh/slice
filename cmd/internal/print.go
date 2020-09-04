package internal

import (
	"fmt"
	"os"
)

// PrintErr will print in red
func PrintErr(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "\033[31mslice: "+msg+"\n", args...)
}

// PrintSuccess will print in green
func PrintSuccess(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "\033[32mslice: "+msg+"\n", args...)
}

// PrintInfo will print in cyan
func PrintInfo(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, "\033[36mslice: "+msg+"\n", args...)
}
