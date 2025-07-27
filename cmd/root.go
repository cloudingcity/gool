package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cloudingcity/gool/internal/shell"
)

var rootCmd = &cobra.Command{
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.AddCommand(commands()...)
	rootCmd.AddCommand(versionCmd)
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
