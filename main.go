// Package main acts as the entrypoint to Polyhymnia.
package main

import (
	"log"

	"github.com/pierow2k/polyhymnia/cmd"
)

// runApp executes the datamuse command as the root command.
//
//nolint:wrapcheck
func runApp() error {
	return cmd.Execute()
}

// main calls runApp and handle any errors.
func main() {
	if err := runApp(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
