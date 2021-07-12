package reservations

import (
	"strings"
	"testing"

	"github.com/JoseFMP/ru"
	"github.com/JoseFMP/ru/quotes"

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
