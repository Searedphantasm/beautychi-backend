package main

import (
	"fmt"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/dbrepo"
	"github.com/PAPAvision-co/beautychi-backend.git/internals/repository/services"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const port = 8080

type application struct {
	DB       repository.DatabaseRepo
	Services struct {
		ProductServices  repository.ProductRepo
		CategoryServices repository.CategoryRepo
		BrandServices    repository.BrandRepo
	}
	DSN    string
	Domain string
}

func main() {
	// set application config
	var app application

	// read from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.DSN = os.Getenv("DSN")

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	app.Services.ProductServices = &services.ProductServices{
		PostgresDBRepo: &dbrepo.PostgresDBRepo{DB: conn},
	}
	app.Services.CategoryServices = &services.CategoryServices{
		PostgresDBRepo: &dbrepo.PostgresDBRepo{DB: conn},
	}
	app.Services.BrandServices = &services.BrandServices{
		PostgresDBRepo: &dbrepo.PostgresDBRepo{DB: conn},
	}

	defer app.DB.Connection().Close()

	log.Println("Starting server on port", port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
