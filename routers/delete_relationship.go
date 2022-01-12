package routers

import (
	"github.com/danielpk74/tweettor/db/relationships"
	"github.com/danielpk74/tweettor/models"
	"net/http"
)

func DeleteRelationship(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Invalid ID parameter.", http.StatusBadRequest)
		return
	}

	var rel models.Relationship
	rel.UserId = IDUser
	rel.FollowedUserId = ID

	status, err := relationships.DeleteRelationship(rel)

	if err != nil || !status {
		http.Error(w, "There was an error deleting the relationship.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
