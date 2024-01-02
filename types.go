package datetimeparser

import (
	"time"
)

type TimeParser struct {
	timePoints DateTimePoints
	parsedTime ParsedTime

	storedTimeLayout string
	storedDateLayout string

	foundedLocale string

	foundTime        bool
	foundDate        bool
	foundAdverb      bool
	foundMerediem    bool
	foundTimeDetails bool
}

type ParsedTime struct {
	Day   int
	Month int
	Year  int

	Hour   int
	Minute int
	Second int

	Adverb     int
	Meridiem   int
	NamedTZVal int
	TZ         *time.Location
}

type PointsOfElement struct {
	Stored bool
	Start  int
	End    int
}

type DateTimePoints struct {
	Time PointsOfElement
	Date PointsOfElement
}

func (pars *TimeParser) ToUTCTime() (time.Time, bool) {
	if pars.parsedTime.Minute == 0 &&
		pars.parsedTime.Second == 0 &&
		pars.parsedTime.Hour == 0 &&
		pars.parsedTime.Day == 0 &&
		pars.parsedTime.Month == 0 &&
		pars.parsedTime.Year == 0 {
		return time.Time{}, false
	}

	// if PM, just add 12 hours
	if pars.parsedTime.Meridiem == 2 {
		pars.parsedTime.Hour += 12
	}

	// fix -0001-11-30 even if time parsed correctly
	if pars.parsedTime.Day == 0 {
		pars.parsedTime.Day = 1
	}

	if pars.parsedTime.Month == 0 {
		pars.parsedTime.Month = 1
	}

	if pars.parsedTime.Year > 0 && pars.parsedTime.Year <= 99 { // 2-digit num need fix
		pars.parsedTime.Year += 2000
	}

	loc := time.UTC

	if pars.parsedTime.TZ != nil {
		loc = pars.parsedTime.TZ
	}

	resDT := time.Date(
		pars.parsedTime.Year,
		time.Month(pars.parsedTime.Month),
		pars.parsedTime.Day,
		pars.parsedTime.Hour,
		pars.parsedTime.Minute,
		pars.parsedTime.Second,
		0,
		loc,
	)

	if pars.parsedTime.NamedTZVal != 0 {
		resDT = resDT.Add(time.Duration(pars.parsedTime.NamedTZVal) * time.Second)
	}

	if pars.foundAdverb {
		tempDate := time.Now()
		resDT = resDT.AddDate( // -1 to fix 69 and 75 lines
			tempDate.Year(), int(tempDate.Month())-1, tempDate.Day()+pars.parsedTime.Adverb-1)
	}

	return resDT.UTC(), true
}

func (pars *TimeParser) clearValues() {
	pars.parsedTime.Day = 0
	pars.parsedTime.Month = 0
	pars.parsedTime.Year = 0

	pars.parsedTime.Second = 0
	pars.parsedTime.Minute = 0
	pars.parsedTime.Hour = 0

	pars.parsedTime.TZ = nil
	pars.parsedTime.Adverb = 0
	pars.parsedTime.Meridiem = 0
	pars.parsedTime.NamedTZVal = 0

	// not set false foundLocale - may be used
	pars.foundTime = false
	pars.foundDate = false
	pars.foundAdverb = false
	pars.foundMerediem = false
	pars.foundTimeDetails = false
}

func (pars *TimeParser) copyDateTimeValues(tm time.Time) {
	pars.copyDateValues(tm)
	pars.copyTimeValues(tm)
}

func (pars *TimeParser) copyTimeValues(tm time.Time) {
	pars.parsedTime.Hour = tm.Hour()
	pars.parsedTime.Minute = tm.Minute()
	pars.parsedTime.Second = tm.Second()
}

func (pars *TimeParser) copyDateValues(tm time.Time) {
	pars.parsedTime.Year = tm.Year()
	pars.parsedTime.Month = int(tm.Month())
	pars.parsedTime.Day = tm.Day()
}

func (pars *TimeParser) copyTimePoints(left, right int) {
	pars.copyTimePointsTime(left, right)
	pars.copyTimePointsDate(left, right)
}

func (pars *TimeParser) copyTimePointsTime(left, right int) {
	pars.timePoints.Time.Stored = true
	pars.timePoints.Time.Start = left
	pars.timePoints.Time.End = right
}

func (pars *TimeParser) copyTimePointsDate(left, right int) {
	pars.timePoints.Date.Stored = true
	pars.timePoints.Date.Start = left
	pars.timePoints.Date.End = right
}
