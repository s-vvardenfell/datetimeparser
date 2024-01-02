package locales

// Испанский

var EsLocale = Locale{
	Name: "es",
	Months: map[string]int{
		"enero":      1,
		"febrero":    2,
		"marzo":      3,
		"abril":      4,
		"mayo":       5,
		"junio":      6,
		"julio":      7,
		"agosto":     8,
		"septiembre": 9,
		"octubre":    10,
		"noviembre":  11,
		"diciembre":  12,
	},
	Adverbs: map[string]int{
		"anteayer": -2,
		"ayer":     -1,
		"hoy":      0,
		"mañana":   1,
	},
	Merediem: map[string]int{
		"días":   1,
		"noches": 2,
	},
}
