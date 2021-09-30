package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"../utils"
	"../models"
)

func AnswerOnReviewHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "POST") {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if !utils.HandleError(w, utils.ReadBodyError, err) {
		return
	}
	type Params struct {
		ReviewUid string `json:"reviewUid"`
		Answer models.ReviewAnswer `json:"answer"`
	}
	var params Params
	if !utils.HandleError(w, utils.UnmarshalFieldsError, json.Unmarshal(body, &params)) {
		return
	}
	uid, err := reviewsController.AnswerOnReview(params.ReviewUid, params.Answer)
	if !utils.HandleError(w, `"error" : "cant create answer`, err) {
		return
	}
	_, err = w.Write([]byte(uid))
	utils.HandleError(w, utils.WriteBytesError, err)
}
