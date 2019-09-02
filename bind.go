package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2018-02-09" time_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/go", test)
	route.Run(":8080")
}

func test(c *gin.Context) {
	var person Person

	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)

		c.String(http.StatusOK, "ok")
	}
}
