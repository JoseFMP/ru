package property

import (
	"encoding/xml"
	"fmt"

	"github.com/JoseFMP/ru"
)

func (clientConfig *propertyClientConfig) ListSpecificProperty(propertyID ru.BasePropertyID) (*SpecificPropertyResponsePayload, error) {
	commandPayload := createListSpecificPropertyPayload(propertyID)
	commandResponse, errDoingReq := clientConfig.DoRequestPull(ru.GetAllCommands().ListSpecificProperty, commandPayload)

	if errDoingReq != nil {
		return nil, errDoingReq
	}
	parsedPayload, errParsing := parseListSepecificPropertyResponse(commandResponse.Payload.Payload)
	if errParsing != nil {
		return nil, errParsing
	}
	return parsedPayload, nil
}

func parseListSepecificPropertyResponse(payload []byte) (*SpecificPropertyResponsePayload, error) {

	var parsedPayload SpecificPropertyResponsePayload
	errorParsing := xml.Unmarshal(payload, &parsedPayload)
	if errorParsing != nil {
		return nil, errorParsing
	}
	return &parsedPayload, nil
}
func createListSpecificPropertyPayload(propertyID ru.BasePropertyID) string {

	payload := fmt.Sprintf(listSpecificPropertyRequestTemplate, propertyID)
	return payload
}

const listSpecificPropertyRequestTemplate = "<PropertyID>%d</PropertyID>"

type SpecificPropertyResponsePayload struct {
	Name           string     `json:"name"`
	Space          uint64     `json:"space"`
	StandardGuests uint8      `json:"standardGuests"`
	CanSleepMax    uint8      `json:"canSleepMax"`
	Images         []ru.Image `json:"images" xml:">Image"`
}
