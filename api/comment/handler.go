package comment

import (
	"developer-network/api"
	"developer-network/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewComment(c *gin.Context) {
	var newComment database.Comment
	err := c.ShouldBindJSON(&newComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in Binding",
		})
		api.PanicIf(err)
	}
	_, err = database.DB.Model(&newComment).Insert()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error in creating a newComment",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Comment created Successfully",
	})
	
 }

 func ReadComments(c *gin.Context) {
	var commentRows []database.Comment
	c.IndentedJSON(http.StatusOK, commentRows)
 }


 func ReadCommentbyId(c *gin.Context) {
	var newComment database.Comment
	id := c.Param("id")
	err := database.DB.Model(&newComment, id).WherePK().Select()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "comment not found",
		})
	}
	fmt.Println(err)
	c.JSON(http.StatusOK, gin.H{
		"newComment": newComment,
	})
 }


 func UpdateComment(c *gin.Context) {
	var newComment database.Comment
	//id := c.Param("id")
	err := c.ShouldBind(&newComment)
   
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
  
	var updateComment database.Comment
	err = database.DB.Model(&updateComment).WherePK().Select()
  
	if err!= nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "newComment not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"newComment": "Comment Updated Successfully!!",
	})
 }
 
 
 func DeleteComment(c *gin.Context) {
	var newComment database.Comment
	id := c.Param("id")
	err := database.DB.Model(&newComment, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "newComment not found",
		})
		return
	}
	database.DB.Model(&newComment).WherePK().Delete()
	c.JSON(http.StatusOK, gin.H{
		"message": "newComment deleted successfully",
	})
 }
 