package main

import (
	"embed"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed resources
	res  embed.FS
	port uint
)

func init() {
	flag.UintVar(&port, "port", 8080, "Port number")
}

func main() {
	flag.Parse()

	router := gin.Default()
	router.GET("/", indexHandler)
	router.GET("/resources/*rest", resourceHandler)
	router.POST("/game/start", startGameHandler)
	router.POST("/game/oponent/move", oponentMoveHandler)
	router.GET("/game/ai/move", aiMoveHandler)
	router.GET("/game/next", nextPositions)

	portString := fmt.Sprintf(":%d", port)
	router.Run(portString)
}
