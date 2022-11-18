package models

import (
	"time"

	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	TaskID         string    `json:"taskId"`
	Title          string    `json:"title"`
	Priority       string    `json:"priority"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	PlannedEndDate time.Time `json:"plannedEndDate"`
	TagID          int       `json:"tagId"`
	UserID         string    `jason:"userId"`
}

type Tags struct {
	TagID   int    `json:"tagId"`
	TagName string `json:"tagName"`
}
type Profiles struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PIN       string `json:"PIN"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Todos `json:"data"`
	Message string  `json:"message"`
}
