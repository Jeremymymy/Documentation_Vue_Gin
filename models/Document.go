package models

import (
	"documentation/dbconnect"
	"time"

	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	AuthorId   string    `json:"AuthorId" gorm:"type:varchar(255);not null;index"`
	AuthorName string    `json:"AuthorName" gorm:"not null"`
	Title      string    `json:"Title" binding:"required" gorm:"not null"`
	Content    string    `json:"Content" binding:"required"`
	Vers       []Version `json:"Vers" gorm:"foreignKey:DocId"`
}

type Version struct {
	Id        uint      `json:"Id" gorm:"primaryKey;autoIncrement"`
	DocId     uint      `json:"DocId" gorm:"not null"`
	Updater   string    `json:"Updater" binding:"required" gorm:"not null"`
	Title     string    `json:"Title" binding:"required" gorm:"not null"`
	Content   string    `json:"Content" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func CreateDoc(doc Document) (Document, error) {
	err := dbconnect.MySQLcon.Create(&doc).Error
	return doc, err
}
