package service

import "github.com/jhahspu/golang-gin/entity"

// VideoService interface that includes two function to Save and Return videos
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

// implement interface
type videoService struct {
	videos []entity.Video
}

// New VideoService is a constructor that returns pointer to struct
func New() VideoService {
	return &videoService{}
}

// Save video
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

// FindAll return slice of videos
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
