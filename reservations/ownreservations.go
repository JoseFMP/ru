package reservations

import (
	"fmt"
	"log"
	"time"

	"github.com/JoseFMP/ru"

	"github.com/JoseFMP/ru/utils"
)

func CanUseListReservations(now time.Time, from time.Time, to time.Time) bool {
	oneWeekAgo := now.Add(utils.DurationSevenDays)
	if from.Before(oneWeekAgo) {
		return false
	}
	if from.Sub(to) > utils.DurationSevenDays {
		return false
	}
	return true
}

var allStatuses = ru.GetAllStatuses()
var commands = ru.GetAllCommands()

func (config *reservationsHandlerConfig) GetOwnReservations(from time.Time, to time.Time) ([]Reservation, error) {
	log.Printf("[ru] GetOwnReservations {%s, %s}\n", from.Format(utils.DateLayout), to.Format(utils.DateLayout))
	requestPayload := composeGetOwnReservationsRequestPayload(from, to)
	log.Printf("[ru][GetOwnReservations] payload:\n%s", requestPayload)
	response, errorDoingRequest := config.DoRequestPull(commands.GetOwnReservations, requestPayload)
	if errorDoingRequest != nil {
		log.Println(errorDoingRequest)
		return nil, errorDoingRequest
	}
	if response.Status.ID == allStatuses.RateLimitSlidingWindow.ID {
		return nil, fmt.Errorf(response.Status.Status)
	}
	if response.Status.ID != allStatuses.Success.ID {
		if response.Status.ID != allStatuses.RateLimitSlidingWindow.ID {
			log.Printf("There was an error response from RU: %s\n", response.Status.Status)
		}
		return nil, fmt.Errorf("Error response from RU: %s", response.Status.Status)
	}
	reservations, errorParsingResponse := parseListReservationsResponse(response.Payload.Payload)
	if errorParsingResponse != nil {
		log.Printf("Error parsing response: %s", errorParsingResponse)
		return nil, errorParsingResponse
	}
	return reservations, nil
}

func composeGetOwnReservationsRequestPayload(from time.Time, to time.Time) string {
	fromAsString := from.Format(utils.TimeDateLayout)
	toAsString := to.Format(utils.TimeDateLayout)
	payload := fmt.Sprintf(listReservationsRequestPayloadTemplate, fromAsString, toAsString, phuketRegionLocationAsString)
	return payload
}
