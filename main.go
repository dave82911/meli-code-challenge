package main

import (
	"meli-code-challenge/config"
	"meli-code-challenge/controller"
	"meli-code-challenge/helper"
	"meli-code-challenge/model"
	"meli-code-challenge/repository"
	"meli-code-challenge/resources"
	"meli-code-challenge/router"
	"meli-code-challenge/service"
	"net/http"

	"github.com/rs/zerolog/log"
)

// @title 	Items Service API
// @version	1.0
// @description A Items service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api/meli-challenge
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()

	db.Table("items").AutoMigrate(&model.Item{})

	// Repository
	tagsRepository := repository.NewItemsRepositoryImpl(db)

	// Resource
	meliResource := resources.NewMeliResourceImpl()

	// Service
	tagsService := service.NewItemsServiceImpl(tagsRepository, meliResource)

	// Controller
	tagsController := controller.NewItemsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
