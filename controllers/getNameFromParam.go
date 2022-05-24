package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yashrajpahwa/return-name-api/utils"
)

func GetNameFromParam(w http.ResponseWriter, r *http.Request) {
	statement := utils.GetName(w, r, chi.URLParam(r, "name"))
	resp := make(map[string]string)
	resp["message"] = statement
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
