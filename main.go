package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ochom/jumia-interview-task/database"
	"github.com/ochom/jumia-interview-task/rest"
	"github.com/ochom/jumia-interview-task/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, HEAD, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	server := gin.Default()
	server.Use(cors())

	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})
	utils.FailOnError(err)

	repo := database.New(db)

	h := rest.New(repo)

	api := server.Group("/api")
	{
		api.GET("/countries", h.GetCountries())
		api.GET("/numbers/:code/:state", h.GetPhonenumbers())
	}

	err = server.Run(fmt.Sprintf(":%s", utils.GetEnv("PORT", "8000")))
	if err != nil {
		log.Fatalf("server failed to launch: %s", err.Error())
	}
}
