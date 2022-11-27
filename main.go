package main

import (
	inittask "file-store/internal/init"
	"file-store/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	inittask.InitTask()

	r := gin.Default()
	routes.InitRoutes(r)

	r.Run(":8080")
}
