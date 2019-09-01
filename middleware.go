package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	route := gin.Default()

	route.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	route.GET("user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	v1 := route.Group("/v1")
	{
		v1.GET("login2", middleware1, handler)
	}
	route.Run(":8008")
}

func middleware1(c *gin.Context) {
	log.Println("in Middleware1")

	c.Next()

	log.Println("out Middleware1")
}

func handler(c *gin.Context) {
	log.Println("Hello %s", "11")

}
