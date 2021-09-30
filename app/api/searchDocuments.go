package api

import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SearchDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "GET") {
		return
	}
	requestParams, err := ioutil.ReadAll(r.Body)
	if !utils.HandleError(w, utils.ReadBodyError, err) {
		return
	}
	type Params struct {
		Categories []string `json:"categories"`
		Names      []string `json:"names"`
		DocNumbers []string `json:"docNumbers"`
	}
	var params Params
	if !utils.HandleError(w, utils.UnmarshalFieldsError, json.Unmarshal(requestParams, &params)) {
		return
	}
	searchDocuments(w, params.Categories, params.Names, params.DocNumbers)
}

//TODO: добавить ограничение на количество записей
func searchDocuments(w http.ResponseWriter, categories, names, docNumbers []string) {
	docs, err := docController.Search(categories, names, docNumbers)
	if !utils.HandleError(w, `"error" : "search error`, err) {
		return
	}
	b, err := json.Marshal(docs)
	if !utils.HandleError(w, utils.MarshalFieldsError, err) {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}
}
