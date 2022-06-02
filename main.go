package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yutohub/channel-list/web"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", web.IndexHandler)
	router.GET("/signup", web.WillSignupHandler)
	router.POST("/signup", web.SignupHandler)
	router.Run("localhost:8080")
}
