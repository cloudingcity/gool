package cases

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Start(s string) string {
	caser := cases.Title(language.English)
	return caser.String(strings.Join(fields(s), " "))
}
