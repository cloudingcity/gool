package uuid

import (
	"github.com/google/uuid"
)

func Generate(count int) []string {
	if count <= 0 {
		return []string{}
	}

	uuids := make([]string, count)
	for i := 0; i < count; i++ {
		uuids[i] = uuid.New().String()
	}
	return uuids
}
