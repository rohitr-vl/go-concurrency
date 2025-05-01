package helpers

import "time"

var tz, _ = time.LoadLocation("Asia/Kolkata")

func GetLocalTime() string {
	/*	if len(format) == 0 {
			format = []string{"3:04:05 PM"}
		}
	*/
	return time.Now().In(tz).Format("3:04:05 PM")
}
