package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()

	route.POST("jdk", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "water c")

		c.JSON(200, gin.H{
			"code":    200,
			"message": message,
			"nick":    nick,
		})
	})

	route.Run(":8080")
}
