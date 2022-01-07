package middlew

import (
	"github.com/danielpk74/tweettor/routers"
	"net/http"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token Error!"+err.Error(), http.StatusBadRequest)
		}

		next.ServeHTTP(w, r)
	}
}
