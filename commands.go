package ru

import (
	"encoding/xml"
	"fmt"
	"log"
)

// APICommand is a command in RU API
type APICommand string

type commandsAvailable struct {
	ListStatuses                     APICommand
	ListPropTypes                    APICommand
	ListOTAPropTypes                 APICommand
	ListProp                         APICommand
	ListPropertyAvailabilityCalendar APICommand
	ListReservations                 APICommand
	GetOwnReservations               APICommand
	ModifyStay                       APICommand
	ListSpecificProperty             APICommand
	GetPropertyAveragePrice          APICommand
	PutLead                          APICommand
	PutConfirmedReservationMulti     APICommand
}

var commands = commandsAvailable{
	ListStatuses:                     "ListStatuses",
	ListPropTypes:                    "List_PropTypes",
	ListOTAPropTypes:                 "ListOTAPropTypes",
	ListProp:                         "ListProp",
	ListPropertyAvailabilityCalendar: "ListPropertyAvailabilityCalendar",
	ListReservations:                 "ListReservations",
	GetOwnReservations:               "GetOwnReservations",
	ModifyStay:                       "ModifyStay",
	ListSpecificProperty:             "ListSpecProp",
	GetPropertyAveragePrice:          "GetPropertyAvbPrice",
	PutLead:                          "PutLead",
	PutConfirmedReservationMulti:     "PutConfirmedReservationMulti",
}

func GetAllCommands() commandsAvailable {
	return commands
}

func (command *APICommand) toPullString() string {
	outtestXML := fmt.Sprintf("Pull_%s_RQ", string(*command))
	return outtestXML
}

func (command *APICommand) toPushString() string {
	outtestXML := fmt.Sprintf("Push_%s_RQ", string(*command))
	return outtestXML
}

func (client *ClientConfig) createRequestPullPayload(command APICommand, payload string) string {
	outtestXML := command.toPullString()

	authenticationFragment := createAuthenticationFragment(client.UserName, client.Password)

	return createRequestPayload(outtestXML, payload, authenticationFragment)
}

func (client *ClientConfig) createRequestPushPayload(command APICommand, payload []byte) string {
	outtestXML := command.toPushString()

	authenticationFragment := createAuthenticationFragment(client.UserName, client.Password)

	return createRequestPayload(outtestXML, string(payload), authenticationFragment)
}

func createRequestPayload(outtestXML string, payload string, authenticationFragment string) string {
	payloadAsString := fmt.Sprintf(reqTemplate, outtestXML, authenticationFragment, payload, outtestXML)
	return payloadAsString
}

func createAuthenticationFragment(useName string, password string) string {
	authenticationFragment := fmt.Sprintf(authenticationTemplate, useName, password)
	return authenticationFragment
}

var allStatuses = GetAllStatuses()

func parseCommandResponse(httpPayload *[]byte) (*GenericCommandResponse, error) {
	var ruError Error
	errParsingError := xml.Unmarshal(*httpPayload, &ruError)
	if errParsingError != nil {
		log.Println("Error parsing error. Maybe ok :)")
	}
	if ruError.ID != allStatuses.Success.ID {
		log.Printf("Found RU error is %d.\nError message from RU:%s", ruError.ID, ruError.Message)
		return &GenericCommandResponse{
			Status: Status{
				ID:     ruError.ID,
				Status: ruError.Message,
			},
		}, nil
	}

	var response GenericCommandResponse
	errParsingResponse := xml.Unmarshal(*httpPayload, &response)
	if errParsingResponse != nil {
		return nil, errParsingResponse
	}
	elementName := response.Payload.XMLName.Local
	if elementName != "" {
		fixedResponse := fmt.Sprintf("<%s>%s</%s>", elementName, string(response.Payload.Payload), elementName)
		response.Payload.Payload = []byte(fixedResponse)
	}
	return &response, nil
}

const authenticationTemplate = `<Authentication>
  <UserName>%s</UserName>
  <Password>%s</Password>
</Authentication>`
const reqTemplate = `<%s>
  %s
  %s
</%s>`

// GenericCommandResponse response of a command
type GenericCommandResponse struct {
	Status     Status `xml:"Status"`
	ResponseID ResponseID
	Payload    Payload `xml:",any"`
}

type Payload struct {
	XMLName xml.Name
	Payload []byte `xml:",innerxml"`
}
