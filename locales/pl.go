package locales

// Польский

var PlLocale = Locale{
	Name: "pl",
	Months: map[string]int{
		"styczeń":     1,
		"luty":        2,
		"marzec":      3,
		"kwiecień":    4,
		"maj":         5,
		"czerwiec":    6,
		"lipiec":      7,
		"sierpień":    8,
		"wrzesień":    9,
		"październik": 10,
		"listopad":    11,
		"grudzień":    12,
	},
	Adverbs: map[string]int{
		"przedwczoraj": -2,
		"wczoraj":      -1,
		"dzisiaj":      0,
		"jutro":        1,
		"pojutrze":     2,
	},
	Merediem: map[string]int{
		"dnia":     1,
		"wieczory": 2,
	},
}
