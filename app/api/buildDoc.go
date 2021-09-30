package api

import (
	"../controllers"
	"../utils"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

var docController controllers.DocumentController

func BuildDocHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.CheckMethod(r, w, "POST") {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if !utils.HandleError(w, utils.ReadBodyError, err) {
		return
	}
	type Params struct {
		UID    string            `json:"uid"`
		Fields map[string]string `json:"fields"`
	}
	var params Params
	if !utils.HandleError(w, utils.UnmarshalFieldsError, json.Unmarshal(body, &params)) {
		return
	}
	checkPart := strings.Replace(params.UID, "doc", "", 1)
	_, err = uuid.FromString(checkPart)
	if !utils.HandleError(w, `"error" : "Invalid uid`, err) {
		return
	}
	fields := params.Fields
	fmt.Println(fields["name"])
	err = docController.BuildDocument(w, params.UID, fields)
	utils.HandleError(w, `"error" : "build document error `, err)
}
