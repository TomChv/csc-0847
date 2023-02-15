package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "backend",
	Short: "Manage project 1 backend",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
