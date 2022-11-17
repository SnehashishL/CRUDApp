package models

import (
	"time"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	TaskID         string `json:"TaskID"`
	Title          string `json:"TaskName"`
	Description    string
	Status         string
	PlannedEndDate time.Time
	TagID          int
	UserID         string
}

type Tags struct {
	TagID   int
	TagName string
}
type Profiles struct {
	ID        string
	FirstName string
	LastName  string
	PIN       string
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Todos `json:"data"`
	Message string  `json:"message"`
}
