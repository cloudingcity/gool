package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var formatJSONCmd = &cobra.Command{
	Use:                   "format-json [text]",
	Short:                 "Cleans and format JSON",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(string(pretty.Color(pretty.Pretty([]byte(args[0])), nil)))
	},
}
