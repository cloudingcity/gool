package timestamp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tests := []struct {
			timestamp int64
			want      string
		}{
			{timestamp: 1598161702, want: "2020-08-23"},
		}

		for _, tt := range tests {
			got := ToDate(tt.timestamp).Format("2006-01-02")
			assert.Equal(t, tt.want, got)
		}
	})
}
