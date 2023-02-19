package model

type User struct {
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	SessionID string     `json:"sessionid"`
	Fav       []Keyboard `json:"fav"`
	BestKeys  []Keyboard `json:"bestkeys"`
	WorstKeys []Keyboard `json:"worstkeys"`
}
