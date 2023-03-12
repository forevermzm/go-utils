package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	// Create a new root command
	rootCmd := &cobra.Command{
		Use:   "utils",
		Short: "A collection of utils CLI commands",
	}

	// Create a new subcommand for greeting
	toEpochCommand := &cobra.Command{
		Use:   "to_epoch",
		Short: "Convert to epoch time",
		Run: func(cmd *cobra.Command, args []string) {
			timestampStr, _ := cmd.Flags().GetString("timestamp")
			seconds, _ := cmd.Flags().GetBool("short")
			t := time.Now()
			if len(timestampStr) != 0 {
				layout := "2006-01-02T15:04:05Z"
				timestamp, err := time.Parse(layout, timestampStr)
				if err != nil {
					panic(fmt.Errorf("In correct timestamp str: %s", timestampStr))
				}
				t = timestamp
			}
			if seconds {
				fmt.Println(t.UnixMilli() / 1000)
			} else {
				fmt.Println(t.UnixMilli())
			}
		},
	}

	toEpochCommand.Flags().StringP("timestamp", "t", "", "Timestamp str in ISO 8601 format. Example 2006-01-02T15:04:05Z")
	toEpochCommand.Flags().BoolP("short", "s", false, "Whether to print epoch time in second")
	// Add the greet command to the root command
	rootCmd.AddCommand(toEpochCommand)

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
