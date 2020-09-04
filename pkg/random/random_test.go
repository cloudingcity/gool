package random

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("valid length", func(t *testing.T) {
		reg := regexp.MustCompile("^[A-Za-z0-9]+$")
		tests := []struct {
			length int
		}{
			{length: 5},
			{length: 10},
			{length: 19},
		}
		for _, tt := range tests {
			s := String(tt.length)
			assert.Len(t, s, tt.length)
			assert.True(t, reg.MatchString(s))
		}
	})

	t.Run("invalid length", func(t *testing.T) {
		tests := []struct {
			length int
		}{
			{length: 0},
			{length: -10},
			{length: -119},
		}
		for _, tt := range tests {
			s := String(tt.length)
			assert.Empty(t, s)
		}
	})
}
