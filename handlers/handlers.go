package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"

	"github.com/danielpk74/tweettor/middlew"
	"github.com/danielpk74/tweettor/routers"
)

// Handlers Web Server
func Handlers() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/sign_up", middlew.CheckConnection(routers.SignUp)).Methods("POST")

	PORT := os.Getenv("PORT")
	if (PORT) == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}