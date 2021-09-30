package api

import (
	"../controllers"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"../utils"
	"../models"
)

var usersController controllers.UsersController

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "POST") {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if !utils.HandleError(w, utils.ReadBodyError, err) {
		return
	}
	var user models.User
	if !utils.HandleError(w, utils.UnmarshalFieldsError, json.Unmarshal(body, &user)) {
		return
	}
	uid, err := usersController.CreateUser(user)
	if !utils.HandleError(w, `"error" : "cant create user`, err) {
		return
	}
	_, err = w.Write([]byte(uid))
	utils.HandleError(w, utils.WriteBytesError, err)
}