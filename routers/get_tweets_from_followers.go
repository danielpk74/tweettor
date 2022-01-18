package routers

import (
	"encoding/json"
	"github.com/danielpk74/tweettor/db/tweets"
	"net/http"
	"strconv"
)

func GetTweetsFromFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Invalid page parameter. It must be an integer.", http.StatusBadRequest)
		return
	}

	userTweets, status := tweets.GetTweetsFromFollowers(IDUser, page)
	if !status {
		http.Error(w, "There was an error getting the Tweets.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(userTweets)
}
