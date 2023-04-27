package models

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Email    string `json:"UserEmail"`
	Password string `json:"UserPassword"`
}
