package util

import "time"

func GetNowFormat() string {
	utc := time.FixedZone("UTC", 0)
	now := time.Now().In(utc)
	return now.Format("20060102150405")
}

func GetNow() time.Time {
	// utc := time.FixedZone("UTC+7", 7*60*60)
	return time.Now().In(time.FixedZone("UTC+7", 7*60*60)).UTC()
}
