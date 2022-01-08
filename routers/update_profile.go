package routers

import (
	"encoding/json"
	"github.com/danielpk74/tweettor/db/users"
	"github.com/danielpk74/tweettor/models"
	"net/http"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid profile data: "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err := users.ModifyProfile(user, IDUser)
	if err != nil || !status {
		http.Error(w, "There was an error updating the profile data: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
