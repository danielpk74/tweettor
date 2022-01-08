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

	// Users
	router.HandleFunc("/sign_up", middlew.CheckConnection(routers.SignUp)).Methods(http.MethodPost)
	router.HandleFunc("/login", middlew.CheckConnection(routers.Login)).Methods(http.MethodPost)
	router.HandleFunc("/profile", middlew.CheckConnection(middlew.ValidateJWT(routers.ViewProfile))).Methods(http.MethodGet)
	router.HandleFunc("/update_profile", middlew.CheckConnection(middlew.ValidateJWT(routers.UpdateProfile))).Methods(http.MethodPut)

	// Tweets
	router.HandleFunc("/create_tweet", middlew.CheckConnection(middlew.ValidateJWT(routers.CreateTweet))).Methods(http.MethodPost)
	router.HandleFunc("/get_tweets", middlew.CheckConnection(middlew.ValidateJWT(routers.GetTweets))).Methods(http.MethodGet)

	PORT := os.Getenv("PORT")
	if (PORT) == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
