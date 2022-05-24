package controllers

import (
	"net/http"

	"github.com/yashrajpahwa/return-name-api/utils"
)

func GetNameFromQueryParam(w http.ResponseWriter, r *http.Request) {
	statement := utils.GetName(w, r, r.URL.Query().Get("name"))
	w.Write([]byte(statement))
}
