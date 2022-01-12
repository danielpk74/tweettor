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

	// Users
	router.HandleFunc("/sign_up", middlew.CheckConnection(routers.SignUp)).Methods(http.MethodPost)
	router.HandleFunc("/login", middlew.CheckConnection(routers.Login)).Methods(http.MethodPost)
	router.HandleFunc("/profile", middlew.CheckConnection(middlew.ValidateJWT(routers.ViewProfile))).Methods(http.MethodGet)
	router.HandleFunc("/update_profile", middlew.CheckConnection(middlew.ValidateJWT(routers.UpdateProfile))).Methods(http.MethodPut)

	// Tweets
	router.HandleFunc("/create_tweet", middlew.CheckConnection(middlew.ValidateJWT(routers.CreateTweet))).Methods(http.MethodPost)
	router.HandleFunc("/get_tweets", middlew.CheckConnection(middlew.ValidateJWT(routers.GetTweets))).Methods(http.MethodGet)
	router.HandleFunc("/delete_tweet", middlew.CheckConnection(middlew.ValidateJWT(routers.DeleteTweet))).Methods(http.MethodDelete)

	// Media
	router.HandleFunc("/upload_avatar", middlew.CheckConnection(middlew.ValidateJWT(routers.UploadAvatar))).Methods(http.MethodPost)
	router.HandleFunc("/upload_banner", middlew.CheckConnection(middlew.ValidateJWT(routers.UploadBanner))).Methods(http.MethodPost)
	router.HandleFunc("/get_avatar", middlew.CheckConnection(routers.GetAvatar)).Methods(http.MethodGet)
	router.HandleFunc("/get_banner", middlew.CheckConnection(routers.GetBanner)).Methods(http.MethodGet)

	// Relationships
	router.HandleFunc("/create_relationship", middlew.CheckConnection(middlew.ValidateJWT(routers.CreateRelationship))).Methods(http.MethodPost)
	router.HandleFunc("/delete_relationship", middlew.CheckConnection(middlew.ValidateJWT(routers.DeleteRelationship))).Methods(http.MethodDelete)

	PORT := os.Getenv("PORT")
	if (PORT) == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
