package tools

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

func StringToTime(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", s, loc)
	return theTime
}

func StringToDate(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	theDate, _ := time.ParseInLocation("2006-01-02", s, loc)
	return theDate
}

func IsErrRecordNotFound(err error) bool {
	if err == nil {
		return false
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}
