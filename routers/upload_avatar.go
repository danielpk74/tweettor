package routers

import (
	"net/http"
	"os"
)

type MediaHandler interface {
	OpenMediaFile() error
	CopyMediaFile() error
	UpdateMediaUserProfile() (bool, error)
	GetMediaFile(ID string) (*os.File, error)
}

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	var mediaHandler = Handler{
		MediaType:         "avatar",
		DestinationFolder: "avatars/",
		Request:           r,
	}

	err := mediaHandler.OpenMediaFile()
	if err != nil {
		http.Error(w, "There was an error uploading the file", http.StatusBadRequest)
		return
	}

	err = mediaHandler.CopyMediaFile()
	if err != nil {
		http.Error(w, "There was an error copying the file: "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err := mediaHandler.UpdateMediaUserProfile()
	if err != nil || !status {
		http.Error(w, "There was an error updating the user profile: "+err.Error(), http.StatusBadRequest)
		return
	}

	//file, handler, err := r.FormFile("avatar")
	//if err != nil {
	//	http.Error(w, "There was an error reading the file", http.StatusBadRequest)
	//	return
	//}
	//
	//extension := strings.Split(handler.Filename, ".")[1]
	//avatar := "uploads/avatars/" + IDUser + "." + extension
	//
	//f, err := os.OpenFile(avatar, os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	http.Error(w, "There was an error uploading the file", http.StatusBadRequest)
	//	return
	//}
	//
	//_, err = io.Copy(f, file)
	//if err != nil {
	//	http.Error(w, "There was an error copying the file: "+err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//var user models.User
	//var status bool
	//
	//user.Avatar = IDUser + "." + extension
	//status, err = users.ModifyProfile(user, IDUser)
	//if err != nil || !status {
	//	http.Error(w, "There was an error updating the user profile: "+err.Error(), http.StatusBadRequest)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
