package locales

// Фарси (Персидский)

var FaLocale = Locale{
	Name: "fa",
	Months: map[string]int{
		"ژانویه":  1,
		"فوریه":   2,
		"مارس":    3,
		"فروردین": 4,
		"مه":      5,
		"خرداد":   6,
		"جولای":   7,
		"اوت":     8,
		"سپتامبر": 9,
		"اکتبر":   10,
		"نوامبر":  11,
		"دسامبر":  12,
	},
	Adverbs: map[string]int{
		"دیروز": -1,
		"امروز": 0,
		"فردا":  1,
	},
	Merediem: map[string]int{
		"روزها": 1,
		"عصرها": 2,
	},
}
