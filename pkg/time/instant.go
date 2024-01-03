package time

import "time"

/* __________________________________________________ */

type Instant struct {
	time.Time
}

func NewInstant(time time.Time) Instant {
	return Instant{
		Time: time.UTC(),
	}
}

func NowInstant() Instant {
	return NewInstant(time.Now())
}

func (i Instant) String() string {
	return i.Format(time.RFC3339Nano)
}

/* __________________________________________________ */
