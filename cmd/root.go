// Package cmd handles the command-line interface for Polyhymnia.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands.
var (
	// Build date is rewritten at build time by the make file.
	BuildDate = "YYYY-MM-DDTHH:MM:SSSZ"
	// Version is rewritten at build time by the make file.
	Version = "1.0.0"
	// RootCmd will execute the datamuse command logic by default.
	RootCmd = &cobra.Command{
		Use:     "polyhymnia [search term]",
		Short:   "Polyhymnia enables users to search for words\nbased on meaning, sound, spelling, and relationships.",
		Long:    "Polyhymnia leverages the Datamuse API to enable users to search for words\nbased on meaning, sound, spelling, and relationships.",
		Version: fmt.Sprintf("%s - Build Date: %s", Version, BuildDate),
		RunE:    runDatamuseQuery,
	}
)

// Execute adds all child commands to the root command and sets flags.
func Execute() error {
	if err := RootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
