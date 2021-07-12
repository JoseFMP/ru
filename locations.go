package ru



func GetAllLocations() *Locations {
	return &Locations{
		Worldwide: &Location{
			LocationTypeID: allLocationsTypes.Worldwide,
			LocationID:     1,
		},
		Europe: &Location{
			LocationTypeID: allLocationsTypes.Continent,
			LocationID:     2,
		},
		North: &Location{
			LocationTypeID: allLocationsTypes.Continent,
			LocationID:     3,
		},
		South: &Location{
			LocationTypeID: allLocationsTypes.Continent,
			LocationID:     4,
		},
		Asia: &Location{
			LocationTypeID: allLocationsTypes.Continent,
			LocationID:     5,
		},
		Australia: &Location{
			LocationTypeID: allLocationsTypes.Continent,
			LocationID:     6,
		},
		Antarctica: &Location{
			LocationTypeID: allLocationsTypes.Continent,
			LocationID:     7,
		},
		Africa: &Location{
			LocationTypeID: allLocationsTypes.Continent,
			LocationID:     8,
		},
		PhuketRegion: &Location{
			LocationTypeID: allLocationsTypes.Region,
			LocationID:     10977,
		},
		Thailand: &Location{
			LocationTypeID: allLocationsTypes.Country,
			LocationID:     340,
		},
	}
}

var allLocationsTypes = LocationTypes{
	Worldwide: 6,
	Continent: 1,
	Country:   2,
	Region:    3,
	City:      4,
	District:  5,
}
