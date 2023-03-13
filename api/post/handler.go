package post

import (
	"developer-network/api"
	"developer-network/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewPost(c *gin.Context) {
	// db := database.Dbconnection()
	var newPost database.Post
	err := c.BindJSON(&newPost)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error in Binding",
		})
		api.PanicIf(err)
	}
	_, err = database.DB.Model(&newPost).Insert()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error in INsertion",
		})
	}
	api.PanicIf(err)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Post created successfull"})
}
