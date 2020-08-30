package cases

import (
	"unicode"
)

func fields(s string) []string {
	var (
		res []string
		pos int
		pre int
	)
	for i, r := range s {
		cur := parseCase(r)

		switch pre {
		case invalidCase:
			if pre != cur {
				pos = i
			}
		case numberCase, lowerCase:
			if pre != cur {
				res = append(res, s[pos:i])
				pos = i
			}
		case upperCase:
			switch cur {
			case lowerCase:
				if i-2 >= 0 && unicode.IsUpper(rune(s[i-2])) {
					res = append(res, s[pos:i-1])
					pos = i - 1
				}
			case numberCase, invalidCase:
				res = append(res, s[pos:i])
				pos = i
			}
		}
		if i == len(s)-1 && cur != invalidCase {
			res = append(res, s[pos:])
			break
		}
		pre = cur
	}
	return res
}

const (
	invalidCase = iota + 1
	upperCase
	lowerCase
	numberCase
)

func parseCase(r rune) int {
	switch {
	case unicode.IsNumber(r):
		return numberCase
	case unicode.IsLower(r):
		return lowerCase
	case unicode.IsUpper(r):
		return upperCase
	default:
		return invalidCase
	}
}
