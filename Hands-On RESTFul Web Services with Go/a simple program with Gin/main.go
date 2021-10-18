package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/pingTime", func(c *gin.Context) {
		c.JSON(http.StatusOK, time.Now().Unix())
	})

	log.Fatal(r.Run(":8000"))
}
