package app

import (
	"github.com/FreeCodeUserJack/bookstore_oauth/src/http"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/repository/db"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/repository/rest"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepo()
	usersRepository := rest.NewRepository()

	atService := access_token.NewService(usersRepository, dbRepository)
	handler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", handler.GetById)
	router.POST("/oauth/access_token", handler.Create)

	router.Run(":8080")
}