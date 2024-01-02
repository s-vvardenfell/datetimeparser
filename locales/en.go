package locales

// Английский

var EnLocale = Locale{
	Name: "en",
	Months: map[string]int{
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
		"jan":       1,
		"feb":       2,
		"mar":       3,
		"apr":       4,
		"jun":       6,
		"jul":       7,
		"aug":       8,
		"sep":       9,
		"oct":       10,
		"nov":       11,
		"dec":       12,
	},
	Adverbs: map[string]int{
		"yesterday": -1,
		"today":     0,
		"tomorrow":  1,
	},
	Merediem: map[string]int{
		"am":   1,
		"pm":   2,
		"a.m.": 1,
		"p.m.": 2,
	},
}
