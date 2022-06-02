package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yutohub/channel-list/model"
)

func IndexHandler(c *gin.Context) {
	user := model.User{ID: 0, Name: "yutohub"}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"User": user,
	})
}
