package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// const (
// 	DB_USER     = "mysql"
// 	DB_PASSWORD = "sneh@mysql123"
// 	DB_NAME     = "todos"
// )

// DB set up
func SetupDB() *sql.DB {
	db, err := sql.Open("mysql", "root:sneh@mysql123@tcp(127.0.0.1:3306)/todos")
	// error handler
	if err != nil {
		//log.Print(err.Error())
		CheckErr(err)
	}
	//defer db.Close()
	return db
}

type Task struct {
	TaskID   string `json:"taskid"`
	TaskName string `json:"taskname"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Task `json:"data"`
	Message string `json:"message"`
}

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

	// Delete all
	router.HandleFunc("/deleteAll/", DeleteAll).Methods("DELETE")

	// Update a task
	router.HandleFunc("/updateTask/{taskid}", UpdateTask).Methods("PUT")

	return router
}

// retrieve all tasks from Todos table
func AllTasks(w http.ResponseWriter, r *http.Request) {

	db := SetupDB()
	defer db.Close()

	printMessage("Getting ToDos...")

	rows, err := db.Query("SELECT * FROM todos.todos")

	CheckErr(err)

	var tasks []Task

	for rows.Next() {
		var id int
		var taskid string
		var taskname string

		err = rows.Scan(&id, &taskid, &taskname)

		CheckErr(err)

		tasks = append(tasks, Task{TaskID: taskid, TaskName: taskname})
	}

	var response = JsonResponse{Type: "success", Data: tasks}

	json.NewEncoder(w).Encode(response)
}

// Get Task by ID
func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID := params["taskid"]

	db := SetupDB()
	defer db.Close()

	printMessage("Getting ToDo with specified id ...")

	//row := db.QueryRow("SELECT * FROM todos.todos WHERE taskID = {taskID}", taskID)
	row := db.QueryRow("SELECT * FROM todos.todos where taskid = ?", taskID)

	var task []Task
	var id int
	var taskid string
	var taskname string

	err := row.Scan(&id, &taskid, &taskname)

	CheckErr(err)

	task = append(task, Task{TaskID: taskid, TaskName: taskname})

	var response = JsonResponse{Type: "success", Data: task}

	json.NewEncoder(w).Encode(response)

}

// Add task to db
func CreateTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.FormValue("taskid")
	taskName := r.FormValue("taskname")

	var response = JsonResponse{}

	if taskID == "" || taskName == "" {
		response = JsonResponse{Type: "error", Message: "You are missing taskID or taskName parameter."}
	} else {

		//db setup
		db := SetupDB()
		defer db.Close()

		printMessage("Inserting task into DB")

		fmt.Println("Inserting new task with ID: " + taskID + " and name: " + taskName)

		//var lastInsertID int
		_, err := db.Exec("INSERT INTO todos.todos(taskid, taskname) VALUES(?, ?)", taskID, taskName)
		CheckErr(err)

		response = JsonResponse{Type: "success", Message: "The task has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// delete specific task by ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	taskID := params["taskid"]

	var response = JsonResponse{}

	if taskID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing taskID parameter."}
	} else {
		// db setup
		db := SetupDB()
		defer db.Close()

		printMessage("Deleting Task from DB")

		_, err := db.Exec("DELETE FROM todos.todos where taskID = ?", taskID)

		CheckErr(err)

		response = JsonResponse{Type: "success", Message: "The task has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// delete all tasks from db
func DeleteAll(w http.ResponseWriter, r *http.Request) {

	// db setup
	db := SetupDB()
	defer db.Close()

	printMessage("Deleting all tasks...")

	_, err := db.Exec("DELETE FROM todos.todos")

	CheckErr(err)

	printMessage("All tasks have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All tasks have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskName := r.FormValue("taskname")
	taskID := params["taskid"]

	var response = JsonResponse{}

	if taskID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing taskID parameter."}
	} else {
		// db setup
		db := SetupDB()
		defer db.Close()

		printMessage("Updating Task in DB")

		fmt.Println("Updating new task with ID: " + taskID + " and name: " + taskName)

		_, err := db.Exec("UPDATE todos.todos SET taskname = ? where taskID = ?", taskName, taskID)

		CheckErr(err)

		response = JsonResponse{Type: "success", Message: "The task has been updated successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}
