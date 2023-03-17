package main

import (
	"developer-network/api/comment"
	"developer-network/api/media"
	"developer-network/api/post"
	"developer-network/api/user"
	"developer-network/database"

	"github.com/gin-gonic/gin"
)

// func init() {
// 	database.Dbconnection()
// }

func main() {
	r := gin.Default()
	database.Dbconnection()
	defer database.DB.Close()

	database.CreateT()
	post.Init(r)
	comment.Init(r)
	user.Init(r)
	media.Init(r)

	r.Run()
}
