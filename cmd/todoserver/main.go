package main

import (
	"fmt"
	"todo/api"
	"todo/dal"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sqlx.Connect("postgres", "user=user password=password dbname=todo sslmode=disable")
	if err != nil {
		fmt.Println("verificar se o banco esta sendo executado com o `docker-compose ps`")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("verificar se o banco esta sendo executado com o `docker-compose ps`")
		panic(err)
	}

	dalInterface := dal.NewDataAccessLayerSQL(db)
	server := api.NewServer(dalInterface)

	err = server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
