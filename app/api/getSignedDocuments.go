package api
//
//import (
//	"encoding/json"
//	"net/http"
//	"../utils"
//)
//
//func GetSignedDocumentsHandler(w http.ResponseWriter, r *http.Request) {
//	ok := utils.CheckMethod(r, w, "GET")
//	if !ok {
//		return
//	}
//	err := r.ParseForm()
//	ok = utils.HandleError(w, utils.ParseFormError, err)
//	if !ok {
//		return
//	}
//	uid := r.Form.Get("uids")
//	doc, err := documentController.GetSignedDocument(uid)
//	ok = utils.HandleError(w, `"error" : "error get signed document"`, err)
//	if !ok {
//		return
//	}
//	b, err := json.Marshal(doc)
//	if !utils.HandleError(w, utils.MarshalFieldsError, err) {
//		return
//	}
//	_, err = w.Write(b)
//	if !utils.HandleError(w, utils.WriteBytesError, err) {
//		return
//	}
//}
