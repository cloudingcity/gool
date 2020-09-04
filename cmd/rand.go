package cmd

import (
	"fmt"
	"strconv"

	"github.com/cloudingcity/gool/pkg/random"
	"github.com/spf13/cobra"
)

var randStringCmd = &cobra.Command{
	Use:                   "rand-string [length]",
	Short:                 "Generate random string of given length (a-z, A-Z, 0-9)",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		length, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid length!")
			return
		}
		fmt.Println(random.String(length))
	},
}
