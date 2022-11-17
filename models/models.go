package models

import (
	"time"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	TaskID         string `json:"Task ID"`
	Title          string
	Description    string
	Status         string
	PlannedEndDate time.Time `json:"Planned End Date"`
	TagID          int       `json:"Tag ID"`
	UserID         string    `jason:"User ID"`
}

type Tags struct {
	TagID   int    `json:"Tag ID"`
	TagName string `json:"Tag Name"`
}
type Profiles struct {
	ID        string
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	PIN       string
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Todos `json:"data"`
	Message string  `json:"message"`
}
