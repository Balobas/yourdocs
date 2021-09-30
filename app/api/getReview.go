package api

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"../utils"
)

func GetReviewHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	err := r.ParseForm()
	if !utils.HandleError(w, utils.ParseFormError, err) {
		return
	}
	uid := r.Form.Get("uid")
	review, isFound, err := reviewsController.GetReview(uid)
	if !utils.HandleError(w, `"error" : "cant get review`, err) {
		return
	}
	if !isFound {
		utils.HandleError(w, `"error" : "review not found"`, errors.New(""))
	}
	b, err := json.Marshal(review)
	if !utils.HandleError(w, utils.MarshalFieldsError, err) {
		return
	}
	_, err = w.Write(b)
	if !utils.HandleError(w, utils.WriteBytesError, err) {
		return
	}
}