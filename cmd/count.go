package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:                   "count [text]",
	Short:                 "Get the characters length",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(len(strings.TrimSpace(args[0])))
	},
}
