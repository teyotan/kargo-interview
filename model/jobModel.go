package model

import (
	"math/rand"
	"time"
)

// Job ....
type Job struct {
	ID           int
	UserID       int
	Destination  string
	Origin       string
	ShipmentDate time.Time
	Status       string
}

func randate(seed int) time.Time {
	min := time.Date(2019, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	rand.Seed(int64(seed))
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func GetJobData() []Job {
	var job Job
	var jobList []Job

	destinationList := []string{"Jakarta", "Semarang", "Surabaya", "Makassar", "London"}

	jobList = make([]Job, 0)
	for i := 1; i < 10; i++ {
		rand.Seed(int64(1+i))
		userID := rand.Intn(4) + 1

		rand.Seed(int64(2+i))
		destinationKey := rand.Intn(5)

		rand.Seed(int64(3+i))
		originKey := rand.Intn(5)

		job = Job{
			ID:           i,
			UserID:       userID,
			Destination:  destinationList[destinationKey],
			Origin:       destinationList[originKey],
			ShipmentDate: randate(i),
			Status:       "Active",
		}
		jobList = append(jobList, job)
	}

	return jobList
}
