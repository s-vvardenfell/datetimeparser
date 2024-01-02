package datetimeparser

import (
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func (pars *TimeParser) parseTimeDetails(inp string) {
	if len(inp) < 3 {
		return
	}

	for pos, run := range inp {
		// +- 0700 07:00
		if !pars.foundTimeDetails && (run == '+' || run == '-') {
			runeLen := utf8.RuneLen(run)

			if pos-runeLen < 0 {
				continue
			}

			_, prevRuneLen := utf8.DecodeRuneInString(inp[pos-runeLen : pos])
			_, nextRuneLen := utf8.DecodeRuneInString(inp[pos+runeLen:])

			left, right := parseSeparatedDigits(inp, pos+runeLen, runeLen, prevRuneLen, nextRuneLen)

			if right > len(inp) {
				right = len(inp)
			}

			for i := range tzOffsetLayouts {
				if tm, err := time.Parse(tzOffsetLayouts[i], inp[left:right]); err == nil {
					pars.parsedTime.TZ = tm.Location()
					pars.foundTimeDetails = true

					break
				}
			}
		}

		if !unicode.IsLetter(run) { // words like 'yesterday', 'januar' or 'MST'-like TZ, PM, p.m. etc
			continue
		}

		var (
			runeLen      = utf8.RuneLen(run)
			tempPosRight = pos + runeLen
			nextRune     rune
			nextRuneLen  int
			wordBeg      int
			wordEnd      int
			word         string
		)

		// found entire word
		for tempPosRight+nextRuneLen <= len(inp) {
			nextRune, nextRuneLen = utf8.DecodeRuneInString(inp[tempPosRight:])
			if !unicode.IsLetter(nextRune) {
				break
			}

			tempPosRight += nextRuneLen
		}

		wordBeg = pos
		wordEnd = tempPosRight
		word = inp[wordBeg:wordEnd]

		// 'four hours ago' / '31 minutes ago'
		if word == "ago" {
			pars.ParseTimeAgoFormat(inp, run, pos, runeLen)
		}

		var (
			loc           string
			value         int
			foundMonth    bool
			foundAdverb   bool
			foundMerediem bool
		)

		// PM or pm
		if !pars.foundMerediem && (word == "PM" || word == "pm") {
			pars.parsedTime.Meridiem = 2
			pars.foundMerediem = true
			pars.foundedLocale = "en"

			continue
		}

		// P.M. or p.m.
		if !pars.foundMerediem && (run == 'p' || run == 'P' && nextRune == '.') {
			posRight := tempPosRight

			if posRight+nextRuneLen > len(inp) {
				continue
			}

			nextR1, _ := utf8.DecodeRuneInString(inp[posRight+nextRuneLen:])

			if nextR1 == 'm' || nextR1 == 'M' {
				pars.parsedTime.Meridiem = 2
				pars.foundMerediem = true
				pars.foundTimeDetails = true
				pars.foundedLocale = "en"

				continue
			}
		}

		// MST-like
		if wordLen := len(word); (wordLen >= 3 && wordLen <= 5) && pars.parsedTime.NamedTZVal == 0 {
			if val, ok := namedTimeZones[word]; ok {
				pars.parsedTime.NamedTZVal = val
				pars.foundTimeDetails = true

				continue
			}
		}

		// for months, adverbs and merediem in non-eng languages
		word = strings.ToLower(word)

		if !pars.foundTimeDetails && !pars.foundMerediem {
			value, loc, foundMerediem = searchMerediemByName(word, pars.foundedLocale)
			if foundMerediem {
				pars.parsedTime.Meridiem = value
				pars.foundedLocale = loc
				pars.foundMerediem = true
				pars.foundTimeDetails = true

				continue
			}
		}

		if !pars.foundAdverb {
			value, loc, foundAdverb = searchAdverbsByName(word, pars.foundedLocale)
			if foundAdverb {
				pars.parsedTime.Adverb = value
				pars.foundedLocale = loc
				pars.foundTimeDetails = true
				pars.foundAdverb = true
				pars.foundDate = true

				return
			}
		}

		if !pars.foundDate && pars.parsedTime.Month == 0 {
			value, loc, foundMonth = searchMonthsByName(word, pars.foundedLocale)
			if foundMonth {
				pars.parsedTime.Month = value
				pars.foundedLocale = loc

				pars.parseNearbyDigits(inp[:wordBeg])
				pars.parseNearbyDigits(inp[wordEnd:])

				return
			}
		} // inp = inp[wordEnd:] // don't work properly for all test cases
	}
}

func (pars *TimeParser) ParseTimeAgoFormat(inp string, run rune, pos, runeLen int) {
	// search prev word
	tempPosLeft := pos - runeLen

	_, prevRuneLen := utf8.DecodeRuneInString(inp[tempPosLeft:pos])
	for tempPosLeft-prevRuneLen >= 0 {
		prevRune, prevRuneLen := utf8.DecodeRuneInString(inp[tempPosLeft-prevRuneLen : tempPosLeft])
		if !unicode.IsLetter(prevRune) {
			break
		}

		tempPosLeft -= prevRuneLen
	}

	shiftLeft := tempPosLeft
	wordBeg := shiftLeft
	wordEnd := pos - prevRuneLen
	word := inp[wordBeg:wordEnd]
	tempPosLeft -= prevRuneLen

	switch word {
	case "minute", "minutes":
		_, prevRuneLen := utf8.DecodeRuneInString(inp[tempPosLeft:pos])

		for tempPosLeft-prevRuneLen >= 0 {
			prevRune, prevRuneLen := utf8.DecodeRuneInString(inp[tempPosLeft-prevRuneLen : tempPosLeft])
			if !unicode.IsDigit(prevRune) {
				break
			}

			tempPosLeft -= prevRuneLen
		}

		shiftLeft := tempPosLeft
		wordEnd = wordBeg // second word end its third word beg-runeLen
		word := inp[shiftLeft : wordEnd-prevRuneLen]

		if num, err := strconv.Atoi(word); err == nil {
			tm := time.Now().Add(time.Minute * -time.Duration(num))

			pars.foundDate = true
			pars.foundTime = true
			pars.copyDateTimeValues(tm)

			return
		}
	case "hour", "hours":
		_, prevRuneLen := utf8.DecodeRuneInString(inp[tempPosLeft:pos])

		for tempPosLeft-prevRuneLen >= 0 {
			prevRune, prevRuneLen := utf8.DecodeRuneInString(inp[tempPosLeft-prevRuneLen : tempPosLeft])
			if !unicode.IsLetter(prevRune) {
				break
			}

			tempPosLeft -= prevRuneLen
		}

		shiftLeft := tempPosLeft
		wordEnd = wordBeg // second word end its third word beg-runeLen
		word := inp[shiftLeft : wordEnd-prevRuneLen]

		if num, ok := numbersInWords[word]; ok {
			tm := time.Now().Add(time.Hour * -time.Duration(num))

			pars.foundDate = true
			pars.foundTime = true
			pars.copyDateTimeValues(tm)

			return
		}
	}
}
