package user

import (
	"net/http"
	"strconv"

	"developer-network/database"

	"github.com/gin-gonic/gin"
)

func Getusers(c *gin.Context) {
	var userRows []database.User
	c.IndentedJSON(http.StatusOK, userRows)
}
func CreateNewuser(c *gin.Context) {

	var newUser database.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = database.DB.Model(&newUser).Insert()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "New user created successfully"})

}
func GetuserbyID(c *gin.Context) {

	id := c.Param("id")
	ids, _ := strconv.ParseFloat(id, 64)
	userget := &database.User{Id: ids}
	err := database.DB.Model(userget).WherePK().Select()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userget)

}
func Updateuser(c *gin.Context) {
	id := c.Param("id")
	idFloat, err := strconv.ParseFloat(id, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user from the database
	updateUser := &database.User{Id: idFloat}
	err = database.DB.Model(updateUser).WherePK().Select()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the user fields
	err = c.Bind(updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated user back to the database
	_, err = database.DB.Model(updateUser).WherePK().Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

}
func Deleteuser(c *gin.Context) {
	id := c.Param("id")
	idFloat, err := strconv.ParseFloat(id, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delUser := &database.User{Id: idFloat}

	_, err = database.DB.Model(delUser).WherePK().Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
