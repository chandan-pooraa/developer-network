package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is Running..."})
	})
	
		router.GET("/comments", ReadComments)
		router.GET("/getcomments/:id", ReadCommentbyId)
		router.POST("/newcomment", CreateNewComment)
		router.PATCH("/upcomment/:id", UpdateComment)
		router.DELETE("/delcomment/:id", DeleteComment)
	
	
}