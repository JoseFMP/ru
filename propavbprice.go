package ru

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/JoseFMP/ru/utils"
)

func (clientConfig *ClientConfig) GetPropertyAveragePrice(propertyID BasePropertyID, from time.Time, to time.Time) ([]PropertyPrice, error) {

	errValidatingArguments := validateArguments(from, to)
	if errValidatingArguments != nil {
		return nil, errValidatingArguments
	}

	requestPayload := createGetPropertyAveragePricePayload(propertyID, from, to)
	log.Printf("[GetPropertyAveragePrice] Payload for HTTP call to RU:\n%s", requestPayload)
	commandResponse, errDoingRequest := clientConfig.DoRequestPull(commands.GetPropertyAveragePrice, requestPayload)
	if errDoingRequest != nil {
		return nil, errDoingRequest
	}

	if commandResponse.Status.ID != allStatuses.Success.ID {
		return nil, fmt.Errorf(commandResponse.Status.Status)
	}

	parsedPayload, errParsingPayload := parsePropertyAveragePricePayload(commandResponse.Payload.Payload)
	if errParsingPayload != nil {
		return nil, errParsingPayload
	}
	return parsedPayload, nil
}

func createGetPropertyAveragePricePayload(propertyID BasePropertyID, from time.Time, to time.Time) string {
	payload := fmt.Sprintf(payloadTemplate, propertyID, from.Format(utils.DateLayout), to.Format(utils.DateLayout))
	return payload
}

const payloadTemplate = `
<PropertyID>%d</PropertyID>
<DateFrom>%s</DateFrom>
<DateTo>%s</DateTo>
`

func validateArguments(from time.Time, to time.Time) error {
	if from.IsZero() {
		return fmt.Errorf("From cannot be zero!")
	}
	if to.IsZero() {
		return fmt.Errorf("to cannot be zero!")
	}

	if to.Before(from) {
		return fmt.Errorf("To cannot be before from")
	}
	return nil
}

func parsePropertyAveragePricePayload(rawPayload []byte) ([]PropertyPrice, error) {

	var parsedPayload PropertyAveragePriceResponse
	errorParsing := xml.Unmarshal(rawPayload, &parsedPayload)
	if errorParsing != nil {
		return nil, errorParsing
	}
	result := make([]PropertyPrice, len(parsedPayload.PropertyPrices))
	for i, price := range parsedPayload.PropertyPrices {
		cleanedPrice := strings.ReplaceAll(price.Price, " ", "")
		cleanedPrice = strings.ReplaceAll(cleanedPrice, "\n", "")
		stayPrice, errConverting := strconv.ParseFloat(cleanedPrice, 64)
		if errConverting != nil {
			log.Printf("Error converting this pricing: %s", errConverting)
			continue
		}
		result[i] = PropertyPrice{
			Price:           stayPrice,
			Cleaning:        price.Cleaning,
			ExtraPerson:     price.ExtraPerson,
			SecurityDeposit: price.SecurityDeposit,
			NumberOfGuests:  price.NumberOfGuests,
		}
	}
	return result, nil
}

type PropertyAveragePriceResponse struct {
	PropertyPrices []PropertyPriceInternal `xml:"PropertyPrice"`
}

type PropertyPrice struct {
	NumberOfGuests  uint8   `json:"numberOfGuests"`
	Cleaning        float64 `json:"cleaning"`
	ExtraPerson     float64 `json:"extraPerson"`
	Deposit         float64 `json:"deposit"`
	SecurityDeposit float64 `json:"securityDeposit"`
	Price           float64 `json:"price"`
}

type PropertyPriceInternal struct {
	NumberOfGuests  uint8   `xml:"NOP,attr"`
	Cleaning        float64 `xml:"Cleaning,attr"`
	ExtraPerson     float64 `xml:"ExtraPersonPrice,attr"`
	Deposit         float64 `xml:"Deposit,attr"`
	SecurityDeposit float64 `xml:"SecurityDeposit,attr"`
	Price           string  `xml:",innerxml"`
}
