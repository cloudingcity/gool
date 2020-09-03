package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "gool",
	Short:         "Gool make your life easier",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.AddCommand(shellCmd)
	rootCmd.AddCommand(commands()...)
	rootCmd.AddCommand(versionCmd)
}

func commands() []*cobra.Command {
	return []*cobra.Command{
		timestampToDateCmd,
		dateToTimestampCmd,
		jwtDecodeCmd,
		md5Cmd,
		CamelCaseCmd,
		KebabCaseCmd,
		LowerCaseCmd,
		SnakeCaseCmd,
		StartCaseCmd,
		UpperCaseCmd,
	}
}
