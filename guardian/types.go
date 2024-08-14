package guardian

type GuardianInput struct {
	Text           string `json:"text"`
	APIKey         string
	ScoreThreshold float64 `json:"score_threshold"`
}

type GuardianOutput struct {
	Flagged        bool                   `json:"flagged"`
	Categories     GuardianCategories     `json:"categories"`
	CategoryScores GuardianCategoryScores `json:"category_scores"`
}

type GuardianCategories struct {
	Defamation bool `json:"defamation"`
	Hate       bool `json:"hate"`
	SelfHarm   bool `json:"self_harm"`
	Sexual     bool `json:"sexual"`
	Violence   bool `json:"violence"`
}
type GuardianCategoryScores struct {
	Defamation float64 `json:"defamation"`
	Hate       float64 `json:"hate"`
	SelfHarm   float64 `json:"self_harm"`
	Sexual     float64 `json:"sexual"`
	Violence   float64 `json:"violence"`
}
