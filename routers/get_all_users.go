package routers

import (
	"encoding/json"
	"github.com/danielpk74/tweettor/db/users"
	"net/http"
	"strconv"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	requestedPage := r.URL.Query().Get("page")
	searchText := r.URL.Query().Get("search_text")
	searchType := r.URL.Query().Get("search_type")

	page, err := strconv.Atoi(requestedPage)
	if err != nil || page <= 0 {
		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
		return
	}

	allUsers, status := users.GetAllUsers(IDUser, int64(page), searchText, searchType)
	if !status {
		http.Error(w, "There was an error getting Tweets.", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(allUsers)
}
