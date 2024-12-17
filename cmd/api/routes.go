package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	// products route
	mux.Get("/api/products", app.AllProductsHandler)
	mux.Get("/api/products/{product_id}", app.OneProductHandler)

	mux.Post("/api/products", app.CreateProductHandler)

	mux.Get("/api/categories", app.AllCategoriesHandler)

	mux.Get("/api/brands", app.AllBrandsHandler)

	mux.Get("/api/sub-categories", app.AllSubCategoriesHandler)

	return mux
}
