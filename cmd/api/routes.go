package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Get("/api/products", app.AllProductsHandler)

	mux.Get("/api/categories", app.AllCategoriesHandler)

	mux.Get("/api/brands", app.AllBrandsHandler)

	return mux
}
