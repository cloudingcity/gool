package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"

	"github.com/cloudingcity/gool/pkg/text"
)

var textEscapeCmd = &cobra.Command{
	Use:                   "text-escape [text]",
	Short:                 "Escape quotes and backslashes in text",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(text.Escape(args[0]))
	},
}

var textUnescapeCmd = &cobra.Command{
	Use:                   "text-unescape [text]",
	Short:                 "Unescape quotes and backslashes in text (pretty prints JSON)",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		result := text.Unescape(args[0])
		if json.Valid([]byte(result)) {
			fmt.Print(string(pretty.Color(pretty.Pretty([]byte(result)), nil)))
		} else {
			fmt.Println(result)
		}
	},
}
