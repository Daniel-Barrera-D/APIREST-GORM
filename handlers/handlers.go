package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	users := models.Users{}
	db.Database.Find(&users)
	sendData(rw, users, http.StatusOK)

}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, error := getUserById(r); error != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)
	}
}

func getUserById(r *http.Request) (models.User, *gorm.DB) {

	//Obtener ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}

	if error := db.Database.First(&user, userId); error.Error != nil {
		return user, error
	} else {
		return user, nil
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	user := models.User{}

	decoder := json.NewDecoder(r.Body)

	if error := decoder.Decode(&user); error != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&user)
		sendData(rw, user, http.StatusCreated)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	var userId int64

	if user_ant, error := getUserById(r); error != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		userId = user_ant.Id

		user := models.User{}

		decoder := json.NewDecoder(r.Body)

		if error := decoder.Decode(&user); error != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.Database.Save(&user)
			sendData(rw, user, http.StatusOK)
		}
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	if user, error := getUserById(r); error != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.Database.Delete(&user)
		sendData(rw, user, http.StatusOK)
	}
}
