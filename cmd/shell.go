package cmd

import (
	"os"

	"github.com/cloudingcity/gool/internal/shell"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:                   "shell",
	Short:                 "Run Gool Shell",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		shell := shell.New(os.Stdin, os.Stdout)
		shell.Register(commands()...)
		shell.Run()
	},
}
