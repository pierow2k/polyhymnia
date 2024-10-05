// Package cmd handles the command-line interface for Polyhymnia.
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands.
var (
    BuildDate = "YYYY-MM-DDTHH:MM:SSSZ" // Build date is rewritten at build time.
    Version   = "1.0.0"   // Version is rewritten at build time.

    // RootCmd will execute the fraction logic by default.
    RootCmd = &cobra.Command{
        Use:   "polyhymnia",
        Short: "Query the Datamuse API",
        Version: fmt.Sprintf("%s - Build Date: %s", Version, BuildDate),
        RunE:    runDatamuseQuery, // Run the datamuse command logic by default
    }
)

// Execute adds all child commands to the root command and sets flags.
func Execute() error {
    if err := RootCmd.Execute(); err != nil {
        return err
    }
    return nil
}
