package database

import (
	"context"
	"fmt"
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

func Dbconnection() {
	DB = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123",
		Addr:     "localhost:5432",
		Database: "postgres",
	})
	ctx := context.Background()
	err := DB.Ping(ctx)
	if err != nil {
		log.Fatal("Connection Failed!!")
	} else {
		fmt.Println("DB Connected")
	}
//	PanicIf(err)
}

func CreateT(db *pg.DB, entity []interface{}) {

	for _, model := range entity {

		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		PanicIf(err)
	}
}