package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title    string `json:"title"`
	Brand    string `json:"brand"`
	Platform string `json:"platform"`
	DueDate  string `json:"due_date"`
	Payment  string `json:"payment"`
	Status   string `json:"status"`
}
