package service

import (
	"github.com/gin-gonic/gin"
	"github.com/teyotan/kargo-interview/model"
)


func GetJobs(c *gin.Context) {
	shipmentDate := c.DefaultQuery("date", "ASC")

	if shipmentDate != "ASC" && shipmentDate != "DESC" {
		c.JSON(400, gin.H{
			"message": "date param must be 'ASC' or 'DESC'",
		})
		return
	}

	jobs := model.GetJobData()

	c.JSON(200, gin.H{
		"data": jobs,
	})
}