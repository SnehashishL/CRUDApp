package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SnehashishL/crudapp/apis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := apis.RouteInit()
	http.Handle("/", r)

	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

	// db, err := sql.Open("mysql", "root:sneh@mysql123@tcp(127.0.0.1:3306)/todos")
	// // if there is an error opening the connection, handle it
	// if err != nil {
	// 	log.Print(err.Error())
	// }
	// defer db.Close()

	// // Execute the query
	// rows, err := db.Query("SELECT * FROM todos.todos")

	// CheckErr(err)

	// var tasks []Task

	// for rows.Next() {
	// 	var id int
	// 	var taskid string
	// 	var taskname string

	// 	err = rows.Scan(&id, &taskid, &taskname)

	// 	CheckErr(err)

	// 	tasks = append(tasks, Task{TaskID: taskid, TaskName: taskname})
	// }

	// fmt.Println("Here are all the tasks from db:")
	// fmt.Printf("%v", tasks)

}
