package reservations

import (
	"testing"

	"github.com/JoseFMP/ru"
	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {

	responseMock := "<ReservationID>333</ReservationID>"
	responseMockAsBytes := []byte(responseMock)

	parsedResponse, errParsingResponse := parsePutLeadResponse(responseMockAsBytes)
	assert.Nil(t, errParsingResponse)
	assert.NotNil(t, parsedResponse)
	assert.Equal(t, ru.ReservationID(333), parsedResponse)
}
