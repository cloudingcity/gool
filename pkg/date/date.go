package date

import (
	"fmt"

	"github.com/araddon/dateparse"
)

func ToTimestamp(date string) (int64, error) {
	time, err := dateparse.ParseAny(date)
	if err != nil {
		return 0, fmt.Errorf("date parse failed: %w", err)
	}
	return time.Unix(), nil
}
