package main

import (
	"github/richieoscar/bigshop/cmd/api"
	"log"
)

func main() {
	var server = api.NewApiServer(":8080", nil)
	var err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
