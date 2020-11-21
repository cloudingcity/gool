package cmd

import (
	"fmt"
	"net/url"

	goolurl "github.com/cloudingcity/gool/pkg/url"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
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

var urlToJSON = &cobra.Command{
	Use:                   "url-to-json [text]",
	Short:                 "Convert url to JSON",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		u, err := goolurl.Parse(args[0])
		if err != nil {
			fmt.Println("Invalid query")
		}
		fmt.Println(string(pretty.Color(pretty.Pretty(u.ToJSON()), nil)))
	},
}
