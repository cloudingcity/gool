package cases

import (
	"strings"
)

func Upper(s string) string {
	return strings.ToUpper(strings.Join(fields(s), " "))
}
