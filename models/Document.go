package models

import "time"

type Document struct {
	DocID     uint      `json:"DocID" gorm:"primaryKey"`
	Author    string    `json:"Author" binding:"required" gorm:"not null"`
	Title     string    `json:"Title" binding:"required" gorm:"not null"`
	Content   string    `json:"Content" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Vers      []Version `json:"Vers" gorm:"foreignKey:VerID"`
}

type Version struct {
	VerID     uint      `json:"VerID" gorm:"primaryKey"`
	Updater   string    `json:"Updater" binding:"required" gorm:"not null"`
	Title     string    `json:"Title" binding:"required" gorm:"not null"`
	Content   string    `json:"Content" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
