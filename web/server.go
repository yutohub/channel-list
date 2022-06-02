package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yutohub/channel-list/model"
	"github.com/yutohub/channel-list/repository"
)

func IndexHandler(c *gin.Context) {
	user := findUser(c)
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"User": user,
	})
}

func WillSignupHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}

func SignupHandler(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	if err := repository.CreateNewUser(name, password); err != nil {
		c.Error(err)
		return
	}
	user, err := repository.FindUserByName(name)
	if err != nil {
		c.Error(err)
		return
	}
	expiresAt := time.Now().Add(24 * time.Hour)
	token, err := repository.CreateNewToken(user.ID, expiresAt)
	if err != nil {
		c.Error(err)
		return
	}
	c.SetCookie("Name", "Channel_List", 86400, "/", "localhost", false, true)
	c.SetCookie("Value", token, 86400, "/", "localhost", false, true)
	c.SetCookie("Expires", expiresAt.String(), 86400, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}

func WillSigninHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.tmpl", gin.H{})
}

func SigninHandler(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	if ok, err := repository.LoginUser(name, password); err != nil || !ok {
		c.Error(fmt.Errorf("user not found or invalid password"))
		return
	}
	user, err := repository.FindUserByName(name)
	if err != nil {
		c.Error(err)
		return
	}
	expiresAt := time.Now().Add(24 * time.Hour)
	token, err := repository.CreateNewToken(user.ID, expiresAt)
	if err != nil {
		c.Error(err)
		return
	}
	c.SetCookie("Name", "Channel_List", 86400, "/", "localhost", false, true)
	c.SetCookie("Value", token, 86400, "/", "localhost", false, true)
	c.SetCookie("Expires", expiresAt.String(), 86400, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}

// Identify users by obtaining tokens from cookies
func findUser(c *gin.Context) *model.User {
	token, err := c.Cookie("Value")
	if err != nil {
		return nil
	}
	if err == nil && token != "" {
		user, _ := repository.FindUserByToken(token)
		return user
	}
	return nil
}

func SignoutHandler(c *gin.Context) {
	c.SetCookie("Name", "Channel_List", 86400, "/", "localhost", false, true)
	c.SetCookie("Value", "", 86400, "/", "localhost", false, true)
	c.SetCookie("Expires", time.Unix(0, 0).String(), 86400, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}
