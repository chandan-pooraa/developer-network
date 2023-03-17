// Contains all the common methods related to database
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

func CreateT() {
	models := []interface{}{
		(*Post)(nil), (*User)(nil), (*Comment)(nil), (*Media)(nil),
	}
	for _, model := range models {

		err := DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		PanicIf(err)
	}
}
