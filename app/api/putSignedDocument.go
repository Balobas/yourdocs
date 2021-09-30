package api

import (
	"../models"
	"../utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PutSignedDocumentHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "POST") {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if !utils.HandleError(w, utils.ReadBodyError, err) {
		return
	}
	var doc models.SignedDocument
	if !utils.HandleError(w, utils.UnmarshalFieldsError, json.Unmarshal(body, &doc)) {
		return
	}
	uid, err := documentController.PutSignedDocument(doc)
	if !utils.HandleError(w, `"error" : "cant put signed file`, err) {
		return
	}
	_, err = w.Write([]byte(uid))
	utils.HandleError(w, utils.WriteBytesError, err)
}