package ru

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru/utils"
)

func (clientConfig *ClientConfig) ListAvailabilityCalendar(from time.Time, to time.Time, propertyID BasePropertyID) (*CalendarChunk, error) {
	shrankFrom, shrankTo := shirnkFromTo(from, to)

	commandPayload := getAvailabilityCommandPayload(shrankFrom, shrankTo, propertyID)
	commandResponse, err := clientConfig.DoRequestPull(commands.ListPropertyAvailabilityCalendar, commandPayload)
	if err != nil {
		return nil, err
	}
	calDays, errParsing := parseAvailabilityCommandPayload(commandResponse.Payload.Payload)
	if errParsing != nil {
		return nil, errParsing
	}
	calendarChunk := createCalendarChunk(calDays)
	return calendarChunk, nil
}

func shirnkFromTo(originalFrom time.Time, originalTo time.Time) (time.Time, time.Time) {

	var from, to time.Time
	if originalFrom.IsZero() || originalFrom.Before(time.Now()) {
		from = time.Now()
	} else {
		from = originalFrom
	}

	if originalTo.Before(from) {
		to = from.Add(time.Duration(utils.Duration1Day * 3))
	} else {
		to = originalTo
	}

	return from, to
}

func createCalendarChunk(calDays []CalDay) *CalendarChunk {
	calendarChunk := make(CalendarChunk)
	for _, cd := range calDays {
		date, errParsing := time.Parse(utils.DateLayout, cd.Date)
		if errParsing != nil {
			log.Println("Error parsing date")
			continue
		}
		calendarChunk[date] = cd
	}
	return &calendarChunk
}

func getAvailabilityCommandPayload(from time.Time, to time.Time, propertyID BasePropertyID) string {
	idAsString := propertyID.GetBasePropertyIDAsString()
	fromAsString := from.Format(utils.DateLayout)
	toAsString := to.Format(utils.DateLayout)
	payload := fmt.Sprintf(commandPayloadTemplate, idAsString, fromAsString, toAsString)
	return payload
}

func parseAvailabilityCommandPayload(payload []byte) ([]CalDay, error) {

	var parsedPayload struct {
		PropertyCalendar []CalDay `xml:"CalDay"`
	}
	errorParsing := xml.Unmarshal(payload, &parsedPayload)
	if errorParsing != nil {
		return nil, errorParsing
	}
	return parsedPayload.PropertyCalendar, nil
}

type CalendarChunk = map[time.Time]CalDay

type CalDay struct {
	Date         string `xml:"Date,attr"`
	Units        uint32 `xml:"Units,attr"`
	Reservations uint32 `xml:"Reservations,attr"`
	IsBlocked    bool   `xml:"IsBlocked"`
	MinStay      uint32
	Changeover   uint32
}

const commandPayloadTemplate = `
  <PropertyID>%s</PropertyID>
  <DateFrom>%s</DateFrom>
  <DateTo>%s</DateTo>
  `
