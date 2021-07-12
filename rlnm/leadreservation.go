package rlnm

import (
	"encoding/xml"

	"github.com/JoseFMP/ru"
)

type PutLeadReservation struct {
	XMLName        xml.Name          `xml:"LNM_PutLeadReservation_RQ"`
	Authentication ru.Authentication `xml:"Authentication"`
}
