package main

import (
	"fmt"
	"forevermzm/go-utils/pkg/commands"

	"github.com/spf13/cobra"
)

func main() {
	// Create a new root command
	rootCmd := &cobra.Command{
		Use:   "utils",
		Short: "A collection of utils CLI commands",
	}

	rootCmd.AddCommand(commands.GetToEpochCommand())
	rootCmd.AddCommand(commands.GetFromEpochCommand())

	// Create a new subcommand for saying goodbye
	goodbyeCmd := &cobra.Command{
		Use:   "goodbye",
		Short: "Say goodbye",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Goodbye!")
		},
	}

	// Add the goodbye command to the root command
	rootCmd.AddCommand(goodbyeCmd)

	// Execute the root command
	rootCmd.Execute()
}
