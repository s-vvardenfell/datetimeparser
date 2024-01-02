package locales

// Португальский

var PtLocale = Locale{
	Name: "pt",
	Months: map[string]int{
		"janeiro":   1,
		"fevereiro": 2,
		"março":     3,
		"abril":     4,
		"maio":      5,
		"junho":     6,
		"julho":     7,
		"agosto":    8,
		"setembro":  9,
		"outubro":   10,
		"novembro":  11,
		"dezembro":  12,
	},
	Adverbs: map[string]int{
		"anteontem": -2,
		"ontem":     -1,
		"hoje":      0,
		"amanhã":    1,
	},
	Merediem: map[string]int{
		"dia":    1,
		"noites": 2,
	},
}
