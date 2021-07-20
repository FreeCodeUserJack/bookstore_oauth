package http

import (
	"net/http"
	"strings"

	"github.com/FreeCodeUserJack/bookstore_oauth/src/domain/access_token"
	"github.com/FreeCodeUserJack/bookstore_oauth/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
	// UpdateExpirationTime(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(s access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: s,
	}
}

func (a *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := a.service.GetById(accessTokenId)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (a *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body", "bad request")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := a.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}