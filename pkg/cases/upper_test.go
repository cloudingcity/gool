package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpper(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "Foo Bar", want: "FOO BAR"},
		{s: "--Foo-Bar--", want: "FOO BAR"},
		{s: "-@@-Foo@\\+-BarBaz-!@%^&A-", want: "FOO BAR BAZ A"},
		{s: "123. Foo Bar", want: "123 FOO BAR"},
		{s: "HTMLDecode", want: "HTML DECODE"},
		{s: "FFFooBar", want: "FF FOO BAR"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Upper(tt.s))
	}
}
