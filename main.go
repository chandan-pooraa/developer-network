package main

import (
	"developer-network/api/comment"
	"developer-network/database"
	"github.com/gin-gonic/gin"
)

// func init() {
// 	database.Dbconnection()
// }

// func main() {
// 	r := gin.Default()
// 	database.Dbconnection()
// 	defer database.DB.Close()

// 	database.CreateT()

// 	r.Run()
// }

func main() {
	r := gin.Default()
	database.Dbconnection()
	defer database.DB.Close()
 
	model := []interface{}{
		 (*(database.Post))(nil),
		 (*(database.User))(nil),
		 (*(database.Comment))(nil),
	 }
	 database.CreateT(database.DB, model)
 
	comment.Init(r)
	
	r.Run("localhost:5000")
 }
