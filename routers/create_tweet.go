package routers

import (
	"encoding/json"
	"github.com/danielpk74/tweettor/db/tweets"
	"github.com/danielpk74/tweettor/models"
	"net/http"
	"time"
)

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	tweet := models.Tweet{
		UserID: IDUser,
		Date:   time.Now(),
	}

	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		http.Error(w, "Error decoding the tweet.", http.StatusBadRequest)
		return
	}

	_, status, err := tweets.CreateTweet(tweet)
	if err != nil || !status {
		http.Error(w, "Error creating the tweet.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
