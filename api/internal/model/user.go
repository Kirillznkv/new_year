package model

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Text       string `json:"text"`
}
