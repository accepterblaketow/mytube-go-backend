package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()
	InitB2()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/videos", GetVideos)
	r.GET("/video/:id", GetVideoByID)
	r.POST("/upload_file", UploadVideoFile)
	r.POST("/add_video", AddVideo) //手動新增影片
	r.POST("/delete_video", DeleteVideoByID)

	r.Run("0.0.0.0:8080")
}
