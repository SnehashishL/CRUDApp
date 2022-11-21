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
	recs := gdb.Table("Todos").Where("ID", taskID).First(&task)
	fmt.Print(task.ID, task.Title)
	CheckErr(recs.Error)

	json.NewEncoder(w).Encode(task)

}

// Add task to db
func CreateTask(w http.ResponseWriter, r *http.Request) {
	//taskID := r.FormValue("taskid")
	title := r.FormValue("title")
	priority := r.FormValue("priority")
	description := r.FormValue("description")
	//status := r.FormValue("status")
	endDate := r.FormValue("date")
	tagID := r.FormValue("tagid")
	userID := r.FormValue("userid")

	var response = models.JsonResponse{}

	if userID == "" || title == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing UserID or title parameter."}
	} else {
		gdb := db.SetupDB()

		printMessage("Inserting task into DB")
		fmt.Println("Inserting new task..." + title)
		task := models.Todos{Title: title, Priority: priority, Description: description, Status: "TBD", PlannedEndDate: endDate, TagID: tagID, UserID: userID}
		gdb.Table("Todos").Create(&task)

		response = models.JsonResponse{Type: "success", Message: "The task has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// delete specific task by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID := params["taskid"]

	var response = models.JsonResponse{}

	if ID == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing taskID parameter."}
	} else {

		gdb := db.SetupDB()

		printMessage("Deleting Task from DB")
		res := gdb.Table("Todos").Where("ID", ID).Unscoped().Delete(&models.Todos{})
		CheckErr(res.Error)

		response = models.JsonResponse{Type: "success", Message: "The task has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateTitle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	newTitle := r.FormValue("taskname")
	ID := params["taskid"]

	var response = models.JsonResponse{}

	if ID == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing taskID parameter."}
	} else {

		db := db.SetupDB()
		task := models.Todos{}

		rec := db.Table("Todos").Where("ID", ID).First(&task)
		if rec.Error != nil {
			CheckErr(rec.Error)
			return
		}
		task.Title = newTitle
		printMessage("Updating Task in DB")
		fmt.Println("Updating new task with ID: " + ID + " and title: " + newTitle)
		rec = db.Table("Todos").Where("ID", ID).Updates(&task)
		CheckErr(rec.Error)

		response = models.JsonResponse{Type: "success", Message: "The task has been updated successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
