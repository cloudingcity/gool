package timestamp

import "time"

func ToDate(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}
