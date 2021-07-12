package ru

type CreditCard struct {
	CCNUmber       CCNumber
	CVC            string
	NameOnCard     string
	Expiration     string
	BillingAddress string
	CardType       CardType
	Comments       string
}

type CardType string
type CCNumber string

type cardTypes struct {
	Visa         CardType
	Mastercard   CardType
	VisaElectron CardType
	MaestroUk    CardType
	Diners       CardType
}

func GetAllCardTypes() cardTypes {
	return cardTypes{
		Diners:       "DINERS",
		MaestroUk:    "MAESTRO_UK",
		Mastercard:   "MASTERVARD",
		Visa:         "VISA",
		VisaElectron: "VISA_ELECTRON",
	}
}
