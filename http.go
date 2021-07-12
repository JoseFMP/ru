package ru

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (config *ClientConfig) DoRequestPull(command APICommand, requestPayload string) (*GenericCommandResponse, error) {
	return config.doRequestInternal(command, true, []byte(requestPayload))
}

func (config *ClientConfig) DoRequestPushWithBytes(command APICommand, requestPayload []byte) (*GenericCommandResponse, error) {
	return config.doRequestInternal(command, false, requestPayload)
}

func (config *ClientConfig) doRequestInternal(command APICommand, pull bool, requestPayload []byte) (*GenericCommandResponse, error) {

	var commandPayload string
	if pull {
		commandPayload = config.createRequestPullPayload(command, string(requestPayload))
	} else {
		commandPayload = config.createRequestPushPayload(command, requestPayload)

	}

	log.Printf("Raw payload:\n%s", commandPayload)
	reader := bytes.NewReader([]byte(commandPayload))
	request, errCreatingRequest := http.NewRequest(
		http.MethodPost,
		baseURLString,
		reader)

	if errCreatingRequest != nil {
		errorMessage := fmt.Sprintf("Error creating request: %v", errCreatingRequest)
		log.Println(errorMessage)
		return nil, errCreatingRequest
	}
	httpClient := http.Client{}

	log.Printf("About to send request to RU, %s\n", command)
	resp, errDoingReq := httpClient.Do(request)
	if errDoingReq != nil {
		return nil, errDoingReq
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP response code: %d", resp.StatusCode)
	}

	bodyPayload, errReadingBody := ioutil.ReadAll(resp.Body)
	if errReadingBody != nil {
		return nil, errReadingBody
	}

	commandResponse, errParsingRes := parseCommandResponse(&bodyPayload)
	if errParsingRes != nil {
		return nil, errParsingRes
	}
	return commandResponse, nil
}
