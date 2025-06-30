package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func UploadVideo(c *gin.Context) {
    title := c.PostForm("title")
    description := c.PostForm("description")
    tags := c.PostForm("tags")
    url := c.PostForm("url")

    if title == "" || url == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "title and url are required"})
        return
    }

    _, err := db.Exec(`INSERT INTO videos (title, description, tags, url) VALUES (?, ?, ?, ?)`,
        title, description, tags, url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Video uploaded!"})
}
