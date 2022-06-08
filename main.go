package main

import(
	"github.com/gin-gonic/gin"
	"test/Repository"
	"test/models"
	"test/routes"
)


func main(){
	r:=gin.Default()
	routes.UserRoutes(r)
	
	err := models.ConnectDB()
UserRepository.CheckErr(err)
r.Run()

}