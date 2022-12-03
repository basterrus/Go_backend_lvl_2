package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/lapitskyss/go_backend_2/internal/server/handler"
)

func Routes(r *chi.Mux, handler *handler.Handler) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/product", handler.CreateProduct)
		r.Get("/product/search", handler.SearchProduct)
	})
}
