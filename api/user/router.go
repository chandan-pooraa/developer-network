package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is Running__user"})
	})

	router.GET("/users", Getusers)
	router.POST("/newUser", CreateNewuser)
	router.GET("/user/:id", GetuserbyID)
	router.PUT("/user/:id", Updateuser)
	router.DELETE("/user/:id", Deleteuser)

}