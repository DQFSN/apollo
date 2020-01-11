package router

import (
	"net/http"

	"github.com/chalvern/apollo/app/controllers/home"
	"github.com/gin-gonic/gin"
)

// pong for ping
func pong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// 定义 router
func routerInit() {
	get("ping_pong", "/ping", pong)
	get("home_page", "/", home.Index)
	get("about", "/about", pong)

	// account
	get("signup", "/signup", pong)
	get("signin", "/signin", pong)
	get("signout", "/signout", pong)

	// user
	get("user_info", "/user/info", pong)
}
