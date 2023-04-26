package helpers

import "time"

func ConvertDateToUnix(date time.Time) int64 {

	result := date.Unix()

	return result

}
