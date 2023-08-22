package models

import "time"

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Book struct {
	Id           int       `json:"id" gorm:"column:id"`
	Title        string    `json:"title" gorm:"column:title"`
	Descriptions string    `json:"descriptions" gorm:"column:descriptions"`
	Author       string    `json:"author" gorm:"column:author"`
	CreateAt     time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt     time.Time `json:"update_at" gorm:"column:update_at"`
	DeleteAt     time.Time `json:"delete_at" gorm:"column:delete_atd"`
}
