package app

import (
	"github.com/FreeCodeUserJack/bookstore_oauth/src/domain/access_token"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/http"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepo()
	atService := access_token.NewService(dbRepository)
	handler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", handler.GetById)

	router.Run(":8080")
}