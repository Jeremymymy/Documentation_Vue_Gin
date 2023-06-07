package models

import (
	"documentation/dbconnect"

	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	AuthorId   string    `json:"AuthorId" gorm:"type:varchar(255);not null;index"`
	AuthorName string    `json:"AuthorName" gorm:"not null"`
	Belong     string    `json:"Belong" binding:"oneof=HR Sales Marketing IT Finance Public"`
	Title      string    `json:"Title" binding:"required" gorm:"not null"`
	Content    string    `json:"Content" binding:"required"`
	Vers       []Version `json:"Vers" gorm:"foreignKey:DocId"`
}

type Version struct {
	gorm.Model
	DocId        uint   `json:"DocId" gorm:"not null"`
	VerNum       int    `json:"VerNum" gorm:"default:1"`
	UpdaterId    string `json:"UpdaterId" binding:"required" gorm:"type:varchar(255);not null;index"`
	UpdaterEmail string `json:"UpdaterEmail" binding:"required" gorm:"not null"`
	UpdaterName  string `json:"UpdaterName" binding:"required" gorm:"not null"`
	Title        string `json:"Title" binding:"required" gorm:"not null"`
	Content      string `json:"Content" binding:"required"`
}

type Collect struct {
	gorm.Model
	UserId     string `json:"UserId" gorm:"type:varchar(255);size:255;not null"`
	AuthorId   string `json:"AuthorId" gorm:"type:varchar(255);not null"`
	AuthorName string `json:"AuthorName" gorm:"not null"`
	Belong     string `json:"Belong"`
	DocId      uint   `json:"DocId" gorm:"not null"`
	Title      string `json:"Title" binding:"required" gorm:"not null"`
	Content    string `json:"Content" binding:"required"`
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

func GetDepartmentDocs(department string) ([]Document, error) {
	var docs []Document
	err := dbconnect.MySQLcon.Table("documents").Where("belong = ?", department).Find(&docs).Error
	return docs, err
}

func DeleteDoc(docId uint) bool {
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

func GetVerById(verId uint) (Version, error) {
	var ver Version
	err := dbconnect.MySQLcon.Table("versions").Where("id = ?", verId).First(&ver).Error
	return ver, err
}

func DocAddNewVer(doc Document, newVer Version) error {
	doc.Vers = append(doc.Vers, newVer)
	err := dbconnect.MySQLcon.Save(&doc).Error
	return err
}

func GetAllVers() []Version {
	var vers []Version
	dbconnect.MySQLcon.Find(&vers)
	return vers
}

func CreateCollect(collect Collect) (Collect, error) {
	err := dbconnect.MySQLcon.Create(&collect).Error
	return collect, err
}

func GetAllCollects() []Collect {
	var cols []Collect
	dbconnect.MySQLcon.Find(&cols)
	return cols
}

func DeleteCol(colId uint) bool {
	var col Collect
	result := dbconnect.MySQLcon.Table("collects").Where("id = ?", colId).Delete(&col)
	return result.RowsAffected > 0
}
