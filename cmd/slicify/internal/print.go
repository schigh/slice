package internal

import (
	"fmt"
	"os"
)

// PrintErr will print in red
func PrintErr(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "\033[31mslicify: "+msg+"\033[0m\n", args...)
}

// PrintSuccess will print in green
func PrintSuccess(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, "\033[32mslicify: "+msg+"\033[0m\n", args...)
}

// PrintInfo will print in cyan
func PrintInfo(msg string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, "\033[36mslicify: "+msg+"\033[0m\n", args...)
}
