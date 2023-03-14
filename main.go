package main

import (
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

	r.Run()
}
