package cases

import (
	"strings"
)

func Lower(s string) string {
	return strings.ToLower(strings.Join(fields(s), " "))
}
