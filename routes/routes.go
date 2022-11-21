package routes

import (
	"github.com/SnehashishL/crudapp/apis"
	"github.com/gorilla/mux"
)

func RouteInit() *mux.Router {
	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints
	// Get all tasks
	router.HandleFunc("/tasks/", apis.AllTasks).Methods("GET")
	// Get task by ID
	router.HandleFunc("/myTask/{taskid}", apis.GetTask).Methods("GET")
	// Create a task
	router.HandleFunc("/addTask/", apis.CreateTask).Methods("POST")
	// Delete a specific task by the taskID
	router.HandleFunc("/deleteTask/{taskid}", apis.DeleteTask).Methods("DELETE")
	// Update a task
	router.HandleFunc("/updateTask/title/{taskid}", apis.UpdateTitle).Methods("PUT")

	// User handles
	// get all profiles
	router.HandleFunc("/users/", apis.GetAllUsers).Methods("GET")
	// create user
	router.HandleFunc("/addUser/", apis.CreateUser).Methods("POST")
	// update user
	router.HandleFunc("/updateUser/", apis.UpdateUser).Methods("PUT")
	// delete user
	router.HandleFunc("/deleteUser/{name}", apis.DeleteUser).Methods("DELETE")
	// get one user
	router.HandleFunc("/getUser/{name}", apis.GetUser).Methods("GET")

	return router
}
