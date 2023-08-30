package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type ClientResponse struct {
	Message string
}

func main() {
	port := "9999"
	goServiceUrl, isUrlExist := os.LookupEnv("GO_SERVICE_URL")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		if !isUrlExist {
			c.JSON(500, gin.H{
				"response": "Service is not available",
			})
			return
		}
		resp, getErr := http.Get(goServiceUrl)
		if getErr != nil {
			c.JSON(500, gin.H{
				"response": "An error occurred on service request",
			})
			return
		}
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			c.JSON(500, gin.H{
				"response": "An error occurred when reading response",
			})
			return
		}
		resData := ClientResponse{}
		err := json.Unmarshal(body, &resData)
		if err != nil {
			c.JSON(500, gin.H{
				"response": "An error occurred when marshalling response",
			})
			return
		}
		c.JSON(200, gin.H{
			"response": resData.Message,
		})
	})
	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
