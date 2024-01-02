package datetimeparser

import (
	"fmt"
	"strconv"
	"time"
	"unicode"
)

func (pars *TimeParser) Parse(inp string) (time.Time, error) {
	pars.clearValues()
	pars.parseTime(inp)

	res, found := pars.ToUTCTime()
	if !found {
		return time.Time{}, fmt.Errorf("no date or time data found in '%s'", inp)
	}

	return res, nil
}

func (pars *TimeParser) parseNearbyDigits(inp string) {
	if len(inp) == 0 {
		return
	}

	var (
		digStart int
		digEnd   int
		started  bool
	)

	for pos, run := range inp {
		if unicode.IsDigit(run) {
			if !started {
				digStart = pos
				started = true
			}
		}

		if !unicode.IsDigit(run) && started {
			started = false
			digEnd = pos

			if dig, err := strconv.Atoi(inp[digStart:digEnd]); err == nil {
				pars.tryToUseFoundedDigit(dig)
			}
		}

		if pos == len(inp)-1 && started {
			started = false
			digEnd = pos + 1

			if dig, err := strconv.Atoi(inp[digStart:digEnd]); err == nil {
				pars.tryToUseFoundedDigit(dig)
			}
		}
	}
}

func (pars *TimeParser) tryToUseFoundedDigit(dig int) {
	if pars.foundDate {
		return
	}

	if pars.parsedTime.Year == 0 && validateLongYear(dig) {
		pars.parsedTime.Year = dig
		return
	}

	if pars.parsedTime.Day == 0 && validateDay(dig) {
		pars.parsedTime.Day = dig
		return
	}

	if pars.parsedTime.Month == 0 && validateMonth(dig) {
		pars.parsedTime.Month = dig
		return
	}

	if pars.parsedTime.Year == 0 && validateShortYear(dig) {
		pars.parsedTime.Year = dig
		return
	}
}
