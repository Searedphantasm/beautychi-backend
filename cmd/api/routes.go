package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	// products route
	mux.Get("/api/products", app.AllProductsHandler)
	mux.Post("/api/products", app.CreateProductHandler)
	mux.Get("/api/products/{product_identifier}", app.OneProductHandler)
	// product reviews
	mux.Get("/api/products/{product_id}/reviews", app.OneProductReviews)
	mux.Post("/api/products/{product_id}/reviews", app.CreateProductReview)
	mux.Put("/api/products/{product_id}", app.UpdateProductHandler)
	mux.Put("/api/products/{product_id}/product-images", app.UpdateProductImageHandler)
	mux.Delete("/api/products/{product_id}", app.DeleteProductHandler)

	// category route
	mux.Get("/api/categories", app.AllCategoriesHandler)
	mux.Post("/api/categories", app.CreateCategoryHandler)
	mux.Get("/api/categories/{category_identifier}", app.OneCategoryHandler)
	mux.Put("/api/categories/{category_id}", app.UpdateCategoryHandler)
	mux.Delete("/api/categories/{category_id}", app.DeleteCategoryHandler)

	// brand route
	mux.Get("/api/brands", app.AllBrandsHandler)
	mux.Post("/api/brands", app.CreateBrandHandler)
	mux.Get("/api/brands/{brand_identifier}", app.OneBrandHandler)
	mux.Put("/api/brands/{brand_id}", app.UpdateBrandHandler)
	mux.Delete("/api/brands/{brand_id}", app.DeleteBrandHandler)

	// sub-category route
	mux.Get("/api/sub-categories", app.AllSubCategoriesHandler)
	mux.Post("/api/sub-categories", app.CreateSubCategoryHandler)
	mux.Get("/api/sub-categories/{sub_category_identifier}", app.OneSubCategoryHandler)
	mux.Put("/api/sub-categories/{sub_category_id}", app.UpdateSubCategoryHandler)
	mux.Delete("/api/sub-categories/{sub_category_id}", app.DeleteSubCategoryHandler)

	// FIXME: ADDING authentication for admin.
	// customer route
	mux.Get("/api/customers", app.AllCustomersHandler)
	mux.Get("/api/customers/{customer_id}", app.OneCustomerHandler)

	return mux
}
