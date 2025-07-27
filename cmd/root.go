package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/cloudingcity/gool/internal/shell"
)

var Root = &cobra.Command{
	Use:           "gool",
	Short:         "Gool make your life easier",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		shell := shell.New(os.Stdin, os.Stdout)
		shell.Register(commands()...)
		shell.SetHistoryPath(os.TempDir())
		shell.Run()
	},
}

func init() {
	cobra.EnableCommandSorting = false
	Root.AddCommand(commands()...)
}

func commands() []*cobra.Command {
	return []*cobra.Command{
		timestampToDateCmd,
		dateToTimestampCmd,
		jwtDecodeCmd,
		md5Cmd,
		urlEncodeCmd,
		urlDecodeCmd,
		urlToJSON,
		base64EncodeCmd,
		base64DecodeCmd,
		randStringCmd,
		countCmd,
		CamelCaseCmd,
		KebabCaseCmd,
		LowerCaseCmd,
		SnakeCaseCmd,
		StartCaseCmd,
		UpperCaseCmd,
		formatJSONCmd,
	}
}
