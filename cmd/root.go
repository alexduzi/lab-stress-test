/*
Copyright © 2026 Alex Duzi <duzihd@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stress",
	Short: "Stress is a CLI tool to test http calls using concurrency",
	Long:  "Stress makes a N amount of calls to a given URL and it also accepts the number of concurrent calls",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
