package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr : "localhost:8000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("aplikasi berjalan di port 8000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}