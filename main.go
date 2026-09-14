package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/raflimuhammadd/mrt-schedule/modules/station"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}
	log.Printf("API_BASE_URL: %s", os.Getenv("API_BASE_URL"))

	initiateRouter()
}

func initiateRouter() {
	var (
		router = gin.Default()
		api = router.Group("/v1/api")
	)

	station.Initiate(api)

	router.Run(":8080")
}