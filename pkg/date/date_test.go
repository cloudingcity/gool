package date

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToTimestamp(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tests := []struct {
			date string
			want int64
		}{
			{date: "2020-08-23 13:48:22 +0800 CST", want: 1598161702},
			{date: "Sun, 23 Aug 2020 05:48:22 GMT", want: 1598161702},
		}

		for _, tt := range tests {
			got, err := ToTimestamp(tt.date)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		}
	})

	t.Run("failed", func(t *testing.T) {
		_, err := ToTimestamp("foo")
		assert.Error(t, err)
	})
}
