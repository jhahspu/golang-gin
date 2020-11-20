package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jhahspu/golang-gin/controller"
	"github.com/jhahspu/golang-gin/middleware"
	"github.com/jhahspu/golang-gin/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	// init new server with gin
	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth(), gindump.Dump())

	// serve static files
	server.Static("/css", "./templates/css")

	// serve html templates
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input valid"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
