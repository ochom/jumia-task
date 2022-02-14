package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ochom/jumia-interview-task/database"
	"github.com/ochom/jumia-interview-task/rest"
	"github.com/ochom/jumia-interview-task/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()

	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{})
	utils.FailOnError(err)

	repo := database.New(db)

	h := rest.New(repo)

	api := server.Group("/api")
	{
		api.GET("/numbers", h.GetPhonenumbers())
		api.GET("/numbers/:code", h.GetCountryPhoneNumbers())
	}

	server.Run(fmt.Sprintf(":%s", utils.GetEnv("PORT", "8000")))
}
