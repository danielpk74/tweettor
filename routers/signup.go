package routers

import (
	"encoding/json"
	"net/http"

	"github.com/danielpk74/tweettor/db/users"
	"github.com/danielpk74/tweettor/models"
)

// SignUp create a new record for a user in the database.
func SignUp(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid Request: "+err.Error(), http.StatusBadRequest)
	}

	validateFields(&u, w, r)
	_, exists, _ := users.UserAlreadyExists(u.Email)
	if exists {
		http.Error(w, "The user already exists: ", http.StatusBadRequest)
		return
	}

	_, status, err := users.CreateUser(&u)
	if err != nil {
		http.Error(w, "Error creating a new user: "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "User could not be created.: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func validateFields(u *models.User, w http.ResponseWriter, r *http.Request) {
	if len(u.Email) == 0 {
		http.Error(w, "Invalid Email: ", http.StatusBadRequest)
		return
	}

	if len(u.Password) == 0 {
		http.Error(w, "Invalid Password: Password is required", http.StatusBadRequest)
		return
	}

	if len(u.Password) < 7 {
		http.Error(w, "Invalid Password: Password must have at least 7 characters ", http.StatusBadRequest)
		return
	}
}
