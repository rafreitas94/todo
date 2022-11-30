package main

import (
	"fmt"
	"todo/api"
	"todo/dal"

	"github.com/go-redis/redis/v9"
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

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	dalInterface := dal.NewDataAccessLayerSQL(db, redisClient)
	server := api.NewServer(dalInterface)

	err = server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
