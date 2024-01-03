package time

import "time"

/* __________________________________________________ */

func UTC(timestamp *time.Time) *time.Time {
	if timestamp == nil {
		return nil
	}
	utc := timestamp.UTC()
	return &utc
}

func NowUTC() time.Time {
	return time.Now().UTC()
}

/* __________________________________________________ */
