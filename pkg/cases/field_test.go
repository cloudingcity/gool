package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFields(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{s: "Foo Bar", want: []string{"Foo", "Bar"}},
		{s: "--Foo-Bar--", want: []string{"Foo", "Bar"}},
		{s: "-@@-Foo@\\+-BarBaz-!@%^&A-", want: []string{"Foo", "Bar", "Baz", "A"}},
		{s: "123. Foo Bar", want: []string{"123", "Foo", "Bar"}},
		{s: "HTMLDecode", want: []string{"HTML", "Decode"}},
		{s: "FFFooBar", want: []string{"FF", "Foo", "Bar"}},
		{s: "a b", want: []string{"a", "b"}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, fields(tt.s))
	}
}

func BenchmarkFields(b *testing.B) {
	s := "-@@-Foo@\\+-BarBaz-!@%^&A-"
	for i := 0; i < b.N; i++ {
		Snake(s)
	}
}
