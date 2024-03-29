package locales

// Иврит

var HeLocale = Locale{
	Name: "he",
	Months: map[string]int{
		"ינואר":   1,
		"פברואר":  2,
		"מרץ":     3,
		"אפריל":   4,
		"מאי":     5,
		"יוני":    6,
		"יולי":    7,
		"אוגוסט":  8,
		"ספטמבר":  9,
		"אוקטובר": 10,
		"נובמבר":  11,
		"דצמבר":   12,
	},
	Adverbs: map[string]int{
		"שלשום":   -2,
		"אתמול":   -1,
		"היום":    0,
		"מחר":     1,
		"מחרתיים": 2,
	},
	Merediem: map[string]int{
		"היום":  1,
		"ערבים": 2,
	},
}
