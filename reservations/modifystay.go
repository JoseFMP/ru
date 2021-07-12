package reservations

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/JoseFMP/ru/utils"

	"github.com/JoseFMP/ru"
)

func (client *reservationsHandlerConfig) ModifyStay(res ru.ReservationID, currentStay CurrentStay, newAlreadyPaid *float64) error {

	log.Printf("[ru] ModifyStay\n")
	rpload, errCreatingPayload := composeModifyStayPayload(res, currentStay, newAlreadyPaid)
	if errCreatingPayload != nil {
		return errCreatingPayload
	}
	result, errorIssuingCommand := client.DoRequestPushWithBytes(commands.ModifyStay, rpload)

	if errorIssuingCommand != nil {
		return errorIssuingCommand
	}
	if result.Status.ID != allStatuses.Success.ID {
		return fmt.Errorf("Error issuing ModifyStay command")
	}

	return nil
}

func composeModifyStayPayload(res ru.ReservationID, currentStay CurrentStay, newAlreadyPaid *float64) ([]byte, error) {

	payload := ModifyStayPayload{
		ReservationID: res,
		Current: CurrentStayPayload{
			PropertyID: currentStay.PropertyID,
			DateFrom:   utils.ToURUDateString(&currentStay.DateFrom),
			DateTo:     utils.ToURUDateString(&currentStay.DateTo),
		},
		Modify: ModifyPayload{
			AlreadyPaid: newAlreadyPaid,
		},
	}

	marshalled, errorMarshalling := xml.MarshalIndent(payload, "", " ")
	if errorMarshalling != nil {
		return nil, errorMarshalling
	}
	return marshalled, nil
}

type CurrentStay struct {
	PropertyID ru.BasePropertyID
	DateFrom   time.Time
	DateTo     time.Time
}

type ModifyStayPayload struct {
	ReservationID ru.ReservationID   `xml:"ReservationID"`
	Current       CurrentStayPayload `xml:"Current"`
	Modify        ModifyPayload      `xml:"Modify"`
}

type CurrentStayPayload struct {
	PropertyID ru.BasePropertyID `xml:"PropertyID"`
	DateFrom   string
	DateTo     string
}

type ModifyPayload struct {
	AlreadyPaid *float64 `xml:"AlreadyPaid"`
}
