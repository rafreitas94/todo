package main

import (
	"todo/api"
	"todo/dal"
)

func main() {
	dalInterface := dal.NewDataAccessLayer()
	server := api.NewServer(dalInterface)

	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
