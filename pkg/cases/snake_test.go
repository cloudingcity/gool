package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnake(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "Foo Bar", want: "foo_bar"},
		{s: "--Foo-Bar--", want: "foo_bar"},
		{s: "-@@-Foo@\\+-BarBaz-!@%^&A-", want: "foo_bar_baz_a"},
		{s: "123. Foo Bar", want: "123_foo_bar"},
		{s: "HTMLDecode", want: "html_decode"},
		{s: "FFFooBar", want: "ff_foo_bar"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Snake(tt.s))
	}
}
