package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var base64EncodeCmd = &cobra.Command{
	Use:                   "base64-encode [text]",
	Short:                 "Base64 encode",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		s := base64.StdEncoding.EncodeToString([]byte(args[0]))
		fmt.Println(s)
	},
}

var base64DecodeCmd = &cobra.Command{
	Use:                   "base64-decode [text]",
	Short:                 "Base64 decode",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := base64.StdEncoding.DecodeString(args[0])
		if err != nil {
			fmt.Println("Invalid text!")
			return
		}
		fmt.Println(string(b))
	},
}
