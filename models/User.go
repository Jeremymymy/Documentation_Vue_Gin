package models

import (
	"documentation/dbconnect"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	EmployeeId string     `json:"EmployeeId" binding:"required" gorm:"primaryKey;index;unique"`
	Name       string     `json:"Name" binding:"gte=2"`
	Email      string     `json:"Email" binding:"required,email" gorm:"unique"`
	Password   string     `json:"Password" binding:"required,min=4"`
	Department string     `json:"Department" binding:"required,oneof=HR Sales Marketing IT Finance"`
	PostDocs   []Document `json:"PostDocs" gorm:"foreignKey:AuthorId"`
	// EditDocs   []Document `json:"EditDocs" gorm:"foreignKey:AuthorId"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func FindAllUsers() []User {
	var users []User
	dbconnect.MySQLcon.Find(&users)
	return users
}

func FindByUserId(id string) User {
	var user User
	dbconnect.MySQLcon.Where("employee_id = ?", id).First(&user)
	return user
}

func FindByUserIdWithPostDocsPreload(id string) (User, error) {
	var user User
	err := dbconnect.MySQLcon.Preload("PostDocs").First(&user, id).Error
	return user, err
}

func FindByUserEmail(email string) User {
	var user User
	dbconnect.MySQLcon.Where("email = ?", email).First(&user)
	return user
}

// func GetMyDetail(id string) (User, error) {

// }

func AddNewDoc(user User, newDoc Document) error {
	user.PostDocs = append(user.PostDocs, newDoc)
	err := dbconnect.MySQLcon.Save(&user).Error
	return err
}

func CreateUser(user User) User {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(bytes)
	dbconnect.MySQLcon.Create(&user)
	return user
}

func DeleteUser(userId string) bool {
	var user User
	result := dbconnect.MySQLcon.Where("employee_id = ?", userId).Delete(&user)
	return result.RowsAffected > 0
}

func UpdateUser(userId string, user User) (User, error) {
	err := dbconnect.MySQLcon.Model(&user).Where("employee_id = ?", userId).Updates(user).Error
	return user, err
}

func CheckUserPassword(email string, password string) bool {
	user := User{}
	dbconnect.MySQLcon.Where("email = ?", email).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
