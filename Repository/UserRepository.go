package UserRepository

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"test/models"
	"strconv"
)


func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}


func GetUser(c *gin.Context){
	user,err:=models.GetUser()
	CheckErr(err)
	if user==nil {
		c.JSON(http.StatusBadRequest,gin.H{"message": "data tidak ditemukan"})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{"message": user})

	}

}
func GetUserById(c *gin.Context){
	id:=c.Param("id")
	user,err:=models.GetUserById(id)
	CheckErr(err)
	if user.Username=="" {
		c.JSON(http.StatusBadRequest,gin.H{"message": "Username tidak ditemukan"})
	}else{
		c.JSON(http.StatusOK,gin.H{"data":user})
	}

}
func AddUser(c *gin.Context){
	var user models.User

	if err:=c.ShouldBindJSON(&user);err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	success,err:=models.AddUser(user)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	}else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}


	
}
func UpdateUser(c *gin.Context){

	var json models.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := models.UpdateUser(json, personId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}



}
func DeleteUser(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid Id"})
	}
	success,err:=models.DeleteUser(id)
	if success {
		c.JSON(http.StatusOK,gin.H{"message":"Success"})
	}else{
		c.JSON(http.StatusBadRequest,gin.H{"error":err})
	}


}



