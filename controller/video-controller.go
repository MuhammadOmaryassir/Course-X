package controller

import (
	"encoding/json"
	"net/http"

	"../entity"
	"../errors"
	"../service"
)

type controller struct{}

var videoService service.VideoService

// VideoController video controller interface
type VideoController interface {
	GetVideos(response http.ResponseWriter, request *http.Request)
	AddVideo(response http.ResponseWriter, request *http.Request)
}

// NewVideoController constructor function
func NewVideoController(service service.VideoService) VideoController {
	videoService = service
	return &controller{}
}

func (*controller) GetVideos(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	videos, err := videoService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the videos", Code: 500})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(videos)
}

func (*controller) AddVideo(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var video entity.Video
	err := json.NewDecoder(request.Body).Decode(&video)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data", Code: 500})
		return
	}
	err1 := videoService.Validate(&video)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error(), Code: 500})
		return
	}

	result, err2 := videoService.Create(&video)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the video", Code: 500})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
