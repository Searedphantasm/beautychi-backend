package main

import (
	"net/http"
)

func (app *application) AllProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := app.Services.ProductServices.AllProductsService()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, products)
}

func (app *application) AllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := app.Services.CategoryServices.AllCategoryService()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, categories)
}

func (app *application) AllBrandsHandler(w http.ResponseWriter, r *http.Request) {
	brands, err := app.Services.BrandServices.AllBrandsService()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, brands)
}
