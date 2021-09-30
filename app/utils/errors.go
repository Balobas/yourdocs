package utils

import (
	"fmt"
	"net/http"
)

const (
	WriteBytesError      = `"error" : "cant write bytes`
	WrongMethodError     = `"error" : "wrong method`
	ParseFormError       = `"error" : "cant parse form`
	MarshalFieldsError   = `"error" : "cant marshal fields`
	UnmarshalFieldsError = `"error" : "cant unmarshal fields`
	ReadBodyError        = `"error" : "cant read request body`
)

func HandleError(w http.ResponseWriter, message string, err error) bool {
	if err != nil {
		_, err = w.Write([]byte(message + " " + err.Error() + `"`))
		if err != nil {
			fmt.Println(err.Error())
		}
		return false
	}
	return true
}
