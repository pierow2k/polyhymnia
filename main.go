// Package main acts as the entrypoint to Polyhymnia.
package main

import (
	"log"

	"github.com/pierow2k/polyhymnia/cmd"
)

func main() {
	// Execute the datamuse command as the root command
	if err := cmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
