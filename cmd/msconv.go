package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cloudingcity/gool/pkg/date"
)

var msconvCmd = &cobra.Command{
	Use:                   "msconv [input]",
	Short:                 "Smart millisecond timestamp/date converter (default: show current time)",
	Args:                  cobra.MaximumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			now := time.Now()
			fmt.Printf("%s\n", now.Format(time.RFC3339))
			fmt.Printf("%s\n", strconv.FormatInt(now.UnixMilli(), 10))
			return
		}

		result, err := smartConvertMs(args[0])
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}
		fmt.Println(result)
	},
}

func smartConvertMs(input string) (string, error) {
	input = strings.TrimSpace(input)

	// Try to parse as timestamp (integer)
	if unix, err := strconv.ParseInt(input, 10, 64); err == nil {
		// Validate timestamp must be non-negative
		if unix >= 0 {
			// Convert millisecond timestamp to date
			date := time.UnixMilli(unix).UTC()
			return date.Format(time.RFC3339), nil
		}
		// If it's a negative number, it's an invalid timestamp
		return "", fmt.Errorf("invalid timestamp: %s must be non-negative", input)
	}

	// Try to parse as date string
	ts, err := date.ToTimestamp(input)
	if err != nil {
		return "", fmt.Errorf("invalid input: not a valid timestamp or date")
	}

	// Convert seconds to milliseconds
	return strconv.FormatInt(ts*1000, 10), nil
}
