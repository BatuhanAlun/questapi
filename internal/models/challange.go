package models

type Challange struct {
	Name      string     `json:"challange-name"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Name    string   `json:"quest-name"`
	Options []string `json:"options"`
	Answer  rune     `json:"answer"`
	Score   int      `json:"score"`
}
