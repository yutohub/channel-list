package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yutohub/channel-list/web"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", web.IndexHandler)
	router.Run("localhost:8080")
}
