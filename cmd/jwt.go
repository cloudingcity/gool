package cmd

import (
	"fmt"

	"github.com/cloudingcity/gool/pkg/jwt"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var jwtDecodeCmd = &cobra.Command{
	Use:                   "jwt-decode [JWT]",
	Short:                 "Decode JWT",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := jwt.Decode(args[0])
		if err != nil {
			fmt.Println("Invalid token!")
			return
		}
		fmt.Print(string(pretty.Color(pretty.Pretty([]byte(data)), nil)))
	},
}
