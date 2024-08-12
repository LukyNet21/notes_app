package main

import (
	"fmt"
	"net/http"
	"notes_server/config"
	"notes_server/database"
	"notes_server/webhost"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.LoadConfig()

	r := mux.NewRouter()

	webhost.SetupRoutes(r)

	database.Connect()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4321"}, // Update this with your frontend's URL
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	// Wrap the router with the CORS middleware
	handler := c.Handler(r)

	fmt.Printf("Starting server on http://%s:%d.\n", config.Host, config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handler)
}
