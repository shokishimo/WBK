package model

type Keyboard struct {
	Name    string `json:"name"`
	Make    string `json:"make"`
	Price   string `json:"price"`
	Layout  string `json:"layout"`
	Profile string `json:"profile"`
	Switch  string `json:"switch"`
	Frame   string `json:"frame"`
	Keycap  string `json:"keycap"`
	Ranking string `json:"ranking"`
	Score   string `json:"score"`
	Url     string `json:"url"`
	Other   string `json:"other"`
}
