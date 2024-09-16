package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", message)
	router.GET("/app", app)

	log.Println("app is running on http://localhost:8080/")

	router.Run()
}

func message(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "hellow world"},
	)
}

func app(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		gin.H{"message": "this is the app"},
	)
}
