package reservations

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru"
	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru/utils"
)

var phuketRegionLocationAsString = fmt.Sprintf("%d", ru.GetAllLocations().PhuketRegion.LocationID)

func (config *reservationsHandlerConfig) ListReservations(from time.Time, to time.Time) ([]Reservation, error) {
	log.Printf("[ru] ListReservations {%s, %s}\n", from.Format(utils.DateLayout), to.Format(utils.DateLayout))
	requestPayload := composeListReservationsRequestPayload(from, to)

	response, errorDoingRequest := config.DoRequestPull(commands.ListReservations, requestPayload)
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

func parseListReservationsResponse(response []byte) ([]Reservation, error) {

	var reservations struct {
		Reservations []Reservation `xml:"Reservation"`
	}
	errorParsing := xml.Unmarshal(response, &reservations)
	if errorParsing != nil {
		return nil, errorParsing
	}
	return reservations.Reservations, nil
}

func composeListReservationsRequestPayload(from time.Time, to time.Time) string {
	fromAsString := from.Format(utils.TimeDateLayout)
	toAsString := to.Format(utils.TimeDateLayout)
	payload := fmt.Sprintf(listReservationsRequestPayloadTemplate, fromAsString, toAsString, phuketRegionLocationAsString)
	return payload
}

const listReservationsRequestPayloadTemplate = `<DateFrom>%s</DateFrom>
    <DateTo>%s</DateTo>
    <LocationID>%s</LocationID>`

type Reservation struct {
	ReservationID    ru.ReservationID  `xml:"ReservationID"`
	Status           ReservationStatus `xml:"StatusID"`
	LastMod          string            `xml:"LastMod"`
	StayInfos        []StayInfo        `xml:">StayInfo"`
	PMSReservationID string            `xml:"PMSReservationId"`
	Archived         bool              `xml:"IsArchived"`
	CustomerInfo     ru.CustomerInfo   `xml:"CustomerInfo"`
}

type StayInfo struct {
	PropertyID     ru.BasePropertyID `xml:"PropertyID"`
	DateFrom       string            `xml:"DateFrom"`
	DateTo         string
	NumberOfGuests uint16
	Costs          Costs       `xml:"Costs"`
	ResAptID       ru.ResAptID `xml:"ResapaID"`
	AmountOfUnits  uint8       `xml:"Units"`
}

type Costs struct {
	RUPrice     float64
	ClientPrice float64
	AlreadyPaid float64
}

type ReservationStatus uint8

func GetAllReservationStatuses() AllReservationStatus {
	return AllReservationStatus{
		Confirmed: 1,
		Canceled:  2,
		Modified:  3,
	}
}

var reservationStatusNames = map[ReservationStatus]string{
	GetAllReservationStatuses().Confirmed: "confirmed",
	GetAllReservationStatuses().Canceled:  "canceled",
	GetAllReservationStatuses().Modified:  "modified",
}

func GetReservationStatusName(status ReservationStatus) string {
	return reservationStatusNames[status]
}

type AllReservationStatus struct {
	Confirmed ReservationStatus
	Canceled  ReservationStatus
	Modified  ReservationStatus
}

func ReservationStatusAsString(rs ReservationStatus) {

}
