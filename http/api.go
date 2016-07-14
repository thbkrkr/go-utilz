package http

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// API provides an HTTP API based gin-gonic with cors and base routes
func API(name string, buildDate string, gitCommit string, f func(r *gin.Engine)) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(cORSMiddleware())

	if _, err := os.Stat("./_static"); !os.IsNotExist(err) {
		r.Static("/s", "./_static/")
		r.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/s")
		})
	}

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"buildDate": buildDate,
			"gitCommit": gitCommit,
			"name":      name,
			"ok":        "true",
			"status":    200,
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.JSON(200, nil)
	})

	f(r)

	logrus.WithFields(logrus.Fields{
		"buildDate": buildDate,
		"gitCommit": gitCommit,
		"name":      name,
		"port":      4242,
	}).Info("Start")

	err := r.Run(":4242")
	if err != nil {
		logrus.Fatal(err)
	}
}

func cORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := "*"
		c.Writer.Header().Set("Access-Control-Allow-Origin", domain)
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
