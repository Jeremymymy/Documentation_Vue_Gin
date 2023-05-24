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
	Belong     string    `json:"Belong"`
	Title      string    `json:"Title" binding:"required" gorm:"not null"`
	Content    string    `json:"Content" binding:"required"`
	Vers       []Version `json:"Vers" gorm:"foreignKey:DocId"`
}

type Version struct {
	Id           uint      `json:"Id" gorm:"primaryKey;autoIncrement"`
	DocId        uint      `json:"DocId" gorm:"not null"`
	UpdaterId    string    `json:"UpdaterId" binding:"required" gorm:"not null"`
	UpdaterEmail string    `json:"UpdaterEmail" binding:"required" gorm:"not null"`
	UpdaterName  string    `json:"UpdaterName" binding:"required" gorm:"not null"`
	Title        string    `json:"Title" binding:"required" gorm:"not null"`
	Content      string    `json:"Content" binding:"required"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func CreateDoc(doc Document) (Document, error) {
	err := dbconnect.MySQLcon.Create(&doc).Error
	return doc, err
}

func GetDocById(docId uint) (Document, error) {
	var doc Document
	err := dbconnect.MySQLcon.Table("documents").Where("id = ?", docId).First(&doc).Error
	return doc, err
}

func GetDocByIdWithVersPreload(docId uint) (Document, error) {
	var doc Document
	err := dbconnect.MySQLcon.Preload("Vers").First(&doc, docId).Error
	return doc, err
}

func DeleteDocById(docId uint) bool {
	var doc Document
	result := dbconnect.MySQLcon.Table("documents").Where("id = ?", docId).Delete(&doc)
	return result.RowsAffected > 0
}

func UpdateDoc(docId uint, doc Document) (Document, error) {
	err := dbconnect.MySQLcon.Table("documents").Model(&doc).Where("id = ?", docId).Updates(doc).Error
	return doc, err
}

func CreateVer(ver Version) (Version, error) {
	err := dbconnect.MySQLcon.Create(&ver).Error
	return ver, err
}

func AddNewVer(doc Document, newVer Version) error {
	doc.Vers = append(doc.Vers, newVer)
	err := dbconnect.MySQLcon.Save(&doc).Error
	return err
}
