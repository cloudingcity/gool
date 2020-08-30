package main

import (
	"time"

	"github.com/cloudingcity/gool/cmd"
)

var (
	version = "dev"
	date    = ""
)

func main() {
	cmd.Version = version
	cmd.Date, _ = time.Parse(time.RFC3339, date)
	cmd.Execute()
}
