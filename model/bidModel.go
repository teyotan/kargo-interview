package model

import "math/rand"

// Bid ....
type Bid struct {
	ID          int
	UserID      int
	JobID       int
	Price       int64
	VehicleType string
	Status      string
}

func GetBidData() []Bid {
	var bid Bid
	var bidList []Bid

	bidList = make([]Bid, 0)
	for i := 1; i < 25; i++ {
		rand.Seed(int64(1+i))
		userID := rand.Intn(3)+1
	
		rand.Seed(int64(2+i))
		jobID := rand.Intn(9)
	
		rand.Seed(int64(3+i))
		price := int64(rand.Intn(1000000))

		bid = Bid{
			ID:          i,
			UserID:      userID,
			JobID:       jobID,
			Price:       price,
			VehicleType: "Troton",
			Status:      "Active",
		}
		bidList = append(bidList, bid)
	}

	return bidList
}
