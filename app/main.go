package main

import "github.com/gin-gonic/gin"

func main() {
	status := true
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong ver2",
		})
	})
	r.GET("/healthcheck", func(c *gin.Context) {
		if status {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		} else {
			c.JSON(500, gin.H{
				"message": "ng",
			})
		}
	})
	r.GET("/changestatus", func(c *gin.Context) {
		status = !status
		c.JSON(200, gin.H{
			"message": status,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
