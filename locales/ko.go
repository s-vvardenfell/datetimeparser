package locales

// Корейский

var KoLocale = Locale{
	Name: "ko",
	Months: map[string]int{
		"1 월":  1,
		"2 월":  2,
		"3 월":  3,
		"4 월":  4,
		"5 월":  5,
		"6 월":  6,
		"7 월":  7,
		"8 월":  8,
		"9 월":  9,
		"10 월": 10,
		"11 월": 11,
		"12 월": 12,
	},
	Adverbs: map[string]int{
		"어제": -1,
		"오늘": 0,
		"내일": 1,
	},
	Merediem: map[string]int{
		"일":  1,
		"저녁": 2,
	},
}
