package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func main() {
	// Handle SIGINT
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			fmt.Printf("\nReceived an interrupt, stopping services...\n")
			os.Exit(0)
		}
	}()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"GOOS":         runtime.GOOS,
			"GOARCH":       runtime.GOARCH,
			"Version":      runtime.Version(),
			"NumCPU":       runtime.NumCPU(),
			"NumGoroutine": runtime.NumGoroutine(),
			"NumCgoCall":   runtime.NumCgoCall(),
			"Time":         time.Now().Format(time.RFC1123),
		})
	})
	r.GET("/cpuInfo", func(c *gin.Context) {
		// Read /proc/cpuinfo
		f, err := os.Open("/proc/cpuinfo")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer f.Close()

		// Copy to in-memory buffer
		b, err := io.ReadAll(f)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.String(http.StatusOK, string(b))
	})
	r.Run(":8080") // listen and serve on :8080
}
