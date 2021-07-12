package ru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanParsePayload(t *testing.T) {

	payloadAsBytes := []byte(mockPayloadPropertyPrices)

	parsedPayload, errParsingPayload := parsePropertyAveragePricePayload(payloadAsBytes)

	assert.Nil(t, errParsingPayload)
	assert.Len(t, parsedPayload, 3)
	assert.NotNil(t, parsedPayload[0])
	assert.Equal(t, float64(10.0), parsedPayload[0].Cleaning)

	assert.Equal(t, float64(100), parsedPayload[0].Price)

}

const mockPayloadPropertyPrices = `<PropertyPrices>
   <PropertyPrice NOP="3" Cleaning="10.00" ExtraPersonPrice="0.00" Deposit="15.00" SecurityDeposit="15.00">
   100.00
   </PropertyPrice>
   <PropertyPrice NOP="4" Cleaning="10.00" ExtraPersonPrice="15.00" Deposit="16.50" SecurityDeposit="16.50">
   110.00
   </PropertyPrice>
   <PropertyPrice NOP="5" Cleaning="10.00" ExtraPersonPrice="40.00" Deposit="18.00" SecurityDeposit="18.00">
   120.00
   </PropertyPrice>
</PropertyPrices>
`
