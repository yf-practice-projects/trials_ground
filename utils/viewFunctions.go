package utils

import "time"

//time.Time→文字列日付(yyyy/MM/dd mm:hh:ss)
func ConvertTimeToDateTime(t time.Time) string {
	timeStr := t.Format("2006/01/02 15:04:05")
	return timeStr
}
