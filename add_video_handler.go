package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddVideo(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	tags := c.PostForm("tags")
	url := c.PostForm("url")

	if title == "" || url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title 和 url 為必填"})
		return
	}

	_, err := db.Exec("INSERT INTO videos (title, description, tags, url) VALUES (?, ?, ?, ?)", title, description, tags, url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "資料庫寫入失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "影片已新增"})
}
