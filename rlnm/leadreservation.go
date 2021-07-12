package rlnm

import (
	"encoding/xml"

	
)

type PutLeadReservation struct {
	XMLName        xml.Name          `xml:"LNM_PutLeadReservation_RQ"`
	Authentication ru.Authentication `xml:"Authentication"`
}
