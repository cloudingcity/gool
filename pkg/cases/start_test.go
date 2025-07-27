package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "foo bar", want: "Foo Bar"},
		{s: "--foo-bar--", want: "Foo Bar"},
		{s: "-@@-Foo@\\+-BarBaz-!@%^&A-", want: "Foo Bar Baz A"},
		{s: "123. Foo Bar", want: "123 Foo Bar"},
		{s: "HTMLDecode", want: "Html Decode"},
		{s: "FFFooBar", want: "Ff Foo Bar"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Start(tt.s))
	}
}
