package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	Version string
	Date    time.Time
)

var versionCmd = &cobra.Command{
	Use:                   "version",
	Short:                 "Print version number of gool",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gool version %s (%s)\n", Version, Date.Format("2006-01-02"))
	},
}
