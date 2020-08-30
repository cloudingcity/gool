package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamel(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "Foo Bar", want: "fooBar"},
		{s: "--Foo-Bar--", want: "fooBar"},
		{s: "-@@-Foo@\\+-BarBaz-!@%^&A-", want: "fooBarBazA"},
		{s: "123. Foo Bar", want: "123FooBar"},
		{s: "HTMLDecode", want: "htmlDecode"},
		{s: "FFFooBar", want: "ffFooBar"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, Camel(tt.s))
	}
}
