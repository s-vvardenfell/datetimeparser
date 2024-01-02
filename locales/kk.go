package locales

// Казахский

var KkLocale = Locale{
	Name: "kk",
	Months: map[string]int{
		"қаңтар":    1,
		"ақпан":     2,
		"наурыз":    3,
		"сәуір":     4,
		"мамыр":     5,
		"маусым":    6,
		"шілде":     7,
		"тамыз":     8,
		"қыркүйек":  9,
		"қазан":     10,
		"қараша":    11,
		"желтоқсан": 12,
	},
	Adverbs: map[string]int{
		"кеше":  -2,
		"бүгін": 0,
		"ертең": 1,
	},
	Merediem: map[string]int{
		"күн":    1,
		"кештер": 2,
	},
}