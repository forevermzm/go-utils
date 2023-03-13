package commands

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func GetToEpochCommand() *cobra.Command {
	command := &cobra.Command{
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
	command.Flags().StringP("timestamp", "t", "", "Timestamp str in ISO 8601 format. Example 2006-01-02T15:04:05Z")
	command.Flags().BoolP("short", "s", false, "Whether to print epoch time in second")
	return command
}
