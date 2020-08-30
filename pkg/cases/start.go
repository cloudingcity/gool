package cases

import (
	"strings"
)

func Start(s string) string {
	return strings.Title(strings.Join(fields(s), " "))
}
