package cases

import "strings"

func Camel(s string) string {
	ss := fields(s)
	for i, s := range ss {
		ss[i] = strings.ToLower(s)
		if i != 0 {
			ss[i] = strings.Title(s)
		}
	}
	return strings.Join(ss, "")
}
