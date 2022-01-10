package routers

import (
	"github.com/danielpk74/tweettor/db/tweets"
	"net/http"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Invalid ID parameter.", http.StatusBadRequest)
		return
	}

	err := tweets.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "There was an error deleting the tweet.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
