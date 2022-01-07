package routers

import (
	"encoding/json"
	"github.com/danielpk74/tweettor/db/users"
	"github.com/danielpk74/tweettor/jwt"
	"github.com/danielpk74/tweettor/models"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid Request"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	user, exists := users.Login(u.Email, u.Password)
	if !exists {
		http.Error(w, "Invalid User or Password.", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "JWT error"+err.Error(), http.StatusBadRequest)
		return
	}

	resp := users.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   resp.Token,
		Expires: expirationTime,
	})
}
