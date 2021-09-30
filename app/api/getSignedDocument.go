package api

import (
	"encoding/json"
	"net/http"
	"../utils"
)

func GetSignedDocumentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	err := r.ParseForm()
	if !utils.HandleError(w, utils.ParseFormError, err) {
		return
	}
	uid := r.Form.Get("uid")
	doc, err := documentController.GetSignedDocument(uid)
	if !utils.HandleError(w, `"error" : "error get signed document`, err) {
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
