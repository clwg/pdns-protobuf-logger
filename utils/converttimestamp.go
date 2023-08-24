package utils

import (
	"time"
)

// ConvertToTimestamp takes seconds and microseconds and returns a time.Time value
func ConvertToTimestamp(TimeSec uint32, TimeUsec uint32) time.Time {
	// Convert uint32 values to int64
	TimeSec64 := int64(TimeSec)
	TimeUsec64 := int64(TimeUsec)

	// Convert microseconds to nanoseconds as Go expects nanoseconds
	TimeNsec := TimeUsec64 * 1000

	// Create time.Time value from seconds and nanoseconds
	return time.Unix(TimeSec64, TimeNsec)
}
