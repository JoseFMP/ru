package reservations

import (
	"fmt"
	"testing"
	"time"

	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru/utils"
	"github.com/stretchr/testify/assert"

	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru"
)

func TestCreateModifyStayPayloadWontSetAlreadyPaidIfNil(t *testing.T) {

	dateFrom, _ := time.Parse(utils.DateLayout, "2021-10-10")
	dateTo, _ := time.Parse(utils.DateLayout, "2021-10-20")
	reservationIDMock := ru.ReservationID(1234123)

	currentStayMock := CurrentStay{
		PropertyID: ru.BasePropertyID(123123),
		DateFrom:   dateFrom,
		DateTo:     dateTo,
	}
	payload, errCreatingPayload := composeModifyStayPayload(reservationIDMock, currentStayMock, nil)

	assert.Nil(t, errCreatingPayload)
	assert.NotNil(t, payload)

	payloadAsString := string(payload)
	assert.Contains(t, payloadAsString, "<DateFrom>2021-10-10</DateFrom>")
	assert.Contains(t, payloadAsString, "<DateTo>2021-10-20</DateTo>")
	assert.Contains(t, payloadAsString, fmt.Sprintf("<PropertyID>%d</PropertyID>", currentStayMock.PropertyID))
	assert.Contains(t, payloadAsString, "<Modify>")
	assert.Contains(t, payloadAsString, "</Modify>")

	assert.NotContains(t, payloadAsString, "AlreadyPaid")
}

func TestCreateModifyStayPayloadWorks(t *testing.T) {

	dateFrom, _ := time.Parse(utils.DateLayout, "2021-10-10")
	dateTo, _ := time.Parse(utils.DateLayout, "2021-10-20")
	reservationIDMock := ru.ReservationID(1234123)

	currentStayMock := CurrentStay{
		PropertyID: ru.BasePropertyID(123123),
		DateFrom:   dateFrom,
		DateTo:     dateTo,
	}
	alreadyPaid := float64(200.5)
	payload, errCreatingPayload := composeModifyStayPayload(reservationIDMock, currentStayMock, &alreadyPaid)

	assert.Nil(t, errCreatingPayload)
	assert.NotNil(t, payload)

	payloadAsString := string(payload)
	assert.Contains(t, payloadAsString, "<DateFrom>2021-10-10</DateFrom>")
	assert.Contains(t, payloadAsString, "<DateTo>2021-10-20</DateTo>")
	assert.Contains(t, payloadAsString, fmt.Sprintf("<PropertyID>%d</PropertyID>", currentStayMock.PropertyID))
	assert.Contains(t, payloadAsString, "<Modify>")
	assert.Contains(t, payloadAsString, "</Modify>")

	assert.Contains(t, payloadAsString, fmt.Sprintf("<AlreadyPaid>200.5</AlreadyPaid>"))
}
