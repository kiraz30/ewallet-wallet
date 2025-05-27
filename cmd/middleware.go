package cmd

import (
	"ewallet-wallet/external"
	"ewallet-wallet/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateToken(c *gin.Context) {

	var (
		log = helpers.Logger
	)
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("Authorization header is empty")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Unauthorized", nil)
		c.Abort()
		return

	}
	tokenData, err := external.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Error(err)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Unauthorized", nil)
		c.Abort()
		return
	}
	c.Set("token", tokenData)
	c.Next()
}
