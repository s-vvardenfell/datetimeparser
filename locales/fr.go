package locales

// Французский

var FrLocale = Locale{
	Name: "fr",
	Months: map[string]int{
		"janvier":   1,
		"février":   2,
		"mars":      3,
		"avril":     4,
		"mai":       5,
		"juin":      6,
		"juillet":   7,
		"août":      8,
		"septembre": 9,
		"octobre":   10,
		"novembre":  11,
		"décembre":  12,
	},
	Adverbs: map[string]int{
		"hier":   -1,
		"demain": 1,
	},
	Merediem: map[string]int{
		"journée": 1,
		"soirées": 2,
	},
}
