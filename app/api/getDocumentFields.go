package api

import (
	"../utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetDocumentFieldsHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	err := r.ParseForm()
	if !utils.HandleError(w, utils.ParseFormError, err) {
		return
	}
	uid := r.Form.Get("uid")
	getDocumentFields(w, uid)
}

func getDocumentFields(w http.ResponseWriter, uid string) {
	fields, err := docController.GetDocumentFields(uid)
	if !utils.HandleError(w, `"error" : "get fields error `, err) {
		return
	}
	b, err := json.Marshal(fields)
	if !utils.HandleError(w, utils.MarshalFieldsError, err) {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}
}
