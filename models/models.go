package models

type Todos struct {
	//gorm.Model
	TaskID   string `json:"taskid"`
	TaskName string `json:"taskname"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Todos `json:"data"`
	Message string  `json:"message"`
}
