package apis

import (
	"encoding/json"
	"net/http"

	"github.com/SnehashishL/crudapp/db"
	"github.com/SnehashishL/crudapp/models"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	userName := r.FormValue("name")
	response := models.JsonResponse{}

	if userName == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing user name."}
	} else {
		gdb := db.SetupDB()
		user := models.Profiles{Name: userName}
		res := gdb.Table("Profiles").Create(&user)
		CheckErr(res.Error)
		response = models.JsonResponse{Type: "success", Message: "The user profile has been inserted successfully!"}
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userName := params["name"]

	response := models.JsonResponse{}

	if userName == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing user name."}
	} else {
		gdb := db.SetupDB()
		gdb.Table("Profiles").Where("Name", userName).Unscoped().Delete(&models.Profiles{})
		response = models.JsonResponse{Type: "success", Message: "The user profile has been deleted successfully!"}
	}
	json.NewEncoder(w).Encode(response)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	oldName := r.FormValue("name")
	newName := r.FormValue("newname")
	response := models.JsonResponse{}

	if newName == "" || oldName == "" {
		response = models.JsonResponse{Type: "error", Message: "You are missing user name(s)."}
	} else {
		gdb := db.SetupDB()
		newuser := models.Profiles{}
		rec := gdb.Table("Profiles").Where("Name", oldName).First(&newuser)
		if rec.Error != nil {
			CheckErr(rec.Error)
			return
		} else {
			newuser.Name = newName
			res := gdb.Table("Profiles").Where("Name", oldName).Updates(newuser)
			CheckErr(res.Error)
			response = models.JsonResponse{Type: "success", Message: "The user profile has been updated successfully!"}
		}
	}
	json.NewEncoder(w).Encode(response)

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	gdb := db.SetupDB()
	var profiles []models.Profiles

	recs := gdb.Table("Profiles").Find(&profiles)
	CheckErr(recs.Error)

	json.NewEncoder(w).Encode(profiles)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["name"]
	user := models.Profiles{}

	gdb := db.SetupDB()
	res := gdb.Table("Profiles").Where("Name", username).First(&user)
	CheckErr(res.Error)
	json.NewEncoder(w).Encode(user)
}
