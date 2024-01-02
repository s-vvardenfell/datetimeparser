package datetimeparser

import (
	"unicode"
	"unicode/utf8"
)

func validateDay(num int) bool {
	if num < 0 || num > 31 {
		return false
	}

	return true
}

func validateMonth(num int) bool {
	if num < 0 || num > 12 {
		return false
	}

	return true
}

func validateLongYear(num int) bool {
	if num < startYearLong || num > endYearLong {
		return false
	}

	return true
}

func validateShortYear(num int) bool {
	return num <= endYearShort
}

func buildTimeLayout(parsedTime string) string {
	var (
		layout      string
		symbLen     int
		prepHours   bool
		prepMinuts  bool
		prepSeconds bool
	)

	for i, symb := range parsedTime {
		if !unicode.IsPunct(symb) {
			symbLen++

			if i < len(parsedTime)-1 {
				continue
			}
		}

		if !prepHours {
			if symbLen == 1 {
				layout += "3:"
			}

			if symbLen == 2 {
				layout += "15:"
			}

			symbLen = 0
			prepHours = true

			continue
		}

		if !prepMinuts {
			if symbLen == 1 {
				layout += "4"
			}

			if symbLen == 2 {
				layout += "04"
			}

			symbLen = 0
			prepMinuts = true

			continue
		}

		if !prepSeconds {
			if symbLen == 1 {
				layout += ":5"
			}

			if symbLen == 2 {
				layout += ":05"
			}

			symbLen = 0
			prepSeconds = true
		}
	}

	return layout
}

func parseSeparatedDigits(inp string, pos int, runeLen, prevRuneLen, nextRuneLen int) (int, int) {
	var (
		tempPosLeft  int = pos - runeLen
		tempPosRight int = pos + runeLen
		nextRune     rune
		prevRune     rune
	)

	// looking for digits to the left
	for tempPosLeft-prevRuneLen >= 0 {
		prevRune, prevRuneLen = utf8.DecodeRuneInString(inp[tempPosLeft-prevRuneLen : tempPosLeft])
		if !unicode.IsDigit(prevRune) {
			break
		}

		tempPosLeft -= prevRuneLen
	}

	shiftLeft := tempPosLeft

	// looking for digits to the right
	for tempPosRight+nextRuneLen < len(inp) {
		nextRune, nextRuneLen = utf8.DecodeRuneInString(inp[tempPosRight+nextRuneLen:])
		if !unicode.IsDigit(nextRune) {
			break
		}

		tempPosRight += nextRuneLen
	}

	// jump over second separator if exists
	if unicode.IsPunct(nextRune) {
		tempPosRight += nextRuneLen
	}

	for tempPosRight+nextRuneLen < len(inp) {
		nextRune, nextRuneLen = utf8.DecodeRuneInString(inp[tempPosRight+nextRuneLen:])
		if !unicode.IsDigit(nextRune) {
			break
		}

		tempPosRight += nextRuneLen
	}

	shiftRight := tempPosRight + nextRuneLen

	return shiftLeft, shiftRight
}

func parseSpacesOrEnd(inp string, pos int, runeLen, prevRuneLen, nextRuneLen int) (int, int) {
	var (
		tempPosLeft  int = pos - runeLen
		tempPosRight int = pos + runeLen
		nextRune     rune
		prevRune     rune
		storedRune   rune
	)

	for tempPosLeft-prevRuneLen >= 0 {
		prevRune, prevRuneLen = utf8.DecodeRuneInString(inp[tempPosLeft-prevRuneLen : tempPosLeft])
		if unicode.IsSpace(prevRune) {
			break
		}

		tempPosLeft -= prevRuneLen
	}

	shiftLeft := tempPosLeft

	for tempPosRight+nextRuneLen < len(inp) {
		nextRune, nextRuneLen = utf8.DecodeRuneInString(inp[tempPosRight+nextRuneLen:])
		if unicode.IsSpace(nextRune) {
			break
		}

		tempPosRight += nextRuneLen
		storedRune = nextRune
	}

	shiftRight := tempPosRight

	return shiftLeft, shiftRight + utf8.RuneLen(storedRune)
}
