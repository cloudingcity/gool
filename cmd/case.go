package cmd

import (
	"fmt"

	"github.com/cloudingcity/gool/pkg/cases"
	"github.com/spf13/cobra"
)

var CamelCaseCmd = &cobra.Command{
	Use:                   "camel-case [string]",
	Short:                 "Coverts string to camel case",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cases.Camel(args[0]))
	},
}

var KebabCaseCmd = &cobra.Command{
	Use:                   "kebab-case [string]",
	Short:                 "Coverts string to kebab case",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cases.Kebab(args[0]))
	},
}

var LowerCaseCmd = &cobra.Command{
	Use:                   "lower-case [string]",
	Short:                 "Coverts string to lower case",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cases.Lower(args[0]))
	},
}

var SnakeCaseCmd = &cobra.Command{
	Use:                   "snake-case [string]",
	Short:                 "Coverts string to snake case",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cases.Snake(args[0]))
	},
}

var StartCaseCmd = &cobra.Command{
	Use:                   "start-case [string]",
	Short:                 "Coverts string to start case",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cases.Start(args[0]))
	},
}

var UpperCaseCmd = &cobra.Command{
	Use:                   "upper-case [string]",
	Short:                 "Coverts string to upper case",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cases.Upper(args[0]))
	},
}
