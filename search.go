package datetimeparser

import (
	"strings"
	"time"
)

func searchMonthsByName(input, localeName string) (int, string, bool) {
	if localeName != "" {
		if foundedLocale, ok := loc[localeName]; ok {
			val, found := searchInLocaleMap(input, foundedLocale.Months)
			if found {
				return val, localeName, true
			}
		}
	}

	for i := range loc {
		val, found := searchInLocaleMap(input, loc[i].Months)
		if found {
			return val, loc[i].Name, true
		}
	}

	return 0, "", false
}

func searchMerediemByName(input, localeName string) (int, string, bool) {
	if localeName != "" {
		foundedLocale, ok := loc[localeName]
		if ok {
			val, found := searchInLocaleMap(input, foundedLocale.Merediem)
			if found {
				return val, localeName, true
			}
		}
	}

	for i := range loc {
		val, found := searchInLocaleMap(input, loc[i].Merediem)
		if found {
			return val, loc[i].Name, true
		}
	}

	return 0, "", false
}

func searchAdverbsByName(input, localeName string) (int, string, bool) {
	if localeName != "" {
		foundedLocale, ok := loc[localeName]
		if ok {
			val, found := searchInLocaleMap(input, foundedLocale.Adverbs)
			if found {
				return val, localeName, true
			}
		}
	}

	for i := range loc {
		val, found := searchInLocaleMap(input, loc[i].Adverbs)
		if found {
			return val, loc[i].Name, true
		}
	}

	return 0, "", false
}

func searchInLocaleMap(input string, where map[string]int) (int, bool) {
	for key := range where {
		// need Contains, because 'вчерашний' also gives 'вчера', 'понедельника' gives 'понедельник'
		if strings.Contains(input, key) {
			return where[key], true
		}
	}

	return 0, false
}

func tryISOAndRFC(inp string) (time.Time, string, bool) {
	for i := range continuousFormats {
		tm, err := time.Parse(continuousFormats[i], inp)
		if err == nil {
			return tm, continuousFormats[i], true
		}
	}

	return time.Time{}, "", false
}
