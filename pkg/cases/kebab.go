package cases

import "strings"

func Kebab(s string) string {
	return strings.ToLower(strings.Join(fields(s), "-"))
}
