package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prshendra/demo-day-cicd/handlers"
)

func main() {
	r := gin.Default()
	handlers.RegisterRoutes(r)
	r.Run(":8010")
}
