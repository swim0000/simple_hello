package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := 3000

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get hostname")
			return
		}

		message := fmt.Sprintf("Welcome to %s", hostname)
		fmt.Println(message)
		c.String(http.StatusOK, message)
	})

	router.Run(fmt.Sprintf(":%d", port))
}
