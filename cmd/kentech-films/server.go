package main

import (
	"log"
	"net/http"

	"github.com/migalpha/kentech-films/config"
	_ "github.com/migalpha/kentech-films/docs"
	handler "github.com/migalpha/kentech-films/http"
	"github.com/migalpha/kentech-films/middleware"
	"github.com/migalpha/kentech-films/postgres"
	"github.com/migalpha/kentech-films/redis"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Kentech-Films
// @version 1.0.0
// @description This API provides endpoints to manage films and register users.
// @description [Read me](https://github.com/migalpha/kentech-films)
// @termsOfService http://swagger.io/terms/
// @schemes http
// @host localhost:8000
// @BasePath /api/v1
func setupServer(app application) *http.Server {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "health"),
		gin.Recovery(),
		middleware.CORS(),
	)

	//Loading repos
	favouriteRepo := postgres.NewFavouriteRepository(app.postgres)
	filmRepo := postgres.NewFilmRepository(app.postgres)
	logoutRepo := redis.NewTokenRepository(app.redis)
	userRepo := postgres.NewUserRepository(app.postgres)

	// Users endpoints
	v1 := router.Group("api/v1")
	handlerRegisterUser := handler.RegisterUserHandler{Repo: userRepo}
	v1.POST("register", handlerRegisterUser.ServeHTTP)

	handlerLogin := handler.LoginHandler{Repo: userRepo}
	v1.POST("login", handlerLogin.ServeHTTP)

	v1.Use(middleware.CheckJWT(logoutRepo))

	// Films endpoints
	handlerCreateFilm := handler.CreateFilmHandler{Repo: filmRepo}
	v1.POST("films", handlerCreateFilm.ServeHTTP)

	handlerGetFilmDetails := handler.GetFilmDetailsHandler{Repo: filmRepo}
	v1.GET("films/:id", handlerGetFilmDetails.ServeHTTP)

	handlerGetFilms := handler.GetFilmsHandler{Repo: filmRepo}
	v1.GET("films", handlerGetFilms.ServeHTTP)

	handlerUploadFilm := handler.UpdateFilmHandler{Provider: filmRepo, Updater: filmRepo}
	v1.PATCH("films/:id", handlerUploadFilm.ServeHTTP)

	handlerDeleteFilm := handler.DeleteFilmHandler{
		Provider:  filmRepo,
		Destroyer: filmRepo,
	}
	v1.DELETE("films/:id", handlerDeleteFilm.ServeHTTP)

	handlerExportCSV := handler.ExportCSVHandler{Repo: filmRepo}
	v1.GET("csv/films", handlerExportCSV.ServeHTTP)

	handlerImportCSV := handler.ImportCSVHandler{Repo: filmRepo}
	v1.POST("csv/films", handlerImportCSV.ServeHTTP)

	// Favourites
	handlerAddFavourite := handler.AddFavouriteHandler{Saver: favouriteRepo, Provider: filmRepo}
	v1.POST("favourites", handlerAddFavourite.ServeHTTP)

	handlerRemoveFavourite := handler.RemoveFavouriteHandler{Destroyer: favouriteRepo, Provider: filmRepo}
	v1.DELETE("favourites/:id", handlerRemoveFavourite.ServeHTTP)

	// Logout
	handlerLogout := handler.LogoutHandler{Repo: logoutRepo}
	v1.POST("logout", handlerLogout.ServeHTTP)

	// Health check of the app.
	router.GET("health", handler.HealthCheck)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server listen in %s", config.HTTP().Address)
	return &http.Server{
		Addr:              config.HTTP().Address,
		Handler:           router,
		ReadTimeout:       config.HTTP().ReadTimeout,
		ReadHeaderTimeout: config.HTTP().ReadHeaderTimeout,
		WriteTimeout:      config.HTTP().WriteTimeout,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}
}
