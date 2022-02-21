package main

import (
	"log"
	"net/http"

	"github.com/migalpha/kentech-films/config"
	handler "github.com/migalpha/kentech-films/http"
	"github.com/migalpha/kentech-films/middleware"
	"github.com/migalpha/kentech-films/postgres"
	"github.com/migalpha/kentech-films/redis"

	"github.com/gin-gonic/gin"
)

// setupServer returns a Gin server ready to rise up with all the available endpoints.
func setupServer(app application) *http.Server {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/films/health"),
		gin.Recovery(),
		// otelgin.Middleware(config.Config.APMAppName),
	)

	// Users endpoints
	userRepo := postgres.UsersRepo{
		DB: app.postgres,
	}

	logoutRepo := redis.TokenRepo{
		DB: app.redis,
	}

	handlerRegisterUser := handler.RegisterUserHandler{Repo: userRepo}
	router.POST("register", handlerRegisterUser.ServeHTTP)

	handlerLogin := handler.LoginHandler{Repo: userRepo}
	router.POST("login", handlerLogin.ServeHTTP)

	v1 := router.Group("api/v1")
	v1.Use(middleware.CheckJWT(logoutRepo))
	// Films endpoints
	filmRepo := postgres.FilmRepo{
		DB: app.postgres,
	}

	handlerCreateFilm := handler.CreateFilmHandler{Repo: filmRepo}
	v1.POST("films", handlerCreateFilm.ServeHTTP)

	handlerGetFilmDetails := handler.GetFilmDetailsHandler{Repo: filmRepo}
	v1.GET("films/:id", handlerGetFilmDetails.ServeHTTP)

	handlerGetFilms := handler.GetFilmsHandler{Repo: filmRepo}
	v1.GET("films", handlerGetFilms.ServeHTTP)

	handlerUploadFilm := handler.UpdateFilmHandler{Provider: filmRepo, Updater: filmRepo}
	v1.PUT("films/:id", handlerUploadFilm.ServeHTTP)

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
	favoriteRepo := postgres.FavouriteRepo{
		DB: app.postgres,
	}

	handlerAddFavourite := handler.AddFavouriteHandler{Saver: favoriteRepo, Provider: filmRepo}
	v1.POST("favourites", handlerAddFavourite.ServeHTTP)

	handlerRemoveFavourite := handler.RemoveFavoriteHandler{Destroyer: favoriteRepo, Provider: filmRepo}
	v1.DELETE("favourites/:id", handlerRemoveFavourite.ServeHTTP)

	// Logout
	handlerLogout := handler.LogoutHandler{Repo: logoutRepo}
	v1.POST("logout", handlerLogout.ServeHTTP)

	// Health check of the app.
	router.GET("health", handler.HealthCheck)

	// docs.SwaggerInfo.BasePath = "internal-microservices.dev.rappi.com/api/maps"
	// url := ginSwagger.URL("api/maps/swagger/doc.json")
	// router.GET("api/maps/swagger/*any", middlewares.IsEnabledSwagger(), ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	log.Printf("Server listen in %s:%s", config.Commons().Host, config.Commons().Port)
	return &http.Server{
		Addr:              config.HTTP().Address,
		Handler:           router,
		ReadTimeout:       config.HTTP().ReadTimeout,
		ReadHeaderTimeout: config.HTTP().ReadHeaderTimeout,
		WriteTimeout:      config.HTTP().WriteTimeout,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}
}
