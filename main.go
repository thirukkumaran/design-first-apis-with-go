package main

import (
	"log"

	"example.com/users/Users"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a gin router
	router := gin.Default()

	// Create a server implementation
	server := Users.ServerImplementation{}

	// Register the handlers
	Users.RegisterHandlers(router, server)

	// Start the server
	log.Println("Server started on port 8080")
	router.Run()

}
