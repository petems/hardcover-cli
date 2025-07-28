package cmd

import (
	"fmt"
)

// printToStdoutf safely prints to stdout without checking errors (for CLI output).
func printToStdoutf(w interface{ Write([]byte) (int, error) }, format string, args ...interface{}) {
	//nolint:errcheck // CLI output errors are not critical
	fmt.Fprintf(w, format, args...)
}

// printToStdoutLn safely prints a newline to stdout without checking errors.
func printToStdoutLn(w interface{ Write([]byte) (int, error) }, args ...interface{}) {
	//nolint:errcheck // CLI output errors are not critical
	fmt.Fprintln(w, args...)
}
