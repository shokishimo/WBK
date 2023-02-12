package model

type Keyboard struct {
	Name     string   `json:"name"`
	Make     string   `json:"make"`
	Price    string   `json:"price"`
	Layout   string   `json:"layout"`
	Profile  string   `json:"profile"`
	Switch   string   `json:"switch"`
	Material Material `json:"material"`
	Ranking  string   `json:"ranking"`
	Score    string   `json:"score"`
	Url      string   `json:"url"`
}
