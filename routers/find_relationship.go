package routers

import (
	"encoding/json"
	"github.com/danielpk74/tweettor/db/relationships"
	"github.com/danielpk74/tweettor/models"
	"net/http"
)

func FindRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relationship
	t.UserId = IDUser
	t.FollowedUserId = ID

	status, err := relationships.FindRelationship(t)
	if err != nil || status == false {
		t.Status = false
	} else {
		t.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(t)
}
