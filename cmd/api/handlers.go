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

func (app *application) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := app.readJSON(w, r, &category)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	category.ID, err = strconv.Atoi(chi.URLParam(r, "category_id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid category id"))
		return
	}

	err = app.Services.CategoryServices.UpdateCategoryService(category)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, category)
}

// FIXME: DROP database and add image_key field and refactor database functions.

func (app *application) UpdateProductImageHandler(w http.ResponseWriter, r *http.Request) {
	product_id := chi.URLParam(r, "product_id")
	productID, err := strconv.Atoi(product_id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid product id"))
		return
	}

	var reqBody struct {
		ProductImages []models.ProductImage `json:"product_images"`
	}
	err = app.readJSON(w, r, &reqBody)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.Services.ProductServices.UpdateProductImagesService(productID, reqBody.ProductImages)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, reqBody)
}

func (app *application) OneCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category_id := chi.URLParam(r, "category_id")
	categoryID, err := strconv.Atoi(category_id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid category id"))
		return
	}

	category, err := app.DB.GetCategoryByID(categoryID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, category)
}

func (app *application) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category_id := chi.URLParam(r, "category_id")
	categoryID, err := strconv.Atoi(category_id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid category id"))
		return
	}

	err = app.Services.CategoryServices.DeleteCategoryService(categoryID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) CreateBrandHandler(w http.ResponseWriter, r *http.Request) {

	var brand models.Brand
	err := app.readJSON(w, r, &brand)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.Services.BrandServices.CreateBrandService(brand)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusCreated, brand)
}

func (app *application) DeleteBrandHandler(w http.ResponseWriter, r *http.Request) {
	brand_id := chi.URLParam(r, "brand_id")
	brandID, err := strconv.Atoi(brand_id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid brand id"))
		return
	}

	err = app.Services.BrandServices.DeleteBrandService(brandID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) OneBrandHandler(w http.ResponseWriter, r *http.Request) {
	brand_id := chi.URLParam(r, "brand_id")
	brandID, err := strconv.Atoi(brand_id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid brand id"))
		return
	}

	brand, err := app.DB.GetOneBrandByID(brandID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.writeJSON(w, http.StatusOK, brand)
}

func (app *application) UpdateBrandHandler(w http.ResponseWriter, r *http.Request) {
	brand_id := chi.URLParam(r, "brand_id")
	brandID, err := strconv.Atoi(brand_id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid brand id"))
		return
	}

	var brand models.Brand
	err = app.readJSON(w, r, &brand)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	brand.ID = brandID

	err = app.Services.BrandServices.UpdateBrandService(brand)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, brand)
}
