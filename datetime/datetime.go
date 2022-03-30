package datetime

import "time"

const (
	RFC3339 = "2006-01-02T15:04:05Z07:00"
)

func GetUtcNow() time.Time {
	return time.Now().UTC()
}

func GetUtcNowStr() string {
	return GetUtcNow().Format(RFC3339)
}
