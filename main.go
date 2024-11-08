package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type FileResponse struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
}

func main() {
	router := gin.Default()

	// Create uploads directory if it doesn't exist
	if err := os.MkdirAll("uploads", 0755); err != nil {
		panic(err)
	}

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	router.POST("/api/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()

		if err != nil {
			fmt.Printf("Error getting multipart form: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		files := form.File["files"]

		fmt.Println(len(files))

		if len(files) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no files uploaded"})
			return
		}

		uploadedFiles := make([]map[string]string, 0)

		for _, file := range files {
			fmt.Println(file.Filename)
			ext := filepath.Ext(file.Filename)
			filename := strings.TrimSuffix(file.Filename, ext)
			finalPath := filepath.Join("uploads", filename+ext)
			fmt.Println(finalPath)

			counter := 0
			for {
				if _, err := os.Stat(finalPath); os.IsNotExist(err) {
					break
				}
				counter++
				finalPath = filepath.Join("uploads", filename+strconv.Itoa(counter)+ext)
			}

			if err := saveFile(file, finalPath); err != nil {
				fmt.Printf("Error saving file: %v\n", err) // Debug log
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			uploadedFiles = append(uploadedFiles, map[string]string{
				"url":          "/uploads/" + filepath.Base(finalPath),
				"filename":     filepath.Base(finalPath),
				"originalName": file.Filename,
				"size":         fmt.Sprintf("%.2f MB", float64(file.Size)/(1024*1024)),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"files":   uploadedFiles,
			"message": fmt.Sprintf("Successfully uploaded %d files", len(uploadedFiles)),
		})
	})

	router.Static("/uploads", "./uploads")

	router.Run(":8080")
}

func saveFile(fileHeader *multipart.FileHeader, dstPath string) error {
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

