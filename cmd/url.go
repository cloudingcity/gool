package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

var urlEncodeCmd = &cobra.Command{
	Use:                   "url-encode [text]",
	Short:                 "Encode url",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(url.QueryEscape(args[0]))
	},
}

var urlDecodeCmd = &cobra.Command{
	Use:                   "url-decode [text]",
	Short:                 "Decode url",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := url.QueryUnescape(args[0])
		if err != nil {
			fmt.Println("Invalid query")
		}
		fmt.Println(s)
	},
}
