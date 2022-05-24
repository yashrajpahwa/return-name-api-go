package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/yashrajpahwa/return-name-api/controllers"
)

func Router(r chi.Router) {
	r.Get("/json/{name}", controllers.GetNameFromParam)
	r.Route("/string", func(r chi.Router) {
		r.Get("/", controllers.GetNameFromQueryParam)
		r.Get("/split", controllers.GetNameSplit)
	})
}
