package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saleh-ghazimoradi/GopherInn/internal/gateway/handlers"
)

type UserRoutes struct {
	userHandler *handlers.UserHandler
}

func (u *UserRoutes) UserRoute(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.POST("/users", u.userHandler.Register)
	v1.GET("/users/:id", u.userHandler.GetUserById)
	v1.GET("/users", u.userHandler.GetUsers)
	v1.PUT("/users/:id", u.userHandler.UpdateUser)
	v1.DELETE("/users/:id", u.userHandler.DeleteUser)
}

func NewUserRoutes(userHandler *handlers.UserHandler) *UserRoutes {
	return &UserRoutes{
		userHandler: userHandler,
	}
}
