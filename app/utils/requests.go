package utils

import (
	"fmt"
	"net/http"
)

func CheckMethod(r *http.Request, w http.ResponseWriter, method string) bool {
	if r.Method != method {
		_, err := w.Write([]byte(WrongMethodError))
		if err != nil {
			fmt.Println(err)
		}
		return false
	}
	return true
}
