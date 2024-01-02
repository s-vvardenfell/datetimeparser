package locales

// Турецкий

var TrLocale = Locale{
	Name: "tr",
	Months: map[string]int{
		"ocak":    1,
		"şubat":   2,
		"mart":    3,
		"nisan":   4,
		"mayıs":   5,
		"haziran": 6,
		"temmuz":  7,
		"ağustos": 8,
		"eylül":   9,
		"ekim":    10,
		"kasım":   11,
		"aralık":  12,
	},
	Adverbs: map[string]int{
		"dün":   -1,
		"bugün": 0,
		"yarın": 1,
	},
	Merediem: map[string]int{
		"gün":      1,
		"akşamlar": 2,
	},
}
