package ru

import "fmt"

type PropertyID struct {
	ID         BasePropertyID `xml:",chardata"`
	BuildingID int64          `xml:"BuildingID,attr"`
}

func (propertyID *BasePropertyID) GetBasePropertyIDAsString() string {
	return fmt.Sprintf("%d", uint64(*propertyID))
}

type Property struct {
	ProviderUID        PropertyID `xml:"PUID"`
	ID                 PropertyID `xml:"ID"`
	Name               string     `xml:"Name"`
	OwnerID            uint64
	DateCreated        string
	DetailedLocationID LocationID
	LastMod            PropertyLastMod
}

type PropertyLastMod struct {
	NLA     bool   `xml:"NLA,attr"`
	LastMod string `xml:",innerxml"`
}

type LocationTypes struct {
	Worldwide LocationTypeID
	Continent LocationTypeID
	Country   LocationTypeID
	Region    LocationTypeID
	City      LocationTypeID
	District  LocationTypeID
}

type Location struct {
	LocationID     LocationID
	LocationTypeID LocationTypeID
}

// Locations All the possible locations out there
type Locations struct {
	Worldwide    *Location
	Europe       *Location
	North        *Location
	South        *Location
	Asia         *Location
	Australia    *Location
	Antarctica   *Location
	Africa       *Location
	PhuketRegion *Location
	Thailand     *Location
}
