package routes

import (
	"github.com/gin-gonic/gin"
	"test/Repository"
)

func UserRoutes(route *gin.Engine){
	v1:=route.Group("/siswa")
	{
	v1.GET("/", UserRepository.GetUser)
	v1.GET("/:id",UserRepository.GetUserById)
	// v1.GET("/login",UserRepository.Login)
	v1.POST("/", UserRepository.AddUser)
	v1.PUT("/:id",UserRepository.UpdateUser)
	v1.DELETE("/:id",UserRepository.DeleteUser)
	}
}