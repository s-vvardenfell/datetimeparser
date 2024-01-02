package locales

// Китайский

var ZhLocale = Locale{
	Name: "zh",
	Months: map[string]int{
		"一月":  1,
		"二月":  2,
		"三月":  3,
		"四月":  4,
		"五月":  5,
		"六月":  6,
		"七月":  7,
		"八月":  8,
		"九月":  9,
		"十月":  10,
		"十一月": 11,
		"十二月": 12,
	},
	Adverbs: map[string]int{
		"前天": -2,
		"昨天": -1,
		"今天": 0,
		"明天": 1,
		"后天": 2,
	},
	Merediem: map[string]int{
		"天":  1,
		"晚上": 2,
	},
}
