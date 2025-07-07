package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "budget",
	Short: "track your expenses and shi",
	Long:  "a terminal based budget tracker type shi",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(summaryCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(delCmd)
	rootCmd.AddCommand(expCmd)
}
