package media

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"message": "Server is Running"})
	// })

	router.GET("/media", GetAllMedia)
	router.GET("/media/:id", GetMediaById)

	router.POST("/newmedia", CreateNewMedia)

	router.DELETE("/media/:id", DeleteMediaById)

}
