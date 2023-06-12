package router

import (
	"os"
	roomController "sirkelin/backend/app/room/controller"
	userController "sirkelin/backend/app/user/controller"
	"sirkelin/backend/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	middleware *middlewares.Middleware
	user       *userController.UserController
	room       *roomController.RoomController
}

func NewRouter(middleware *middlewares.Middleware, auth *userController.UserController, room *roomController.RoomController) *Router {
	return &Router{
		middleware: middleware,
		user:       auth,
		room:       room,
	}
}

func (router *Router) Handle() *gin.Engine {
	handler := gin.Default()
	gin.SetMode(os.Getenv("APP_MODE"))

	authGroup := handler.Group("/auth")
	{
		authGroup.POST("/sign-in", router.user.SignIn)
		authGroup.POST("/sign-out", router.user.SignOut)
	}

	roomGroup := handler.Group("/room")
	{
		roomGroup.Use(router.middleware.AuthenticatedUser())
		roomGroup.GET("/", router.room.GetRooms)
		roomGroup.POST("/", router.room.CreateRoom)
	}

	userGroup := handler.Group("/user")
	{
		userGroup.Use(router.middleware.AuthenticatedUser())
		userGroup.GET("/", router.user.GetAll)
		userGroup.GET("/:id", router.user.GetByID)
	}

	return handler
}
