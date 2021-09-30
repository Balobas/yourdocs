package api

import (
	"net/http"
	"../utils"
)

func DeleteReviewHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	if !utils.HandleError(w, utils.ParseFormError, r.ParseForm()) {
		return
	}
	userUid := r.Form.Get("userUid")
	reviewUid := r.Form.Get("reviewUid")
	utils.HandleError(w, `"error" : "cant delete review`, reviewsController.DeleteReview(userUid, reviewUid))
}