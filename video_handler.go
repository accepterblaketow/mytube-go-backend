package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVideos(c *gin.Context) {
	search := c.Query("search")
	query := "SELECT id, title, description, tags, url FROM videos"
	var rows *sql.Rows
	var err error

	if search != "" {
		query += " WHERE title LIKE ? OR description LIKE ? OR tags LIKE ?"
		searchTerm := "%" + search + "%"
		rows, err = db.Query(query, searchTerm, searchTerm, searchTerm)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var videos []Video
	for rows.Next() {
		var v Video
		if err := rows.Scan(&v.ID, &v.Title, &v.Description, &v.Tags, &v.URL); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		videos = append(videos, v)
	}
	if videos == nil {
		videos = []Video{} // ✅ 這行是關鍵，強制空陣列不是 null
	}

	c.JSON(http.StatusOK, videos)
}

func GetVideoByID(c *gin.Context) {
	id := c.Param("id") // <- 這裡改了
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}

	query := "SELECT id, title, description, tags, url FROM videos WHERE id = ?"
	row := db.QueryRow(query, id)

	var v Video
	err := row.Scan(&v.ID, &v.Title, &v.Description, &v.Tags, &v.URL)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, v)
}
