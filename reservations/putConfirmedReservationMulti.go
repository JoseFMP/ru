package reservations

import (
	"encoding/xml"
	"fmt"

	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru"
	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru/quotes"
	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru/utils"
)

func (config *reservationsHandlerConfig) PutConfirmedReservationMulti(params *ConfirmedReservationParams,
	cc *ru.CreditCard,
	quoteMode quotes.QuoteMode) (ru.ReservationID, error) {

	err := validatePutConfirmedReservationParam(*params)
	if err != nil {
		return 0, err
	}
	payload, errCreatingPayload := createPayload(params, cc, quoteMode)
	if errCreatingPayload != nil {
		return 0, nil
	}
	response, errDoingReq := config.DoRequestPushWithBytes(commands.PutConfirmedReservationMulti, payload)
	if errDoingReq != nil {
		return 0, errDoingReq
	}
	if response.Status.ID != allStatuses.Success.ID {
		return 0, fmt.Errorf(response.Status.Status)
	}
	reservationID, errParsingResponse := parsePutLeadResponse(response.Payload.Payload)
	if errParsingResponse != nil {
		return 0, errParsingResponse
	}
	return reservationID, nil
}

func validatePutConfirmedReservationParam(params ConfirmedReservationParams) error {

	if params.CustomerInfo.Email == "" {
		return fmt.Errorf("Customer email not specified")
	}

	if params.CustomerInfo.Name == "" {
		return fmt.Errorf("Customer Name not specified")
	}

	from := params.Period.From
	to := params.Period.To

	if from.Year() < 2020 {
		return fmt.Errorf("Year has to be greater than 2020!")
	}

	if to.Year() < from.Year() {
		return fmt.Errorf("To year cannot be lower than From year")
	}

	if to.Year() == from.Year() && to.YearDay() <= from.YearDay() {
		return fmt.Errorf("To day cannot be before from day")
	}

	if params.PropertyID == 0 {
		return fmt.Errorf("The RU property ID cannot be 0")
	}

	return nil
}

func createPayload(params *ConfirmedReservationParams, cc *ru.CreditCard, quoteMode quotes.QuoteMode) ([]byte, error) {

	reservation := PutConfirmedReservationRequestReservation{
		Comments:     params.Comments,
		CustomerInfo: params.CustomerInfo,
		StayInfos: []StayInfo{
			StayInfo{
				DateFrom:       utils.ToURUDateString(&params.Period.From),
				DateTo:         utils.ToURUDateString(&params.Period.To),
				AmountOfUnits:  1,
				NumberOfGuests: params.NumberOfGuests,
				PropertyID:     params.PropertyID,
				Costs:          params.Costs,
			},
		},
		CreditCard: cc,
	}

	rawPayload, errMarshalling := xml.Marshal(reservation)
	if errMarshalling != nil {
		return nil, errMarshalling
	}

	quotesStrings := fmt.Sprintf("<QuoteModeId>%d</QuoteModeId>", quoteMode)
	finalString := fmt.Sprintf("%s\n%s", string(rawPayload), quotesStrings)
	return []byte(finalString), nil
}

type ConfirmedReservationParams struct {
	Period         utils.Period
	PropertyID     ru.BasePropertyID
	NumberOfGuests uint16
	CustomerInfo   ru.CustomerInfo
	Comments       string
	Costs          Costs
}

type PutConfirmedReservationRequestPayload struct {
	XMLName     struct{} `xml:"Push_PutConfirmedReservationMulti_RQ"`
	QuoteModeId quotes.QuoteMode
	Reservation PutConfirmedReservationRequestReservation
}

type PutConfirmedReservationRequestReservation struct {
	XMLName      struct{}        `xml:"Reservation"`
	StayInfos    []StayInfo      `xml:">StayInfo" json:"stayInfos"`
	CustomerInfo ru.CustomerInfo `json:"customerInfo"`
	Comments     string          `json:"comments"`
	CreditCard   *ru.CreditCard  `json:"creditCard"`
}
