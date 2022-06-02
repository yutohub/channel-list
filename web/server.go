package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	// user := model.User{ID: 0, Name: "yutohub"}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"User": nil,
	})
}

func WillSignupHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}
