package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jhahspu/golang-gin/entity"
	"github.com/jhahspu/golang-gin/service"
)

// VideoController has two functions for FindAll and Save videos
type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

// Controller struct to control the interface
type controller struct {
	service service.VideoService
}

// New constructor to return video controller
func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

// FindAll returns slice on videos
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// Save json video and return video
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
