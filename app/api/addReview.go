package api

import (
	"../controllers"
	"../models"
	"../utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var reviewsController controllers.ReviewsController

func AddReviewHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "POST") {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if !utils.HandleError(w, utils.ReadBodyError, err) {
		return
	}
	var review models.Review
	if !utils.HandleError(w, utils.UnmarshalFieldsError, json.Unmarshal(body, &review)) {
		return
	}
	uid, err := reviewsController.AddReview(review)
	if !utils.HandleError(w, `"error" : "cant add review`, err) {
		return
	}
	_, err = w.Write([]byte(uid))
	utils.HandleError(w, utils.WriteBytesError, err)
}