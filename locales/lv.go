package locales

// Латышский

var LvLocale = Locale{
	Name: "lv",
	Months: map[string]int{
		"janvāris":   1,
		"februāris":  2,
		"marts":      3,
		"aprīlis":    4,
		"maijs":      5,
		"jūnijs":     6,
		"jūlijs":     7,
		"augusts":    8,
		"septembris": 9,
		"oktobris":   10,
		"novembris":  11,
		"decembris":  12,
	},
	Adverbs: map[string]int{
		"aizvakar": -2,
		"vakar":    -1,
		"šodien":   0,
		"rīt":      1,
	},
	Merediem: map[string]int{
		"dienas": 1,
		"vakari": 2,
	},
}
