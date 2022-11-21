package apis

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SnehashishL/crudapp/db"
	"github.com/SnehashishL/crudapp/models"
	"github.com/gorilla/mux"
)

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// retrieve all tasks from Todos table
func AllTasks(w http.ResponseWriter, r *http.Request) {
	gdb := db.SetupDB()
	//defer db.Close()

	printMessage("Getting ToDos...")
	var tasks []models.Todos
	recs := gdb.Table("Todos").Find(&tasks)
	CheckErr(recs.Error)

	fmt.Println(tasks)
	json.NewEncoder(w).Encode(tasks)
}

// Get Task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID := params["taskid"]

	gdb := db.SetupDB()
	//defer db.Close()

	var task models.Todos
	printMessage("Getting ToDo by ID...")
	recs := gdb.Table("Todos").Where("TaskID", taskID).First(&task)
	fmt.Print(task.TaskID, task.Title)
	CheckErr(recs.Error)

	json.NewEncoder(w).Encode(task)

}

// Add task to db
func CreateTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.FormValue("taskid")
	title := r.FormValue("title")

	var response = models.JsonResponse{}

	if taskID == "" || title == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing taskID or title parameter."}
	} else {
		gdb := db.SetupDB()

		printMessage("Inserting task into DB")
		fmt.Println("Inserting new task with ID: " + taskID + " and title: " + title)
		task := models.Todos{TaskID: taskID, Title: title}
		gdb.Table("Todos").Create(&task)

		response = models.JsonResponse{Type: "success", Message: "The task has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// delete specific task by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID := params["taskid"]

	var response = models.JsonResponse{}

	if taskID == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing taskID parameter."}
	} else {

		gdb := db.SetupDB()
		//defer db.Close()

		printMessage("Deleting Task from DB")
		res := gdb.Table("Todos").Where("TaskID", taskID).Unscoped().Delete(&models.Todos{})
		CheckErr(res.Error)

		response = models.JsonResponse{Type: "success", Message: "The task has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskName := r.FormValue("taskname")
	taskID := params["taskid"]

	var response = models.JsonResponse{}

	if taskID == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing taskID parameter."}
	} else {

		db := db.SetupDB()
		//defer db.Close()

		if err := db.Where("TaskID", taskID).First(&models.Todos{}).Error; err != nil {
			CheckErr(err)
			return
		}

		printMessage("Updating Task in DB")
		fmt.Println("Updating new task with ID: " + taskID + " and name: " + taskName)
		res := db.Model(&models.Todos{}).Where("TaskID", taskID).Update("TaskName", taskName)
		CheckErr(res.Error)

		response = models.JsonResponse{Type: "success", Message: "The task has been updated successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
