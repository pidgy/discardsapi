package api

// Card contains various metadata pertaining to a single TCG card.
type Card struct {
	Data `json:"data"`
}

// Data contains the entire collection of metadata for a Card.
type Data struct {
	ID                     string
	Name                   string
	Supertype              string
	Subtypes               []string
	HP                     string
	Types                  []string
	EvolvesFrom            string
	Rules                  []string
	Attacks                []Attack
	Weaknesses             []Weakness
	RetreatCost            []string
	ConvertedRetreatCost   int
	Set                    Set
	Number                 string
	Artist                 string
	Rarity                 string
	NationalPokedexNumbers []int
	Legalities             map[string]string
	RegulationMark         string
	Images                 map[string]string
	CardMarket             CardMarket
}

// Attack represents a specific field within any Card Data field.
type Attack struct {
	Name                string
	Cost                []string
	ConvertedEnergyCost int
	Damage              string
	Text                string
}

// Weakness represents a specific field within any Card Data field.
type Weakness struct {
	Type  string
	Value string
}

// Set represents TCG set information for a Card.
type Set struct {
	ID           string
	Name         string
	Series       string
	PrintedTotal int
	Total        int
	Legalities   map[string]string
	PtcgoCode    string
	ReleaseDate  string
	UpdatedAt    string
	Images       map[string]string
}

// CardMarket represents market value prices and metadata for a Card.
type CardMarket struct {
	URL       string
	UpdatedAt string
	Prices    Prices
}

// Prices represents the specific market values for a CardMarket.
type Prices struct {
	AverageSellPrice float64
	LowPrice         float64
	TrendPrice       float64
	ReverseHoloTrend float64
	Avg1             float64
	Avg7             float64
	Avg30            float64
	ReverseHoloAvg1  float64
	ReverseHoloAvg7  float64
	ReverseHoloAvg30 float64
}
