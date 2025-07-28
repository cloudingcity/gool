package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cloudingcity/gool/pkg/date"
)

var dateToTimestampCmd = &cobra.Command{
	Use:                   "date-to-timestamp [date]",
	Aliases:               []string{"d2t"},
	Short:                 "Covert date to unix timestamp",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if strings.ToLower(args[0]) == "" {
			fmt.Println(time.Now().Unix())
			return
		}

		timestamp, err := date.ToTimestamp(args[0])
		if err != nil {
			fmt.Println("Invalid date!")
			return
		}
		fmt.Println(timestamp)
	},
}
