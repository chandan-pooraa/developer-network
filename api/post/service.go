package post

import (
	"developer-network/database"
	"net/http"

	"github.com/gin-gonic/gin"
)


func BadRequest(err error, c *gin.Context) error{
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
	}
	return err
}

func SearchInPostsUid(id string)(*database.Post,error){
	var post database.Post
	err := database.DB.Model(&post).Column("title", "content").Where("user_id = ?", id).Select()
	return &post,err
}
func SearchInPostsPid(id string)(*database.Post,error){
	var post database.Post
	err := database.DB.Model(&post).Column("title", "content").Where("id = ?", id).Select()
	return &post,err
}