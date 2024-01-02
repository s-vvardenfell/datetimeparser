package locales

// Итальянский

var ItLocale = Locale{
	Name: "it",
	Months: map[string]int{
		"gennaio":   1,
		"febbraio":  2,
		"marzo":     3,
		"aprile":    4,
		"maggio":    5,
		"giugno":    6,
		"luglio":    7,
		"agosto":    8,
		"settembre": 9,
		"ottobre":   10,
		"novembre":  11,
		"dicembre":  12,
	},
	Adverbs: map[string]int{
		"ieri":   -1,
		"oggi":   0,
		"domani": 1,
	},
	Merediem: map[string]int{
		"giorni": 1,
		"serate": 2,
	},
}
