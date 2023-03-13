// Contains all the common methods related to database
package database

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

var DB *pg.DB
var m interface{}

func Dbconnection() {
	DB = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "1234",
		Addr:     "localhost:5432",
		Database: "postgres",
	})
	ctx := context.Background()
	err := DB.Ping(ctx)
	if err != nil {
		log.Fatal("Failed to connect to db")
	}
	PanicIf(err)
}

func CreateT() {
	models:= []interface{}{
		(*Post)(nil),
	}
	for _, model := range models {

		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		PanicIf(err)
	}
}
