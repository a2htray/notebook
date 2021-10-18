package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/v1/stations/:stationID", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:stationID", RemoveStation)
}

func GetStation(c *gin.Context) {

}

func CreateStation(c *gin.Context) {

}

func RemoveStation(c *gin.Context) {

}
