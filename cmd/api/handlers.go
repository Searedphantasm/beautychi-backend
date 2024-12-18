package main

import (
	"errors"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (app *application) AllProductsHandler(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	limit, err := strconv.Atoi(queryParams.Get("limit"))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(queryParams.Get("page"))
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	offset := (page - 1) * limit

	products, err := app.Services.ProductServices.AllProductsService(limit, offset)
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

func (app *application) AllSubCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	subCategories, err := app.Services.SubCategoryServices.AllSubCategoryService()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, subCategories)
}

func (app *application) OneProductHandler(w http.ResponseWriter, r *http.Request) {
	product_id := chi.URLParam(r, "product_id")
	productID, err := strconv.Atoi(product_id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid product id"))
		return
	}

	product, err := app.Services.ProductServices.OneProductServiceByProductID(productID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, product)
}

func (app *application) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := app.readJSON(w, r, &product)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.Services.ProductServices.InsertProductService(product)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusCreated, product)
}

func (app *application) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	err := app.readJSON(w, r, &product)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	product.ID, err = strconv.Atoi(chi.URLParam(r, "product_id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid product id"))
		return
	}

	err = app.Services.ProductServices.UpdateProductService(product)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, product)
}

func (app *application) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	product_id := chi.URLParam(r, "product_id")
	productID, err := strconv.Atoi(product_id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid product id"))
		return
	}

	err = app.Services.ProductServices.DeleteProductService(productID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	err := app.readJSON(w, r, &category)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.Services.CategoryServices.CreateCategoryService(category)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusCreated, category)
}
