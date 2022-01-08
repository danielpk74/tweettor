package routers

import (
	"encoding/json"
	"github.com/danielpk74/tweettor/db/users"
	"net/http"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := users.FindProfile(ID)
	if err != nil {
		http.Error(w, "Error finding the profile"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(profile)
}
