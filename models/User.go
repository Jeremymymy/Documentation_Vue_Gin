package models

import (
	"documentation/dbconnect"
	"log"
)

type User struct {
	Id       int    `json:"UserId" binding:"required"`
	Name     string `json:"UserName" binding:"gte=2"`
	Email    string `json:"UserEmail" binding:"required,email"`
	Password string `json:"UserPassword" binding:"required,min=4,max=20"`
}

func FindAllUsers() []User {
	var users []User
	dbconnect.MySQLcon.Find(&users)
	return users
}

func FindByUserId(userId string) User {
	var user User
	dbconnect.MySQLcon.Where("id = ?", userId).First(&user)
	return user
}

func CreateUser(user User) User {
	dbconnect.MySQLcon.Create(&user)
	return user
}

func DeleteUser(userId string) bool {
	var user User
	result := dbconnect.MySQLcon.Where("id = ?", userId).Delete(&user)
	log.Println(result)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func UpdateUser(userId string, user User) User {
	dbconnect.MySQLcon.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}
