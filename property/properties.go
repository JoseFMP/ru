package property

import (
	"encoding/xml"
	"fmt"
	"log"

	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru"
)

type PropertyClient interface {
	ListProperties(includeNLA bool) ([]ru.Property, error)
	ListSpecificProperty(propertyID ru.BasePropertyID) (*SpecificPropertyResponsePayload, error)
}

func GetPropertyClient(userName string, password string) PropertyClient {
	return &propertyClientConfig{
		ClientConfig: ru.ClientConfig{
			UserName: userName,
			Password: password,
		},
	}
}

type propertyClientConfig struct {
	ru.ClientConfig
}

// ListProperties provides a list of properties available.
func (clientConfig *propertyClientConfig) ListProperties(includeNLA bool) ([]ru.Property, error) {

	requestPayload := generateListPropertiesRequestPayload(includeNLA)
	commandResponse, err := clientConfig.DoRequestPull(ru.GetAllCommands().ListProp, requestPayload)
	if err != nil {
		return nil, err
	}
	properties, errParsing := parsePropertiesPayload(commandResponse.Payload.Payload)
	if errParsing != nil {
		log.Printf("Error parsing properties list:\n%s\n", errParsing)
		return nil, errParsing
	}
	return properties, nil
}

const payloadPropertiesTemplate string = `
 <LocationID>%d</LocationID>
 <IncludeNLA>%t</IncludeNLA>
`

var thailandLocationID = ru.GetAllLocations().Thailand.LocationID

func generateListPropertiesRequestPayload(includeNLA bool) string {
	payload := fmt.Sprintf(payloadPropertiesTemplate, thailandLocationID, includeNLA)
	return payload
}

func parsePropertiesPayload(payload []byte) ([]ru.Property, error) {

	//log.Println(string(payload))
	var properties struct {
		Properties []ru.Property `xml:"Property"`
	}
	errUnmarshaling := xml.Unmarshal(payload, &properties)
	if errUnmarshaling != nil {
		log.Printf("Error unmarshaling this:\n%s\n", string(payload))
		return nil, errUnmarshaling
	}
	return properties.Properties, nil
}
