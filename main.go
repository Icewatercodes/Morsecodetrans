package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type injson struct {
	InMsg string `json:"Msg"`
}

/*type outjson struct {
	OutMsg string `json:Msg`
}*/

func main() {
	router := gin.Default()
	router.GET("/api/translate/morse", translateStringMorse)
	router.Run(":4300")
}

func translateStringMorse(c *gin.Context) {
	req := c.Request.Body

	c.JSON(http.StatusOK, req)
}
