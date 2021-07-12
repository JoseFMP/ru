package ru

import (
	"fmt"
	"log"
)

// Authenticate authenticates towards RU API
func (clientConfig *ClientConfig) Authenticate() error {

	commandResponse, err := clientConfig.DoRequestPull(commands.ListStatuses, "")

	if err != nil {
		log.Printf("Authentication failed")
		return err
	}
	if commandResponse.Status.ID != allStatuses.Success.ID {
		return fmt.Errorf("Authentication failed, response to authentication message contains error: %d\n%s", commandResponse.Status.ID, commandResponse.Status.Status)
	}

	log.Printf("Authenticate result: %s\n", commandResponse.Status.Status)
	return nil
}

type Authentication struct {
	UserName string `xml:"UserName"`
	Password string `xml:"Password"`
}
