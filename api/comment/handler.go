package comment

import (
	"developer-network/api"
	"developer-network/database"
	"fmt"
	"net/http"
	"strconv"
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
	// id := c.Param("id")
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        // ... handle error
        // panic(err)
		fmt.Println("Error in Fetching data!!")
    }
	newComment.Id = id
	err = database.DB.Model(&newComment).WherePK().Select()
	fmt.Println(newComment)
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
	id := c.Param("id")
	idint, err := strconv.ParseFloat(id, 64)
	updateComment := &database.Comment{Id: int(idint)}
	err = database.DB.Model(&updateComment).Select()

	//err = database.DB.Model(&newComment).Where("id = ?", idint).Select()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "newComment not found",
		})
		return
	}
	
	err=c.Bind(updateComment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "newComment Binding error",
		})
		return
	}
	_, err = database.DB.Model(updateComment).Update()


	c.JSON(http.StatusOK, gin.H{
		"newComment": "Comment Updated Successfully!!",
	})

 }
 
 
 func DeleteComment(c *gin.Context) {
	var newComment database.Comment
	id := c.Param("id")
	idint, err := strconv.ParseFloat(id, 64)

	err = database.DB.Model(&newComment).Where("id = ?", idint).Select()
	
	//err = database.DB.Model(&newComment).Column("title", "content").Where("id = ?", id).Select()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "newComment not found",
		})
		// fmt.Println(id)
		return
	}
	database.DB.Model(&newComment).WherePK().Delete()
	c.JSON(http.StatusOK, gin.H{
		"message": "newComment deleted successfully",
	})
}