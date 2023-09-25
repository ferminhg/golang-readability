package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "readability",
	Short: "app is a CLI application",
	Run: readibilityCmd,
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
