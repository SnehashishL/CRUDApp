package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SnehashishL/crudapp/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := routes.RouteInit()
	http.Handle("/", r)

	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
