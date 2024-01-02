package locales

// Хинди

var HiLocale = Locale{
	Name: "hi",
	Months: map[string]int{
		"जनवरी":   1,
		"फरवरी":   2,
		"मार्च":   3,
		"अप्रैल":  4,
		"मई":      5,
		"जून":     6,
		"जुलाई":   7,
		"अगस्त":   8,
		"सितंबर":  9,
		"अक्टूबर": 10,
		"नवंबर":   11,
		"दिसंबर":  12,
	},
	Adverbs: map[string]int{
		"कल": -1,
		"आज": 0,
	},
	Merediem: map[string]int{
		"दिन": 1,
		"शाम": 2,
	},
}
