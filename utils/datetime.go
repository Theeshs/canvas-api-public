package utils

import (
	"fmt"
	"time"
)

func ConvertJsonDate(timeString string) (time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02", timeString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Time{}, err
	}

	return parsedTime, nil
}
