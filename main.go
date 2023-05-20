package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/upload", uploadImages)
	router.GET("/test", testConnectivity)
	router.Run(":8080")
}
func testConnectivity(c *gin.Context){
	fmt.Print("connection is fine")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("you can now upload files")})

}

func uploadImages(c *gin.Context) {
	fmt.Print("hello. apk knoked")
	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the assets folder if it doesn't exist
	assetsDir := "assets"
	if _, err := os.Stat(assetsDir); os.IsNotExist(err) {
		if err := os.Mkdir(assetsDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	// Save the file to the assets folder
	filename := filepath.Base(file.Filename)
	dst := fmt.Sprintf("%s/%s", assetsDir, filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File '%s' uploaded successfully.", filename)})
}
