package routers

import (
	"encoding/json"
	"github.com/danielpk74/tweettor/db/tweets"
	"net/http"
	"strconv"
)

var TweetsPerPage = 20

func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 0 {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	userTweets, status := tweets.GetTweetsPerUser(IDUser, int64(page), int64(TweetsPerPage))
	if !status {
		http.Error(w, "There was an error getting Tweets.", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userTweets)
}
