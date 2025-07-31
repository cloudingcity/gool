package uuid

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

	t.Run("valid count", func(t *testing.T) {
		tests := []struct {
			count int
		}{
			{count: 1},
			{count: 5},
			{count: 10},
		}

		for _, tt := range tests {
			uuids := Generate(tt.count)
			assert.Len(t, uuids, tt.count)

			for _, uuid := range uuids {
				assert.True(t, uuidRegex.MatchString(uuid), "Each UUID should match UUIDv4 format")
				assert.Len(t, uuid, 36, "Each UUID should be 36 characters long")
			}

			// Ensure all UUIDs are unique
			uniqueUUIDs := make(map[string]bool)
			for _, uuid := range uuids {
				assert.False(t, uniqueUUIDs[uuid], "All UUIDs should be unique")
				uniqueUUIDs[uuid] = true
			}
		}
	})

	t.Run("invalid count", func(t *testing.T) {
		tests := []struct {
			count int
		}{
			{count: 0},
			{count: -1},
			{count: -10},
		}

		for _, tt := range tests {
			uuids := Generate(tt.count)
			assert.Empty(t, uuids)
		}
	})
}
