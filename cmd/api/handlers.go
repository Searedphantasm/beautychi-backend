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
	//test := r.Header.Get("Cookie")

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

	// TODO : Add filter
	//getting optional params
	search := queryParams.Get("search")
	filter := queryParams.Get("filter")
	category := queryParams.Get("category")

	optionalParams := models.OptionalQueryParams{
		Search:          search,
		Filter:          filter,
		ProductCategory: category,
	}

	offset := (page - 1) * limit

	products, err := app.Services.ProductServices.AllProductsService(limit, offset, optionalParams)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, products)
}

func (app *application) AllCategoriesHandler(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	search := queryParams.Get("search")
	filter := queryParams.Get("filter")

	optionalParams := models.OptionalQueryParams{
		Search: search,
		Filter: filter,
	}

	categories, err := app.Services.CategoryServices.AllCategoryService(optionalParams)
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

func (app *application) CreateSubCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var subCategory models.SubCategory

	err := app.readJSON(w, r, &subCategory)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.Services.SubCategoryServices.InsertSubCategoryService(subCategory)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusCreated, subCategory)
}

func (app *application) OneSubCategoryHandler(w http.ResponseWriter, r *http.Request) {
	identifier := chi.URLParam(r, "sub_category_identifier")
	subCategoryID, err := strconv.Atoi(identifier)
	var subCategory *models.SubCategory
	if err != nil {
		subCategory, err = app.Services.SubCategoryServices.OneSubCategoryServiceByIDOrSlug(0, identifier)
	} else {
		subCategory, err = app.Services.SubCategoryServices.OneSubCategoryServiceByIDOrSlug(subCategoryID, "")
	}

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, subCategory)
}

func (app *application) UpdateSubCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var subCategory models.SubCategory

	err := app.readJSON(w, r, &subCategory)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	subCategory.ID, err = strconv.Atoi(chi.URLParam(r, "sub_category_id"))
	if err != nil {
		app.errorJSON(w, errors.New("invalid sub category id"))
		return
	}

	err = app.Services.SubCategoryServices.UpdateSubCategoryService(subCategory)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, subCategory)
}

func (app *application) DeleteSubCategoryHandler(w http.ResponseWriter, r *http.Request) {
	sub_category_id := chi.URLParam(r, "sub_category_id")
	subCategoryID, err := strconv.Atoi(sub_category_id)
	if err != nil {
		app.errorJSON(w, errors.New("invalid sub category id"))
		return
	}

	err = app.Services.SubCategoryServices.DeleteSubCategoryService(subCategoryID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) OneProductHandler(w http.ResponseWriter, r *http.Request) {
	identifier := chi.URLParam(r, "product_identifier")
	productID, err := strconv.Atoi(identifier)
	var product *models.Product
	if err != nil {
		product, err = app.Services.ProductServices.OneProductServiceByProductIDOrSlug(0, identifier)
	} else {
		product, err = app.Services.ProductServices.OneProductServiceByProductIDOrSlug(productID, "")
	}

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, product)
}

func (app *application) OneProductReviews(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "product_id")
	productID, err := strconv.Atoi(productId)

	if err != nil {
		app.errorJSON(w, errors.New("invalid product id"))
		return
	}

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

	productReviews, err := app.Services.ProductServices.OneProductByIDReviewsService(limit, offset, productID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, productReviews)
}

func (app *application) CreateProductReview(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "product_id")
	productID, err := strconv.Atoi(productId)

	if err != nil {
		app.errorJSON(w, errors.New("invalid product id"))
		return
	}

	var productReview models.ProductReview
	err = app.readJSON(w, r, &productReview)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	productReview.ProductID = productID
	err = app.Services.ProductServices.InsertProductReviewService(productID, productReview)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusCreated, productReview)
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
	identifier := chi.URLParam(r, "category_identifier")
	categoryID, err := strconv.Atoi(identifier)
	var category *models.Category
	if err != nil {
		category, err = app.Services.CategoryServices.OneCategoryServiceByCategoryIDOrSlug(0, identifier)
	} else {
		category, err = app.Services.CategoryServices.OneCategoryServiceByCategoryIDOrSlug(categoryID, "")
	}

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
	identifier := chi.URLParam(r, "brand_identifier")
	brandID, err := strconv.Atoi(identifier)
	var brand *models.Brand
	if err != nil {
		brand, err = app.Services.BrandServices.GetBrandServiceByIDOrSlug(0, identifier)
	} else {
		brand, err = app.Services.BrandServices.GetBrandServiceByIDOrSlug(brandID, "")
	}

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, brand)
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

// CUSTOMERS

func (app *application) AllCustomersHandler(w http.ResponseWriter, r *http.Request) {
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

	// TODO : Add filter
	//getting optional params
	search := queryParams.Get("search")
	filter := queryParams.Get("filter")
	category := queryParams.Get("category")

	optionalParams := models.OptionalQueryParams{
		Search:          search,
		Filter:          filter,
		ProductCategory: category,
	}

	offset := (page - 1) * limit

	customers, err := app.Services.CustomerServices.AllCustomersService(limit, offset, optionalParams)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, customers)
}

func (app *application) OneCustomerHandler(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "customer_id")

	customer, err := app.Services.CustomerServices.OneCustomerServiceByID(customerID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, customer)
}

func (app *application) RegisterAdminUser(w http.ResponseWriter, r *http.Request) {

	var adminUser models.User
	err := app.readJSON(w, r, &adminUser)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.Services.UserServices.RegisterAdminUserService(adminUser)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
