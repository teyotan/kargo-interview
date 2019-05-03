package service

import (
	"github.com/gin-gonic/gin"
	"github.com/teyotan/kargo-interview/model"
)


func GetBids(c *gin.Context) {
	price := c.DefaultQuery("price", "ASC")

	if price != "ASC" && price != "DESC" {
		c.JSON(400, gin.H{
			"message": "Price param must be 'ASC' or 'DESC'",
		})
		return
	}

	bids := model.GetBidData()

	c.JSON(200, gin.H{
		"data": bids,
	})
}