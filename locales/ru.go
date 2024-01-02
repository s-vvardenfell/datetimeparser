package locales

// Русский

var RuLocale = Locale{
	Name: "ru",
	Months: map[string]int{
		"январ":   1,
		"феврал":  2,
		"март":    3,
		"апрел":   4,
		"май":     5,
		"мае":     5,
		"мая":     5,
		"июн":     6,
		"июл":     7,
		"август":  8,
		"сентябр": 9,
		"октябр":  10,
		"ноябр":   11,
		"декабр":  12,
	},
	Adverbs: map[string]int{
		"позавчера":   -2,
		"вчера":       -1,
		"сегодня":     0,
		"завтра":      1,
		"послезавтра": 2,
	},
	Merediem: map[string]int{
		"дня":    1,
		"вечера": 2,
	},
}
