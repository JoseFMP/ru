package ru

// StatusID is the status id defined by RentalsUnited API
type StatusID int16

// ResponseID of a response coming from RentalsUnited
type ResponseID string

// Status packs the status id and the response ID of a response coming from RentalsUnited API
type Status struct {
	ID     StatusID `xml:"ID,attr"`
	Status string   `xml:",chardata"`
}

// Error from RU API
type Error struct {
	//XMLName string   `xml:"Error"`
	ID      StatusID `xml:"ID,attr"`
	Message string   `xml:",chardata"`
}

type AllStatuses struct {
	Success                Status
	RateLimitSlidingWindow Status
}

func GetAllStatuses() AllStatuses {
	return AllStatuses{
		Success: Status{
			ID:     0,
			Status: "Success",
		},
		RateLimitSlidingWindow: Status{
			ID:     -6,
			Status: "This request was rate limited. Maximum number of requests allowed for this API method with specified parameters is {0} per {1} sliding window. You have reached this limit.",
		},
	}
}
