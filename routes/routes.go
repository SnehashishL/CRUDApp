package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SnehashishL/crudapp/db"
	"github.com/SnehashishL/crudapp/models"
	"github.com/gorilla/mux"
)

// const (
// 	DB_USER     = "mysql"
// 	DB_PASSWORD = "sneh@mysql123"
// 	DB_NAME     = "todos"
// )

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

func RouteInit() *mux.Router {
	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all tasks
	router.HandleFunc("/tasks/", AllTasks).Methods("GET")

	// Get task by ID
	router.HandleFunc("/myTask/{taskid}", GetTask).Methods("GET")

	// Create a task
	router.HandleFunc("/addTask/", CreateTask).Methods("POST")

	// Delete a specific task by the taskID
	router.HandleFunc("/deleteTask/{taskid}", DeleteTask).Methods("DELETE")

	// Update a task
	router.HandleFunc("/updateTask/{taskid}", UpdateTask).Methods("PUT")

	return router
}

// retrieve all tasks from Todos table
func AllTasks(w http.ResponseWriter, r *http.Request) {

	gdb := db.SetupDB()
	//defer db.Close()

	printMessage("Getting ToDos...")

	//rows, err := db.Query("SELECT * FROM todos.todos")
	var tasks []models.Todos
	recs := gdb.Find(&tasks)
	CheckErr(recs.Error)
	json.NewEncoder(w).Encode(tasks)
}

// Get Task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	taskID := params["taskid"]

	gdb := db.SetupDB()
	//defer db.Close()

	printMessage("Getting ToDo by ID...")

	var task models.Todos

	gdb.Find(&task, "task_id = ?", taskID)
	fmt.Print(task.TaskID, task.TaskName)
	//CheckErr(recs.Error)

	//var response = JsonResponse{Type: "success", Data: task}
	json.NewEncoder(w).Encode(task)

}

// Add task to db
func CreateTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.FormValue("taskid")
	taskName := r.FormValue("taskname")

	var response = models.JsonResponse{}

	if taskID == "" || taskName == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing taskID or taskName parameter."}
	} else {

		//db setup
		db := db.SetupDB()

		printMessage("Inserting task into DB")

		fmt.Println("Inserting new task with ID: " + taskID + " and name: " + taskName)

		//var lastInsertID int
		task := models.Todos{TaskID: taskID, TaskName: taskName}
		db.Create(&task)

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
		// db setup
		db := db.SetupDB()
		//defer db.Close()

		var task models.Todos

		printMessage("Deleting Task from DB")

		//_, err := db.Exec("DELETE FROM todos.todos where taskID = ?", taskID)
		res := db.Where("task_id = ?", taskID).Delete(&task)
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
		// db setup
		db := db.SetupDB()
		//defer db.Close()

		var task models.Todos
		if err := db.Where("task_id = ?", taskID).First(&task).Error; err != nil {
			CheckErr(err)
			return
		}
		printMessage("Updating Task in DB")

		fmt.Println("Updating new task with ID: " + taskID + " and name: " + taskName)

		res := db.Model(&task).Where("task_id = ?", taskID).Update("task_name", taskName)

		CheckErr(res.Error)

		response = models.JsonResponse{Type: "success", Message: "The task has been updated successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
