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

}
