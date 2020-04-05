package main

import (
	"github.com/KouT127/go-cloud-run/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	initDatabase()
	initRouter()
}

func initDatabase() {
	dbHost := os.Getenv("DB_TCP_HOST")
	if dbHost == "" {
		err := database.InitSocketConnectionPool()
		if err != nil {
			log.Fatalf("Socket connection is unavailable")
		}
	} else {
		err := database.InitTcpConnectionPool()
		if err != nil {
			log.Fatalf("Tcp connection is unavailable")
		}

	}
}

func initRouter() {
	r := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"OPTION", "GET", "POST", "PUT", "DELETE"}
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}

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

	r.Use(cors.New(config))
	log.Fatal(r.Run(":" + port))
}
