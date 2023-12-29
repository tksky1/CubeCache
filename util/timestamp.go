package util

import "time"

type Timestamp struct {
	Second     int64 // unix second
	NanoSecond int64 // nanoSecond left
}

func (t *Timestamp) Time() time.Time {
	return time.Unix(t.Second, t.NanoSecond)
}

func NewTimestamp(second int64, nanoSecond int64) *Timestamp {
	return &Timestamp{
		Second:     second,
		NanoSecond: nanoSecond,
	}
}

func NewTimestampForNow() *Timestamp {
	return &Timestamp{
		Second:     time.Now().Unix(),
		NanoSecond: int64(time.Now().Nanosecond()),
	}
}
