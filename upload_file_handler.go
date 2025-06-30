package main

import (
    "fmt"
    "net/http"
    "path/filepath"
    "time"
    "github.com/gin-gonic/gin"
)
const maxUploadSize = 4 << 30 // 4GB
func UploadVideoFile(c *gin.Context) {
    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

    err := c.Request.ParseMultipartForm(maxUploadSize)
    if err != nil {
        fmt.Println("上傳解析失敗:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    title := c.PostForm("title")
    description := c.PostForm("description")
    tags := c.PostForm("tags")

    file, header, err := c.Request.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "File required"})
        return
    }
    defer file.Close()

    ext := filepath.Ext(header.Filename)
    timestamp := time.Now().Unix()
    filename := fmt.Sprintf("video_%d%s", timestamp, ext)

    url, err := UploadToB2(file, filename, header.Header.Get("Content-Type"))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    _, err = db.Exec("INSERT INTO videos (title, description, tags, url) VALUES (?, ?, ?, ?)", title, description, tags, url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Video uploaded", "url": url})
}