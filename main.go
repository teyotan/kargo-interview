package main

import (
	"github.com/teyotan/kargo-interview/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	
	r.GET("/getJobs", service.GetJobs)	
	r.GET("/getBids", service.GetBids)

	r.Run() // listen and serve on 0.0.0.0:8080

	// fmt.Println(model.GetUserData())
	// fmt.Println()
	// fmt.Println(model.GetJobData())
	// fmt.Println()
	// fmt.Println(model.GetBidData())
}
