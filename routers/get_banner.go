package routers

import (
	"github.com/danielpk74/tweettor/db/users"
	"io"
	"net/http"
	"os"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("ID")
	//handler := Handler{}
	//openFile, err := handler.GetMediaFile(ID)
	//if err != nil {
	//	http.Error(w, "Error copying file" + err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//_, err = io.Copy(w, openFile)
	//if err != nil {
	//	http.Error(w, "Error copying file", http.StatusBadRequest)
	//	return
	//}

	if len(ID) < 1 {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := users.FindProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copying file", http.StatusBadRequest)
		return
	}
}
