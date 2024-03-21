package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ryosantouchh/golang-todo/auth"
)

func NewRouterHandler(route *gin.Engine) {
	// route.GET("/tokenz", auth.AccessToken) // normal version
	route.GET("/tokenz", auth.AccessToken([]byte("+++signature+++"))) // middleware version
}
