package reservations

import (
	"strings"
	"testing"

	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru"
	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru/quotes"

	"github.com/stretchr/testify/assert"
)

func TestConfirmedResCreatePayload(t *testing.T) {

	mockParams := ConfirmedReservationParams{
		Comments: "tralara",
	}
	payload, errCreatingPayload := createPayload(&mockParams, mockCC, quotes.QuoteIgnorePMSAndRU)

	assert.Nil(t, errCreatingPayload)
	assert.NotNil(t, payload)

	asString := string(payload)

	assert.True(t, strings.Contains(asString, "<Reservation>"))
}

var mockCC = ru.CreditCard{}
