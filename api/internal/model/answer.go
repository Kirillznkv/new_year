package model

type Answer struct {
	ID     int    `json:"id"`
	Lvl    int    `json:"lvl"`
	UserId int    `json:"user_id"`
	Image  string `json:"image"`
}
