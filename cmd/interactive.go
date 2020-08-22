package cmd

import (
	"os"

	"github.com/cloudingcity/gool/internal/interactive"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:                   "interactive",
	Aliases:               []string{"i"},
	Short:                 "Start interactive mode",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		i := interactive.New(os.Stdin, os.Stdout)
		i.Register(commands()...)
		i.Run()
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}
