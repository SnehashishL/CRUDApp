package models

type Todos struct {
	//gorm.Model
	TaskID   string `json:"TaskID"`
	TaskName string `json:"TaskName"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Todos `json:"data"`
	Message string  `json:"message"`
}
