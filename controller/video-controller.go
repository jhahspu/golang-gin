package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jhahspu/golang-gin/entity"
	"github.com/jhahspu/golang-gin/service"
	"github.com/jhahspu/golang-gin/validators"
)

// VideoController has two functions for FindAll and Save videos
type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

// Controller struct to control the interface
type controller struct {
	service service.VideoService
}

var validate *validator.Validate

// New constructor to return video controller
func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

// FindAll returns slice on videos
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// Save json video and return video
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

// ShowAll will retrieve all videos
// and pass them to the viewRouts /view component
func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
