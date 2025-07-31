package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cloudingcity/gool/pkg/uuid"
)

var uuidCmd = &cobra.Command{
	Use:                   "uuid [count]",
	Short:                 "Generate UUID v4 (default: 1)",
	Args:                  cobra.MaximumNArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		count := 1

		if len(args) > 0 && args[0] != "" {
			var err error
			count, err = strconv.Atoi(args[0])
			if err != nil || count <= 0 {
				fmt.Println("Invalid count! Please provide a positive integer.")
				return
			}
		}

		uuids := uuid.Generate(count)
		for _, u := range uuids {
			fmt.Println(u)
		}
	},
}
