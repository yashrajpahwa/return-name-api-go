package utils

import "net/http"

func GetName(w http.ResponseWriter, r *http.Request, name string) string {
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return "Please provide a name"
	}

	output := "Hello " + name

	return output
}
