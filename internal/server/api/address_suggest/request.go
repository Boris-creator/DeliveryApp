package address_suggest

type suggestRequest struct {
	Query          string `json:"query"`
	HighestToponym string `json:"highestToponym"`
	LowestToponym  string `json:"lowestToponym"`
}
