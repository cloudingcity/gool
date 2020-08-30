package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLower(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "Foo Bar", want: "foo bar"},
		{s: "--Foo-Bar--", want: "foo bar"},
		{s: "-@@-Foo@\\+-BarBaz-!@%^&A-", want: "foo bar baz a"},
		{s: "123. Foo Bar", want: "123 foo bar"},
		{s: "HTMLDecode", want: "html decode"},
		{s: "FFFooBar", want: "ff foo bar"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Lower(tt.s))
	}
}
