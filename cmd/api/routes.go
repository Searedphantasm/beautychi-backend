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
	mux.Post("/api/products", app.CreateProductHandler)

	mux.Get("/api/products/{product_id}", app.OneProductHandler)
	mux.Put("/api/products/{product_id}", app.UpdateProductHandler)
	mux.Put("/api/products/{product_id}/product-images", app.UpdateProductImageHandler)
	mux.Delete("/api/products/{product_id}", app.DeleteProductHandler)

	mux.Get("/api/categories", app.AllCategoriesHandler)
	mux.Post("/api/categories", app.CreateCategoryHandler)
	mux.Put("/api/categories/{category_id}", app.UpdateCategoryHandler)

	mux.Get("/api/brands", app.AllBrandsHandler)

	mux.Get("/api/sub-categories", app.AllSubCategoriesHandler)

	return mux
}
