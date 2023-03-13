package main

import (
	"developer-network/api/user"
	"developer-network/database"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	database.Dbconnection()
// }

func main() {
	router := gin.Default()
	database.Dbconnection()
	defer database.DB.Close()
	model := []interface{}{
		(*(database.Post))(nil),
		(*(database.User))(nil),
	}
	database.CreateT(database.DB, model)
	user.Init(router)

	router.Run()
}
