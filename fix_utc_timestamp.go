package quickfix

import (
	"time"
)

// TimestampPrecision defines the precision used by FIXUTCTimestamp
type TimestampPrecision bool

const (
	Seconds TimestampPrecision = true
	Millis  TimestampPrecision = false
)

// FIXUTCTimestamp is a FIX UTC Timestamp value, implements FieldValue
type FIXUTCTimestamp struct {
	time.Time
	NoMillis TimestampPrecision
}

const (
	utcTimestampMillisFormat  = "20060102-15:04:05.000"
	utcTimestampSecondsFormat = "20060102-15:04:05"
	utcTimestampMicrosFormat  = "20060102-15:04:05.000000"
	utcTimestampNanosFormat   = "20060102-15:04:05.000000000"
)
