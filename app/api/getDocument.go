package api

import (
	"../controllers"
	"encoding/json"
	"net/http"
	"../utils"
)

var documentController controllers.DocumentController

func GetDocumentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	err := r.ParseForm()
	if !utils.HandleError(w, utils.ParseFormError, err) {
		return
	}
	uid := r.Form.Get("uid")
	//TODO: uid validation
	doc, err := documentController.GetDocument(uid)
	if !utils.HandleError(w, `"error" : "error get document`, err) {
		return
	}
	b, err := json.Marshal(doc)
	if !utils.HandleError(w, utils.MarshalFieldsError, err) {
		return
	}
	_, err = w.Write(b)
	if !utils.HandleError(w, utils.WriteBytesError, err) {
		return
	}

}