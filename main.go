package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	v1 := r.Group("/v1")
	v1.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"OPTION", "GET", "POST", "PUT", "DELETE"}
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))
	log.Fatal(r.Run(":" + port))
}
