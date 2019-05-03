package service

import (
	"github.com/gin-gonic/gin"
	"github.com/teyotan/kargo-interview/model"
)


func GetBids(c *gin.Context) {
	var asc bool
	price := c.DefaultQuery("price", "ASC")

	if price != "ASC" && price != "DESC" {
		c.JSON(400, gin.H{
			"message": "Price param must be 'ASC' or 'DESC'",
		})
		return
	}

	if price == "ASC" {
		asc = true
	} else {
		asc = false
	}

	bids := model.GetBidData()

	b := make([]interface{}, len(bids))
	for i := range bids {
		b[i] = bids[i]
	}

	b = QuickSort(b, getBidPrice, asc)

	c.JSON(200, gin.H{
		"data": b,
	})
}

func getBidPrice(bid interface{}) int{
	return int(bid.(model.Bid).Price)
}