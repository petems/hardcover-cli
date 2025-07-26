package main

import (
	"fmt"
	"os"

	"hardcover-cli/cmd"
)

// version is set by the build process
var version = "dev"

// main is the entry point for the hardcover CLI application
func main() {
	// Set version in cmd package if needed
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("hardcover version %s\n", version)
		os.Exit(0)
	}

	cmd.Execute()
}
