package routers

import (
	"errors"
	"github.com/danielpk74/tweettor/db/users"
	"github.com/danielpk74/tweettor/models"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type Handler struct {
	DestinationFolder string
	file              *os.File
	mediaExtension    string
	mediaFile         multipart.File
	MediaType         string
	Request           *http.Request
}

func (h *Handler) OpenMediaFile() error {
	var handler *multipart.FileHeader
	var err error
	h.mediaFile, handler, err = h.Request.FormFile(h.MediaType)
	if err != nil {
		return err
	}

	h.mediaExtension = strings.Split(handler.Filename, ".")[1]
	destination := "uploads/" + h.DestinationFolder + IDUser + "." + h.mediaExtension

	h.file, err = os.OpenFile(destination, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) CopyMediaFile() error {
	_, err := io.Copy(h.file, h.mediaFile)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) UpdateMediaUserProfile() (bool, error) {

	var user models.User
	var status bool

	userMedia := IDUser + "." + h.mediaExtension
	if h.MediaType == "avatar" {
		user.Avatar = userMedia
	}

	if h.MediaType == "banner" {
		user.Banner = userMedia
	}

	status, err := users.ModifyProfile(user, IDUser)
	if err != nil || !status {
		return false, err
	}

	return true, nil
}

func (h Handler) GetMediaFile(ID string) (*os.File, error) {
	if len(ID) < 1 {
		return nil, errors.New("Invalid ID parameter")
	}

	profile, err := users.FindProfile(ID)
	if err != nil {
		return nil, errors.New("User not found")
	}

	destinationFolder := "uploads/" + h.DestinationFolder
	if h.MediaType == "avatar" {
		destinationFolder += profile.Avatar
	}

	if h.MediaType == "banner" {
		destinationFolder += profile.Banner
	}

	openFile, err := os.Open(destinationFolder)
	if err != nil {
		return nil, errors.New("Image not found")
	}

	return openFile, nil
}
