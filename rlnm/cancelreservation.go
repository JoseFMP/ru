package rlnm

import (
	"encoding/xml"

	"github.com/JoseFMP/ru"
)

type PutCancelReservation struct {
	XMLName        xml.Name          `xml:"LNM_CancelReservation_RQ"`
	Authentication ru.Authentication `xml:"Authentication"`
}
