package locales

// Нидерландский

var NlLocale = Locale{
	Name: "nl",
	Months: map[string]int{
		"januari":   1,
		"februari":  2,
		"maart":     3,
		"april":     4,
		"mei":       5,
		"juni":      6,
		"juli":      7,
		"augustus":  8,
		"september": 9,
		"oktober":   10,
		"november":  11,
		"december":  12,
	},
	Adverbs: map[string]int{
		"eergisteren": -2,
		"gisteren":    -1,
		"vandaag":     0,
		"morgen":      1,
		"overmorgen":  2,
	},
	Merediem: map[string]int{
		"dagen":   1,
		"avonden": 2,
	},
}
