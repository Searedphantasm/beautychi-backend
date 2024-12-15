package main

import (
	"log"
	"net/http"
)

func (app *application) AllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := app.Services.ProductServices.AllProductsService()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(products)
	_ = app.writeJSON(w, http.StatusOK, products)
}
