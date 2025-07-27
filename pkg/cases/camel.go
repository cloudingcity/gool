package cases

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Camel(s string) string {
	ss := fields(s)
	caser := cases.Title(language.English)
	for i, s := range ss {
		ss[i] = strings.ToLower(s)
		if i != 0 {
			ss[i] = caser.String(s)
		}
	}
	return strings.Join(ss, "")
}
