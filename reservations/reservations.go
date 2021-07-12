package reservations

import (
	"time"

	"github.com/JoseFMP/ru"
	"github.com/JoseFMP/ru/quotes"
)

type ReservationsHandler interface {
	PutConfirmedReservationMulti(params *ConfirmedReservationParams, cc *ru.CreditCard, quoteMode quotes.QuoteMode) (ru.ReservationID, error)
	PutLead(lead *PutLeadRequest) (ru.ReservationID, error)
	GetOwnReservations(from time.Time, to time.Time) ([]Reservation, error)
	ModifyStay(res ru.ReservationID, currentStay CurrentStay, newAlreadyPaid *float64) error
}

type reservationsHandlerConfig struct {
	ru.ClientConfig
}

func GetClientReservationsHandler(userName string, password string) ReservationsHandler {
	ruClient := ru.ClientConfig{
		UserName: userName,
		Password: password,
	}
	return GetClientReservationsHandlerFromRUClient(ruClient)
}

func GetClientReservationsHandlerFromRUClient(ruClient ru.ClientConfig) ReservationsHandler {
	return &reservationsHandlerConfig{
		ClientConfig: ruClient,
	}
}
