package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

const layoutStr = "2006-01-02T15:04:05Z"

func GetToEpochCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "to_epoch",
		Short: "Convert to epoch time",
		Run: func(cmd *cobra.Command, args []string) {
			timestampStr, _ := cmd.Flags().GetString("timestamp")
			seconds, _ := cmd.Flags().GetBool("short")
			t := time.Now()
			if len(timestampStr) != 0 {
				timestamp, err := time.Parse(layoutStr, timestampStr)
				if err != nil {
					panic(fmt.Errorf("incorrect timestamp str: %s", timestampStr))
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

func GetFromEpochCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "from_epoch",
		Short: "Convert from epoch time",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				panic("No epoch timestamp is provided!")
			}
			epochTimestamp := args[len(args)-1]
			epochInt, _ := strconv.ParseInt(epochTimestamp, 10, 64)
			unixTimeUTC := time.Unix(epochInt, 0)
			fmt.Println(unixTimeUTC.Format(layoutStr))
		},
	}
	return command
}
