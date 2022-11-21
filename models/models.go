package models

import (
	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	//TaskID         string    `json:"taskId"`
	Title          string `json:"title"`
	Priority       string `json:"priority"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	PlannedEndDate string `json:"plannedEndDate"`
	TagID          string `json:"tagId"`
	UserID         string `jason:"userId"`
}

type Tags struct {
	gorm.Model
	TagName string `json:"tagName"`
}
type Profiles struct {
	gorm.Model
	Name string `json:"Name"`
	//PIN  string `json:"PIN"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Todos `json:"data"`
	Message string  `json:"message"`
}
