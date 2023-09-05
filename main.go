package main

import (
	"fmt"
	"log"
	"os"

	"vix-btpns/database"
	"vix-btpns/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	err = database.ConnectDatabase()
	if err != nil {
		fmt.Println("Failed connect database")
	}

	r.Static("/uploads", "./uploads")

	router.UserRouter(r)
	router.PhotosRouter(r)

	r.Run(":" + os.Getenv("PORT"))
}
