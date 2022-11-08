package db

import (
	"fmt"

	"github.com/SnehashishL/crudapp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgresql"
	DB_NAME     = "demo"
)

// // DB set up
func SetupDB() *gorm.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	gdb, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	gdb.AutoMigrate(&models.Todos{})

	return gdb
}

// DB set up for mysql server
// func SetupDB() *gorm.DB {
// 	//db, err := sql.Open("mysql", "root:sneh@mysql123@tcp(127.0.0.1:3306)/todos")
// 	dsn := "root:sneh@mysql123@tcp(127.0.0.1:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	// error handler
// 	if err != nil {
// 		panic(err)
// 	}
// 	//defer db.Close()
// 	return db
// }
