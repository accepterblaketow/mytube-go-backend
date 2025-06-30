package main

import (
    "database/sql"
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
)

func DeleteVideoByID(c *gin.Context) {
    id := c.PostForm("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
        return
    }

    var url string
    err := db.QueryRow("SELECT url FROM videos WHERE id = ?", id).Scan(&url)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "video not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 從 URL 中擷取檔名
    parts := strings.Split(url, "/")
    filename := parts[len(parts)-1]

    // 刪除 B2 上的檔案
    if err := DeleteFromB2(filename); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 刪除資料庫中的紀錄
    _, err = db.Exec("DELETE FROM videos WHERE id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"warning": "File deleted, but DB delete failed"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Video deleted"})
}
