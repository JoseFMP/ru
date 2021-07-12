package reservations

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"github.com/JoseFMP/ru"
)

func (config *reservationsHandlerConfig) PutLead(lead *PutLeadRequest) (ru.ReservationID, error) {

	requestPayload, errCreatingPayload := createPutLeadRequestPayload(lead)
	if errCreatingPayload != nil {
		return 0, errCreatingPayload
	}
	response, errDoingReq := config.DoRequestPushWithBytes(commands.PutLead, requestPayload)
	if errDoingReq != nil {
		return 0, errDoingReq
	}
	if response.Status.ID != allStatuses.Success.ID {
		log.Printf("RU reported error while trying to put lead.\nLead:%v\nError:%v", lead, response.Status.Status)
		return 0, fmt.Errorf(response.Status.Status)
	}

	reservationID, errParsingResponse := parsePutLeadResponse(response.Payload.Payload)
	if errParsingResponse != nil {
		return 0, errParsingResponse
	}

	return reservationID, nil
}

func parsePutLeadResponse(rawResponse []byte) (ru.ReservationID, error) {
	var response PutLeadResponse
	errUnmarshalling := xml.Unmarshal(rawResponse, &response)
	if errUnmarshalling != nil {
		return 0, errUnmarshalling
	}

	reservationID, errParsingInt := strconv.ParseInt(response.ReservationID, 10, 64)
	if errParsingInt != nil {
		return 0, errParsingInt
	}
	if reservationID == 0 {
		return 0, fmt.Errorf("ReservationID can't be 0, that's weird!")
	}
	return ru.ReservationID(reservationID), nil
}

type PutLeadResponse struct {
	ReservationID string `xml:",innerxml"`
}

func createPutLeadRequestPayload(lead *PutLeadRequest) ([]byte, error) {
	marshalledResult, errMarshalling := xml.Marshal(lead)
	if errMarshalling != nil {
		return nil, errMarshalling
	}
	return marshalledResult, nil
}

type PutLeadRequest struct {
	PropertyID     ru.BasePropertyID
	ExternalLeadID string
	DateFrom       string
	DateTo         string
	NumberOfGuests uint8
	CustomerInfo   ru.CustomerInfo
	Comments       string
}
