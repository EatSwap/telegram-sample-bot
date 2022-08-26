package main

import (
	"github.com/gin-gonic/gin"
	"runtime"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"GOOS":         runtime.GOOS,
			"GOARCH":       runtime.GOARCH,
			"Version":      runtime.Version(),
			"NumCPU":       runtime.NumCPU(),
			"NumGoroutine": runtime.NumGoroutine(),
			"NumCgoCall":   runtime.NumCgoCall(),
			"Time":         time.Now().Format(time.RFC1123),
		})
	})
	r.Run(":10000") // listen and serve on :8080
}
