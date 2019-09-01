package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	route := gin.Default()

	route.POST("json", func(c *gin.Context) {
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(http.StatusOK,"ok")
	})

	route.Run(":8080")
}
