package web

import (
	"fmt"
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

func SignupHandler(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	fmt.Println(name)
	fmt.Println(password)
}

func WillSigninHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.tmpl", gin.H{})
}

func SigninHandler(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	fmt.Println(name)
	fmt.Println(password)
}
