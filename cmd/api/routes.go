package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// products route
	mux.Get("/api/products", app.AllProductsHandler)
	mux.Post("/api/products", app.CreateProductHandler)

	mux.Get("/api/products/{product_id}", app.OneProductHandler)
	mux.Put("/api/products/{product_id}", app.UpdateProductHandler)
	mux.Put("/api/products/{product_id}/product-images", app.UpdateProductImageHandler)
	mux.Delete("/api/products/{product_id}", app.DeleteProductHandler)
	// TODO: Create endpoint for deleting single productImage.
	//mux.Delete("/api/products/{product_id}", app.DeleteProductImageHandler)

	mux.Get("/api/categories", app.AllCategoriesHandler)
	mux.Post("/api/categories", app.CreateCategoryHandler)
	mux.Put("/api/categories/{category_id}", app.UpdateCategoryHandler)
	mux.Get("/api/categories/{category_id}", app.OneCategoryHandler)
	mux.Delete("/api/categories/{category_id}", app.DeleteCategoryHandler)

	mux.Get("/api/brands", app.AllBrandsHandler)
	mux.Post("/api/brands", app.CreateBrandHandler)
	mux.Get("/api/brands/{brand_id}", app.OneBrandHandler)
	mux.Delete("/api/brands/{brand_id}", app.DeleteBrandHandler)
	mux.Put("/api/brands/{brand_id}", app.UpdateBrandHandler)

	mux.Get("/api/sub-categories", app.AllSubCategoriesHandler)
	mux.Post("/api/sub-categories", app.CreateSubCategoryHandler)
	mux.Put("/api/sub-categories/{sub_category_id}", app.UpdateSubCategoryHandler)
	mux.Delete("/api/sub-categories/{sub_category_id}", app.DeleteSubCategoryHandler)

	return mux
}
