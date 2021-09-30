package api

import (
	"encoding/json"
	"net/http"
	"../utils"
)

func GetReviewAnswerHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	err := r.ParseForm()
	if !utils.HandleError(w, utils.ParseFormError, err) {
		return
	}
	uid := r.Form.Get("uid")
	answer, err := reviewsController.GetReviewAnswer(uid)
	if !utils.HandleError(w, `"error" : "cant get review`, err) {
		return
	}
	b, err := json.Marshal(answer)
	if !utils.HandleError(w, utils.MarshalFieldsError, err) {
		return
	}
	_, err = w.Write(b)
	if !utils.HandleError(w, utils.WriteBytesError, err) {
		return
	}
}