package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgresql"
	DB_NAME     = "demo"
)

// // DB set up
// func SetupDB() *sql.DB {
// 	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
// 	db, err := sql.Open("postgres", dbinfo)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return db
// }

// DB set up
func SetupDB() *gorm.DB {
	//db, err := sql.Open("mysql", "root:sneh@mysql123@tcp(127.0.0.1:3306)/todos")
	dsn := "root:sneh@mysql123@tcp(127.0.0.1:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// error handler
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}
