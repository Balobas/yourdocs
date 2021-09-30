package api

import (
	"encoding/json"
	"net/http"
	"../utils"
)

func GetReviewCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	categories := reviewsController.GetReviewCategories()
	b, err := json.Marshal(categories)
	if !utils.HandleError(w, utils.MarshalFieldsError, err) {
		return
	}
	_, err = w.Write(b)
	if !utils.HandleError(w, utils.WriteBytesError, err) {
		return
	}
}
