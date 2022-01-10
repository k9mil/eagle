package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eagle",
	Short: "A simple, fast and elegant Stack Overflow search(er).",
	Long:  `A simple, fast, and fun CLI-based application which functions as a helper to find answers to your programming questions! `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
