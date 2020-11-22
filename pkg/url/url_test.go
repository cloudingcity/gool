package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL_ToJSON(t *testing.T) {
	want := "{\"schema\":\"https\",\"host\":\"github.com\",\"path\":\"/cloudingcity/gool\",\"query\":{\"a\":\"1\",\"b\":\"2\"}}"
	u := &URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "/cloudingcity/gool",
		Query: map[string]string{
			"a": "1",
			"b": "2",
		},
	}

	assert.Equal(t, want, string(u.ToJSON()))
}

func TestParse(t *testing.T) {
	tests := []struct {
		give string
		want *URL
	}{
		{
			give: "https://github.com/cloudingcity/gool",
			want: &URL{
				Scheme: "https",
				Host:   "github.com",
				Path:   "/cloudingcity/gool",
				Query:  map[string]string{},
			},
		},
		{
			give: "https://github.com/cloudingcity/gool?a=1&b=2",
			want: &URL{
				Scheme: "https",
				Host:   "github.com",
				Path:   "/cloudingcity/gool",
				Query: map[string]string{
					"a": "1",
					"b": "2",
				},
			},
		},
	}

	for _, tt := range tests {
		got, _ := Parse(tt.give)
		assert.Equal(t, tt.want, got)
	}
}
