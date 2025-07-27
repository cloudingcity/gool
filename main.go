package main

import (
	"fmt"
	"os"

	"github.com/cloudingcity/gool/cmd"
)

var (
	version = "dev"
)

func main() {
	cmd.Root.Version = version
	if err := cmd.Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
