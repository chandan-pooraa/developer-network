package post

import (
	"developer-network/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNewPost(c *gin.Context) {
	// db := database.Dbconnection()
	var newPost database.Post
	err := c.BindJSON(&newPost)
	if BadRequest(err, c) != nil {
		return
	}
	_, err = database.DB.Model(&newPost).Insert()
	if BadRequest(err, c) != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Post created successfull"})
}

//Funtion to show all the post
func ShowPosts(c *gin.Context) {
	var posts []database.Post
	err := database.DB.Model(&posts).Column().Select()
	if BadRequest(err, c) != nil {
		return
	}
	fmt.Println(posts)
	c.IndentedJSON(http.StatusOK, posts)
}

//
func GetPostById(c *gin.Context) {
	id := c.Param("id")
	p, err := SearchInPostsUid(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, p)
}

func UpdatePostById(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.ParseFloat(id, 64)
	if BadRequest(err, c) != nil {
		return
	}
	updatePost := &database.Post{Id: int(idint)}
	err = database.DB.Model(updatePost).Select()
	if BadRequest(err, c) != nil {
		return
	}

	err = c.Bind(updatePost)
	if BadRequest(err, c) != nil {
		return
	}
	_, err = database.DB.Model(updatePost).Update()
	if BadRequest(err, c) != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func DeletePost(c *gin.Context) {
	var delpost database.Post
	id := c.Param("id")
	idint, err := strconv.ParseFloat(id, 64)
	BadRequest(err, c)
	_, err = SearchInPostsPid(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}
	_, err = database.DB.Model(&delpost).Where("id = ?", idint).Delete()
	if BadRequest(err, c) != nil {
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Post deleted"})
}
