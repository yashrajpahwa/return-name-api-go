package controllers

import (
	"net/http"

	"github.com/yashrajpahwa/return-name-api/utils"
)

func GetNameSplit(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	statement := utils.GetName(w, r, utils.SplitName(name))
	w.Write([]byte(statement))
}
