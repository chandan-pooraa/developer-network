package main

import (
	"developer-network/api/comment"
	"developer-network/database"

	//"go/doc/comment"

	// "fmt"
	// "go/doc/comment"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	database.Dbconnection()
// }

// func main() {
// 	router := gin.Default()
// 	database.Dbconnection()
// 	defer database.DB.Close()
// 	model := []interface{}{
// 		(*(database.Post))(nil),
// 		(*(database.User))(nil),
// 		(*(database.Comment))(nil),
// 	}
// 	database.CreateT(database.DB, model)
// 	comment.Init(router)

// 	router.Run()
// }

func main() {
	r := gin.Default()
	database.Dbconnection()
	defer database.DB.Close()

	database.CreateT()

	comment.Init(r)

	r.Run()
}
