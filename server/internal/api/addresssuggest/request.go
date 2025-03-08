package addresssuggest

type suggestRequest struct {
	Query          string `json:"query"`
	HighestToponym string `json:"highestToponym"`
	LowestToponym  string `json:"lowestToponym"`
}
