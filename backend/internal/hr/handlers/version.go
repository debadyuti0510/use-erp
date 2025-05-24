package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetVersion(c *gin.Context) {
	// Open the .version file and read the version
	cwd, _ := os.Getwd()
	fmt.Println("The current directory is - ", cwd)
	version, err := os.ReadFile("internal/hr/.version")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to retrieve version with error: %s", err)})
	} else {
		c.JSON(
			http.StatusOK, gin.H{
				"version": string(version),
			})
	}

}
