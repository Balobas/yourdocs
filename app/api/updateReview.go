package api

import (
	"../models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"../utils"
)

func UpdateReviewHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "POST") {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if !utils.HandleError(w, utils.ReadBodyError, err) {
		return
	}
	type Params struct {
		ExecutorUid string `json:"executorUid"`
		Review models.Review `json:"review"`
	}
	var params Params
	if !utils.HandleError(w, utils.UnmarshalFieldsError, json.Unmarshal(body, &params)) {
		return
	}
	utils.HandleError(w, `"error" : "cant update review`, reviewsController.UpdateReview(params.ExecutorUid, params.Review))
}