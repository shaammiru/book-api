package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/syamilh/book-api/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Chi"))
	})

	router.Route("/books", loadBookRoutes)

	return router
}

func loadBookRoutes(r chi.Router) {
	bookHandler := &handler.BookHandler{}

	r.Post("/", bookHandler.Create)
	r.Get("/", bookHandler.List)
	r.Get("/{id}", bookHandler.GetByID)
	r.Put("/{id}", bookHandler.UpdateByID)
	r.Delete("/{id}", bookHandler.DeleteByID)
}
