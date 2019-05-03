package service

import (
	"github.com/gin-gonic/gin"
	"github.com/teyotan/kargo-interview/model"
)


func GetJobs(c *gin.Context) {
	var asc bool
	shipmentDate := c.DefaultQuery("date", "ASC")

	if shipmentDate != "ASC" && shipmentDate != "DESC" {
		c.JSON(400, gin.H{
			"message": "date param must be 'ASC' or 'DESC'",
		})
		return
	}

	if shipmentDate == "ASC" {
		asc = true
	} else {
		asc = false
	}

	jobs := model.GetJobData()

	b := make([]interface{}, len(jobs))
	for i := range jobs {
		b[i] = jobs[i]
	}

	b = QuickSort(b, getJobUnixDate, asc)

	c.JSON(200, gin.H{
		"data": b,
	})
}

func getJobUnixDate(job interface{}) int{
	return int(job.(model.Job).ShipmentDate.Unix())
}