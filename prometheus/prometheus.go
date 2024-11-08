package main

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Prometheus Middleware hinzufügen
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Path("/metrics"), // Endpoint, auf dem die Metriken verfügbar sind
	)
	r.Use(p.Instrument())
	r.Use(gin.Recovery())

	// Beispiel-Handler
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})
	r.GET("/panic", func(c *gin.Context) {
		panic("panic")
	})

	r.Run(":8080")
}
