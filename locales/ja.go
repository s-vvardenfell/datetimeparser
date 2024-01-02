package locales

// Японский

var JaLocale = Locale{
	Name: "ja",
	Months: map[string]int{
		"一月":  1,
		"二月":  2,
		"行进":  3,
		"四月":  4,
		"可能":  5,
		"六月":  6,
		"七月":  7,
		"八月":  8,
		"九月":  9,
		"十月":  10,
		"十一月": 11,
		"十二月": 12,
	},
	Adverbs: map[string]int{
		"一昨日": -2,
		"昨日":  -1,
		"今日":  0,
		"明日":  1,
		"明後日": 2,
	},
	Merediem: map[string]int{
		"日": 1,
		"夜": 2,
	},
}
