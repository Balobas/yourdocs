package api

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"../utils"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	err := r.ParseForm()
	if !utils.HandleError(w, utils.ParseFormError, err) {
		return
	}
	uid := r.Form.Get("uid")
	user, isFound, err := usersController.GetUser(uid)
	if !utils.HandleError(w, `"error" : "cant get user`, err) {
		return
	}
	if !isFound {
		utils.HandleError(w, `"error" : "review not found"`, errors.New(""))
	}
	b, err := json.Marshal(user)
	if !utils.HandleError(w, utils.MarshalFieldsError, err) {
		return
	}
	_, err = w.Write(b)
	if !utils.HandleError(w, utils.WriteBytesError, err) {
		return
	}
}