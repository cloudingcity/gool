package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cloudingcity/gool/pkg/timestamp"
)

var timestampToDateCmd = &cobra.Command{
	Use:                   "ts2date [timestamp]",
	Short:                 "Covert unix timestamp to date (default: current time)",
	Args:                  cobra.MaximumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			fmt.Println(time.Now().Format(time.RFC3339))
			return
		}
		unix, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Invalid timestamp!")
			return
		}
		date := timestamp.ToDate(unix)
		fmt.Println(date.Format(time.RFC3339))
	},
}
