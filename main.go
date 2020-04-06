package main

import (
	"github.com/KouT127/go-cloud-run/database"
	"github.com/KouT127/go-cloud-run/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const ReleaseMode = "Release"

var releaseMode string

func main() {
	releaseMode = os.Getenv("RELEASE_MODE")
	database.InitDatabase()
	initRouter()
}

func initRouter() {
	r := gin.Default()

	if releaseMode == ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"OPTION", "GET", "POST", "PUT", "DELETE"}
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	r.Use(cors.New(config))

	v1 := r.Group("/v1")
	v1.Handle("GET", "/", func(c *gin.Context) {
		err := database.Ping()
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"message": "ng"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	logger.Error(r.Run(":" + port))
}
