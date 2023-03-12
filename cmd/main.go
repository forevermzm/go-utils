package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	// Create a new root command
	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "A simple CLI for greeting and saying goodbye",
	}

	// Create a new subcommand for greeting
	greetCmd := &cobra.Command{
		Use:   "greet",
		Short: "Greet someone",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			verbose, _ := cmd.Flags().GetBool("verbose")

			if verbose {
				fmt.Printf("Hello, %s!\n", name)
			} else {
				fmt.Println("Hello!")
			}
		},
	}

	// Add flags to the greet command
	greetCmd.Flags().StringP("name", "n", "World", "A name to greet")
	greetCmd.Flags().BoolP("verbose", "v", false, "Whether to print verbose output")

	// Add the greet command to the root command
	rootCmd.AddCommand(greetCmd)

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
