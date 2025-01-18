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
	"time"
)

const port = 8080

type application struct {
	DB       repository.DatabaseRepo
	Services struct {
		ProductServices     repository.ProductRepo
		CategoryServices    repository.CategoryRepo
		BrandServices       repository.BrandRepo
		SubCategoryServices repository.SubCategoryRepo
		CustomerServices    repository.CustomerRepo
		UserServices        repository.UserRepo
	}
	DSN          string
	Domain       string
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	auth         Auth
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
	app.JWTSecret = os.Getenv("JWTSecret")
	app.JWTIssuer = os.Getenv("JWTIssuer")
	app.JWTAudience = os.Getenv("JWTAudience")
	app.CookieDomain = os.Getenv("CookieDomain")
	app.Domain = os.Getenv("Domain")

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "refresh_token",
		CookieDomain:  app.CookieDomain,
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	app.Services.ProductServices = &services.ProductServices{
		PostgresDBRepo: &dbrepo.PostgresDBRepo{DB: conn},
	}
	app.Services.UserServices = &services.UserService{
		PostgresDBRepo: &dbrepo.PostgresDBRepo{DB: conn},
	}
	app.Services.CustomerServices = &services.CustomerServices{
		PostgresDBRepo: &dbrepo.PostgresDBRepo{DB: conn},
	}
	app.Services.CategoryServices = &services.CategoryServices{
		PostgresDBRepo: &dbrepo.PostgresDBRepo{DB: conn},
	}
	app.Services.BrandServices = &services.BrandServices{
		PostgresDBRepo: &dbrepo.PostgresDBRepo{DB: conn},
	}
	app.Services.SubCategoryServices = &services.SubCategoryServices{
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
