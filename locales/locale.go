package locales

type Locales []Locale

type Locale struct {
	Name     string
	Months   map[string]int
	Adverbs  map[string]int
	Merediem map[string]int
}
