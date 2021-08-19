package middlew

import (
	"github.com/danielpk74/tweettor/db"
	"net/http"
)

// CheckConnection Check if the database is available.
func CheckConnection(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.Conn.CheckConnection() {
			http.Error(w, "MongoDB is not available", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
