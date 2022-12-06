package main

import (
	"fmt"
	"os"
	"todo/api"
	"todo/dal"

	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {

	// variaveis de ambiente
	// POSTGRES_CONNECTION_STRING - define a conexao com banco de dados postgres

	postgresConnectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if postgresConnectionString == "" {
		postgresConnectionString = "user=user password=password dbname=todo sslmode=disable"
	}

	db, err := sqlx.Connect("postgres", postgresConnectionString)
	if err != nil {
		fmt.Println("verificar se o banco esta sendo executado com o `docker-compose ps`")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("verificar se o banco esta sendo executado com o `docker-compose ps`")
		panic(err)
	}

	// REDIS_ADDRESS define hostname e porta para o servidor redis
	redisAddress := os.Getenv("REDIS_ADDRESS")
	if redisAddress == "" {
		redisAddress = "localhost:6379"
	}

	// REDIS_USERNAME username de acesso ao servidor
	redisUsername := os.Getenv("REDIS_USERNAME")
	// REDIS_PASSWORD senha do username de acesso ao servidor
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Username: redisUsername,
		Password: redisPassword,
	})

	dalInterface := dal.NewDataAccessLayerSQL(db, redisClient)
	server := api.NewServer(dalInterface)

	address := os.Getenv("LISTEN_ADDRESS")
	if address == "" {
		address = ":8080"
	}
	err = server.Start(address)
	if err != nil {
		panic(err)
	}
}
