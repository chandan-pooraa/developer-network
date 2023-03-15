package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is Running__post"})
	})
	post := router.Group("/post")
	{
		post.POST("/new", CreateNewPost)
		post.GET("/showposts", ShowPosts)
		post.GET("/show/:id", GetPostById)
		post.PATCH("/update/:id", UpdatePostById)
		post.DELETE("delete/:id", DeletePost)
	}
}
