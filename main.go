package main

import (
	"log"

	"github.com/puang59/internode/routes"
)

var port string = "8080"

func main() {
	router := routes.SetupRouter()

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
