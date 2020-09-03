package cmd

import (
	"crypto/md5"
	"fmt"

	"github.com/spf13/cobra"
)

var md5Cmd = &cobra.Command{
	Use:                   "md5 [text]",
	Short:                 "Computes the checksum of your text",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		data := []byte(args[0])
		fmt.Printf("%x\n", md5.Sum(data))
	},
}
