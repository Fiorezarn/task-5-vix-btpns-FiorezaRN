package main

import (
	"BTPN-API-go/app"
	"BTPN-API-go/controller"
	"BTPN-API-go/exception"
	"BTPN-API-go/helper"
	"BTPN-API-go/repository"
	"BTPN-API-go/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	envErr := godotenv.Load(".env")
	helper.PanicIfError(envErr)

	webPort := os.Getenv("PORT")

	log.Printf("starting service on port %s\n", webPort)

	db := app.NewDB()
	validate := validator.New()

	repository := repository.NewImageRepositoryImpl()

	service := service.NewImageServiceImpl(repository, db, validate)

	controller := controller.NewImageControllerImpl(service)

	router := httprouter.New()

	router.POST("/api/image", controller.Create)
	router.DELETE("/api/image/:imageId", controller.Delete)
	router.GET("/api/image/:imageId", controller.FindById)

	router.PanicHandler = exception.ErrorHandler

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
