// Package main is the entry point for the hardcover-cli application.
package main

import (
	"context"
	"log/slog"
	"os"

	"hardcover-cli/cmd"
)

// version is set by the build process.
var version = "dev"

// main is the entry point for the hardcover CLI application.
func main() {
	if len(os.Args) > 1 && (os.Args[1] == "version" || os.Args[1] == "--version" || os.Args[1] == "-v") {
		slog.InfoContext(context.Background(), "hardcover version", "version", version)
		return
	}

	cmd.Execute()
}
