package utils

import (
	"time"
)

const DateLayout = "2006-01-02"
const TimeDateLayout = "2006-01-02 15:04:05"

const Duration1Microsecond = time.Duration(1000)
const Duration1Millisecond = Duration1Microsecond * 1000
const Duration1Second = Duration1Millisecond * 1000
const Duration1Minute = Duration1Second * 60
const Duration1Day = Duration1Minute * 60 * 24
const DurationSevenDays = Duration1Day * 7

type YearDay struct {
	Year Year      `json:"year"`
	Day  DayInYear `json:"day"`
}

type Year uint64
type DayInYear uint16

type Period struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func (yearDay *YearDay) ToRUString() string {
	asString := yearDay.asTime().Format(DateLayout)
	return asString
}

func ToURUDateString(timeToConvert *time.Time) string {
	asString := timeToConvert.In(thailandTimeZone).Format(DateLayout)
	return asString
}

func (yearDay *YearDay) asTime() time.Time {
	asTime := time.Date(int(yearDay.Year), time.January, 1, 0, 0, 0, 0, thailandTimeZone)
	daysToAdd := int64(yearDay.Day)
	timeToAdd := time.Duration(int64(Duration1Day) * daysToAdd)
	asTime = asTime.Add(timeToAdd)
	return asTime
}

var thailandTimeZone = getThailandTimeZone()

func getThailandTimeZone() *time.Location {
	location := time.FixedZone(ianaThailandDatabaseTimeZone, 7*60*60) // UTC + 7
	return location
}

const ianaThailandDatabaseTimeZone = "Asia/Bangkok"
