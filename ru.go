package ru

import (
	"time"
)

// Client represents a RU API client
type Client interface {
	Authenticate() error
	ListAvailabilityCalendar(from time.Time, to time.Time, propertyID BasePropertyID) (*CalendarChunk, error)
	GetPropertyAveragePrice(propertyID BasePropertyID, from time.Time, to time.Time) ([]PropertyPrice, error)
}

const baseURLString = "https://rm.rentalsunited.com/api/Handler.ashx"
