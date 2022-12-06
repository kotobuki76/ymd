package ymd

import "time"

// type for YYYYMMDDHHMMSS
type Ymdhms string

// convert to string
func (h Ymdhms) String() string {
	return (string)(h)
}

// Get today's "YYYYMMDDHHMMSS" string
func TodayYmdhms() Ymdhms {
	now := time.Now().UTC()
	return Ymdhms(now.Format(YYYYMMDDHHMMSS))
}
