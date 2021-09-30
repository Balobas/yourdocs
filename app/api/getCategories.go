package api

import (
	"../controllers"
	"../utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var categoriesController controllers.CategoriesController

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	categories, err := categoriesController.GetCategories(0, 0)
	if !utils.HandleError(w, `"error" : "error get categories`, err) {
		return
	}
	b, err := json.Marshal(categories)
	if !utils.HandleError(w, utils.MarshalFieldsError, err) {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}
}