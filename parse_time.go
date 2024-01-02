package datetimeparser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func (pars *TimeParser) parseTime(inp string) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Printf("Panic: \n(%v)\n, while parsing <%s>\n", rec, inp)
			return
		}
	}()

	// use stored layout from previously processed input
	if pars.timePoints.Time.Stored {
		if pars.timePoints.Time.Start < len(inp) &&
			pars.timePoints.Time.End <= len(inp) { // here '<=' is correct for 'End'
			supposedTime := inp[pars.timePoints.Time.Start:pars.timePoints.Time.End]
			if tm, err := time.Parse(pars.storedTimeLayout, supposedTime); err == nil {
				pars.foundTime = true
				pars.copyTimeValues(tm)
			}
		}
	}

	var foundTimeHere bool

	for pos, run := range inp {
		if pars.foundTime { // if time was already found
			break
		}

		// skip lone letter, because we need at least two for correct parsing
		if pos < 1 {
			continue
		}

		var (
			nextRune    rune
			prevRune    rune
			nextRuneLen int
			prevRuneLen int
		)

		runeLen := utf8.RuneLen(run)
		if runeLen == -1 {
			continue
		}

		if unicode.IsDigit(run) {
			if pars.checkDigitsForTimeStampAndChineese(
				inp, pos, runeLen, prevRuneLen, nextRuneLen) {
				return
			}
		}

		if pos-runeLen < 0 {
			continue
		}

		prevRune, prevRuneLen = utf8.DecodeRuneInString(inp[pos-runeLen : pos])
		nextRune, nextRuneLen = utf8.DecodeRuneInString(inp[pos+runeLen:])

		// colon-separated time
		if run == ':' && unicode.IsDigit(prevRune) && unicode.IsDigit(nextRune) {
			if pars.ParseColonSepTime(inp, run, pos, runeLen, prevRuneLen, nextRuneLen) {
				foundTimeHere = true
				break
			}
		}
	}

	if pars.foundTime && foundTimeHere {
		if pars.timePoints.Time.Start < len(inp) {
			leftPart := inp[:pars.timePoints.Time.Start]
			pars.parseDate(leftPart)
		}

		if pars.timePoints.Time.End >= 0 {
			rightPart := inp[pars.timePoints.Time.End:]
			pars.parseDate(rightPart)
		}
	} else {
		pars.parseDate(inp)
	}
}

func (pars *TimeParser) checkDigitsForTimeStampAndChineese(
	inp string, pos, runeLen, prevRuneLen, nextRuneLen int) bool {
	left, right := parseSeparatedDigits(inp, pos, runeLen, prevRuneLen, nextRuneLen)
	if left < 0 || right > len(inp) {
		return false
	}

	// unix timestamp
	if right-left >= 10 && len(inp) >= 10 {
		if digits, err := strconv.ParseInt(inp[left:right][:10], 10, 0); err == nil {
			tm := time.Unix(digits, 0)

			if valid := validateLongYear(tm.Year()); valid {
				pars.foundTime = true
				pars.foundDate = true
				pars.copyDateTimeValues(tm)

				return true
			}
		}
	}

	// 2021年06月29日09時20分
	if right < len(inp) {
		nextRune, nextRuneLen := utf8.DecodeRuneInString(inp[right:])
		if right-left == 4 && nextRune == '年' {
			idx := strings.Index(inp, string('分'))
			if idx != -1 {
				if tm, err := time.Parse(chineeseFormat, inp[left:idx+nextRuneLen]); err == nil {
					pars.foundTime = true
					pars.foundDate = true
					pars.copyDateTimeValues(tm)
					pars.copyTimePoints(left, idx+nextRuneLen)
					pars.storedTimeLayout = chineeseFormat
					pars.storedDateLayout = chineeseFormat

					return true
				}
			}
		}
	}

	return false
}

func (pars *TimeParser) ParseColonSepTime(
	inp string, run rune, pos, runeLen, prevRuneLen, nextRuneLen int) bool {
	left, right := parseSeparatedDigits(inp, pos, runeLen, prevRuneLen, nextRuneLen)

	// fix cases '10:55, 5 JAN 2021' - comma after time
	lastRune, lastRuneLen := utf8.DecodeLastRuneInString(inp[:right])
	if unicode.IsPunct(lastRune) {
		right -= lastRuneLen
	}

	// check ISO and RFC formats
	prevRune, _ := utf8.DecodeLastRuneInString(inp[:left])
	if prevRune == 'T' || prevRune == 'Z' {
		left, right := parseSpacesOrEnd(inp, pos, runeLen, prevRuneLen, nextRuneLen)

		if tm, layout, found := tryISOAndRFC(inp[left:right]); found {
			pars.foundTime = true
			pars.foundDate = true
			pars.copyDateTimeValues(tm)
			pars.copyTimePoints(left, right)
			pars.storedTimeLayout = layout
			pars.storedDateLayout = layout

			return true
		}
	}

	parsedTime := inp[left:right]
	layout := buildTimeLayout(parsedTime)

	if tm, err := time.Parse(layout, parsedTime); err == nil {
		pars.foundTime = true
		pars.copyTimeValues(tm)
		pars.copyTimePointsTime(left, right)
		pars.storedTimeLayout = layout

		return true
	}

	return false
}
