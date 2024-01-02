package locales

// Финский

var FiLocale = Locale{
	Name: "fi",
	Months: map[string]int{
		"tammikuu":   1,
		"helmikuuta": 2,
		"maaliskuu":  3,
		"huhtikuu":   4,
		"toukokuu":   5,
		"kesäkuu":    6,
		"heinäkuu":   7,
		"elokuu":     8,
		"syyskuu":    9,
		"lokakuu":    10,
		"marraskuu":  11,
		"joulukuu":   12,
	},
	Adverbs: map[string]int{
		"toissapäivänä": -2,
		"eilen":         -1,
		"tänään":        0,
		"huomenna":      1,
	},
	Merediem: map[string]int{
		"päivät": 1,
		"illat":  2,
	},
}
