package model

type Keyboard struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Make      string   `json:"make"`
	Price     string   `json:"price"`
	Layout    int      `json:"layout"`
	NumOfKeys int      `json:"numOfKeys"`
	Profile   string   `json:"profile"`
	Switch    string   `json:"switch"`
	Material  Material `json:"material"`
	Ranking   int      `json:"ranking"`
}
