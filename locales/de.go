package locales

// Немецкий

var DeLocale = Locale{
	Name: "de",
	Months: map[string]int{
		"januar":    1,
		"februar":   2,
		"märz":      3,
		"april":     4,
		"mai":       5,
		"juni":      6,
		"juli":      7,
		"august":    8,
		"september": 9,
		"oktober":   10,
		"november":  11,
		"dezember":  12,
		"jan":       1,
		"feb":       2,
		"mär":       3,
		"apr":       4,
		"jun":       6,
		"jul":       7,
		"aug":       8,
		"sep":       9,
		"okt":       10,
		"nov":       11,
		"dez":       12,
	},
	Adverbs: map[string]int{
		"vorgestern": -2,
		"gestern":    -1,
		"heute":      0,
		"morgen":     1,
		"übermorgen": 2,
	},
	Merediem: map[string]int{
		"nachts":      1,
		"morgens":     1,
		"nachmittags": 2,
		"abends":      2,
	},
}
