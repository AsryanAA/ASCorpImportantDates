package models

type User struct {
	Login   string `json:"login"`
	Surname string `json:"surname"`
	Name    string `json:"name"`
	DOB     string `json:"dob"`
	RegDate string `json:"reg_date"`
}
