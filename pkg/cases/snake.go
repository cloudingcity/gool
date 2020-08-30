package cases

import (
	"strings"
)

func Snake(s string) string {
	return strings.ToLower(strings.Join(fields(s), "_"))
}
