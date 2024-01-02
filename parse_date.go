package datetimeparser

import (
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)

func (pars *TimeParser) parseDate(inp string) {
	if len(inp) < 3 {
		return
	}

	if pars.timePoints.Date.Stored {
		beg := pars.timePoints.Date.Start
		end := pars.timePoints.Date.End

		if beg < len(inp) && beg >= 0 &&
			end <= len(inp) && end >= 0 && (beg < end) { // here '<=' is correct for 'end'
			supposedDate := inp[beg:end]

			if tm, err := time.Parse(pars.storedDateLayout, supposedDate); err == nil {
				pars.foundDate = true
				pars.copyDateValues(tm)
			}
		}
	}

	// need to avoid splitting input if date was found in other call
	// and rewrite stored timePoints
	var foundDateHere bool

	for pos, run := range inp {
		if pars.foundDate { // if date was already found
			break
		}

		if pos < 1 {
			continue
		}

		runeLen := utf8.RuneLen(run)

		if pos-runeLen < 0 {
			continue
		}

		prevRune, prevRuneLen := utf8.DecodeRuneInString(inp[pos-runeLen : pos])
		nextRune, nextRuneLen := utf8.DecodeRuneInString(inp[pos+runeLen:])

		if !pars.foundDate && unicode.IsPunct(run) && unicode.IsDigit(prevRune) && unicode.IsDigit(nextRune) {
			left, right := parseSeparatedDigits(inp, pos, runeLen, prevRuneLen, nextRuneLen)
			parsedDate := inp[left:right]

			var (
				delim     string
				digits    [3]string
				digitsPos int
				symbStart int
			)

			for idx, symb := range parsedDate {
				if !unicode.IsPunct(symb) {
					if idx < len(parsedDate)-1 {
						continue
					}
				} else {
					delim = string(symb)
					if delim == ":" { // to avoid parsing -07:00 like a date
						break
					}
				}

				// to add last symbol if input ends
				if idx == len(parsedDate)-1 {
					idx++ // last digit in the end of the string
				}

				digits[digitsPos] = parsedDate[symbStart:idx]
				symbStart = idx + utf8.RuneLen(symb)
				digitsPos++
			}

			var (
				fstDigLen  = len(digits[0])
				secDigLen  = len(digits[1])
				thrDigLen  = len(digits[2])
				layoutElem string
			)

			// found 2 digits
			if fstDigLen != 0 && secDigLen != 0 && thrDigLen == 0 {
				var (
					fstDig, fstErr = strconv.Atoi(digits[0])
					secDig, secErr = strconv.Atoi(digits[1])
				)

				if fstErr != nil || secErr != nil {
					continue
				}

				// yyyy.mm
				if !pars.foundDate && fstDigLen == 4 && validateLongYear(fstDig) && validateMonth(secDig) {
					pars.parsedTime.Year = fstDig
					pars.parsedTime.Month = secDig
					pars.foundDate = true

					switch secDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout = "2006" + delim + layoutElem
				}

				// mm.yyyy
				if !pars.foundDate && secDigLen == 4 && validateMonth(fstDig) && validateLongYear(secDig) {
					pars.parsedTime.Year = secDig
					pars.parsedTime.Month = fstDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout = layoutElem + delim + "2006"
				}

				// dd.mm
				if !pars.foundDate && validateDay(fstDig) && validateMonth(secDig) {
					pars.parsedTime.Day = fstDig
					pars.parsedTime.Month = secDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "2"
					case 2:
						layoutElem = "02"
					}

					pars.storedDateLayout = layoutElem + delim

					switch secDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout += layoutElem
				}

				// mm.dd
				if !pars.foundDate && validateMonth(fstDig) && validateDay(secDig) {
					pars.parsedTime.Day = secDig
					pars.parsedTime.Month = fstDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout = layoutElem + delim

					switch secDigLen {
					case 1:
						layoutElem = "2"
					case 2:
						layoutElem = "02"
					}

					pars.storedDateLayout += layoutElem
				}

				// yy.mm
				if !pars.foundDate && validateShortYear(fstDig) && validateMonth(secDig) {
					pars.parsedTime.Year = fstDig
					pars.parsedTime.Month = secDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "6"
					case 2:
						layoutElem = "06"
					}

					pars.storedDateLayout = layoutElem + delim

					switch secDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout += layoutElem
				}

				// mm.yy
				if !pars.foundDate && validateMonth(fstDig) && validateShortYear(secDig) {
					pars.parsedTime.Year = secDig
					pars.parsedTime.Month = fstDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout = layoutElem + delim

					switch secDigLen {
					case 1:
						layoutElem = "6"
					case 2:
						layoutElem = "06"
					}

					pars.storedDateLayout += layoutElem
				}
			}

			// found 3 digits
			if fstDigLen != 0 && secDigLen != 0 && thrDigLen != 0 {
				var (
					fstDig, fstErr = strconv.Atoi(digits[0])
					secDig, secErr = strconv.Atoi(digits[1])
					thrDig, thrErr = strconv.Atoi(digits[2])
				)

				if fstErr != nil || secErr != nil || thrErr != nil {
					continue
				}

				// dd.mm.yyyy
				if !pars.foundDate && thrDigLen == 4 &&
					validateDay(fstDig) && validateMonth(secDig) && validateLongYear(thrDig) {
					pars.parsedTime.Year = thrDig
					pars.parsedTime.Month = secDig
					pars.parsedTime.Day = fstDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "2"
					case 2:
						layoutElem = "02"
					}

					pars.storedDateLayout = layoutElem + delim

					switch secDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout += layoutElem + delim + "2006"
				}

				// yyyy.mm.dd
				if !pars.foundDate && fstDigLen == 4 &&
					validateLongYear(fstDig) && validateMonth(secDig) && validateDay(thrDig) {
					pars.parsedTime.Year = fstDig
					pars.parsedTime.Month = secDig
					pars.parsedTime.Day = thrDig
					pars.foundDate = true

					pars.storedDateLayout = "2006" + delim

					switch secDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout += layoutElem + delim

					switch thrDigLen {
					case 1:
						layoutElem = "2"
					case 2:
						layoutElem = "02"
					}

					pars.storedDateLayout += layoutElem
				}

				// dd.mm.yy
				if !pars.foundDate &&
					validateDay(fstDig) && validateMonth(secDig) && validateShortYear(thrDig) {
					pars.parsedTime.Year = thrDig
					pars.parsedTime.Month = secDig
					pars.parsedTime.Day = fstDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "2"
					case 2:
						layoutElem = "02"
					}

					pars.storedDateLayout = layoutElem + delim

					switch secDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout += layoutElem + delim

					switch thrDigLen {
					case 1:
						layoutElem = "6"
					case 2:
						layoutElem = "06"
					}

					pars.storedDateLayout += layoutElem
				}

				// yy.mm.dd
				if !pars.foundDate &&
					validateShortYear(fstDig) && validateMonth(secDig) && validateDay(thrDig) {
					pars.parsedTime.Year = fstDig
					pars.parsedTime.Month = secDig
					pars.parsedTime.Day = thrDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "6"
					case 2:
						layoutElem = "06"
					}

					pars.storedDateLayout = layoutElem + delim

					switch secDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout += layoutElem + delim

					switch thrDigLen {
					case 1:
						layoutElem = "2"
					case 2:
						layoutElem = "02"
					}

					pars.storedDateLayout += layoutElem
				}

				// yy.mm.dd
				if !pars.foundDate &&
					validateShortYear(fstDig) && validateDay(secDig) && validateMonth(thrDig) {
					pars.parsedTime.Year = fstDig
					pars.parsedTime.Month = thrDig
					pars.parsedTime.Day = secDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "6"
					case 2:
						layoutElem = "06"
					}

					pars.storedDateLayout = layoutElem + delim

					switch secDigLen {
					case 1:
						layoutElem = "2"
					case 2:
						layoutElem = "02"
					}

					pars.storedDateLayout += layoutElem + delim

					switch thrDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout += layoutElem
				}

				// yyyy.mm.dd
				if !pars.foundDate &&
					validateLongYear(fstDig) && validateDay(secDig) && validateMonth(thrDig) {
					pars.parsedTime.Year = fstDig
					pars.parsedTime.Month = thrDig
					pars.parsedTime.Day = secDig
					pars.foundDate = true

					pars.storedDateLayout = "2006" + delim

					switch secDigLen {
					case 1:
						layoutElem = "2"
					case 2:
						layoutElem = "02"
					}

					pars.storedDateLayout += layoutElem + delim

					switch thrDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout += layoutElem
				}

				// mm.dd.yyyy
				if !pars.foundDate && thrDigLen == 4 &&
					validateMonth(fstDig) && validateDay(secDig) && validateLongYear(thrDig) {
					pars.parsedTime.Year = thrDig
					pars.parsedTime.Month = fstDig
					pars.parsedTime.Day = secDig
					pars.foundDate = true

					switch fstDigLen {
					case 1:
						layoutElem = "1"
					case 2:
						layoutElem = "01"
					}

					pars.storedDateLayout = layoutElem + delim

					switch secDigLen {
					case 1:
						layoutElem = "2"
					case 2:
						layoutElem = "02"
					}

					pars.storedDateLayout += layoutElem + delim + "2006"
				}

				if pars.foundDate {
					pars.timePoints.Date.Start = left
					pars.timePoints.Date.End = right
					pars.timePoints.Date.Stored = true
					foundDateHere = true

					break
				}
			}
		}
	}

	if pars.foundDate && foundDateHere {
		leftPart := inp[:pars.timePoints.Date.Start]
		rightPart := inp[pars.timePoints.Date.End:]

		pars.parseTimeDetails(leftPart)
		pars.parseTimeDetails(rightPart)
	} else {
		pars.parseTimeDetails(inp)
	}
}
