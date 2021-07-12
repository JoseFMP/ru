package reservations

import (
	"testing"

	"github.com/JoseFMP/ru"

	"github.com/stretchr/testify/assert"
)

func TestCanParseResponse(t *testing.T) {

	asByteArray := []byte(reservationsPayloadMock)

	reservations, errParsing := parseListReservationsResponse(asByteArray)

	assert.Nil(t, errParsing)
	assert.NotNil(t, reservations)
	assert.Len(t, reservations, 1)
	assert.Equal(t, ru.ReservationID(136468606), reservations[0].ReservationID)
	assert.Equal(t, ReservationStatus(1), reservations[0].Status)
	assert.Equal(t, "2020-06-18 05:40:26", reservations[0].LastMod)
	assert.Len(t, reservations[0].StayInfos, 1)
	assert.Equal(t, "2020-06-18", reservations[0].StayInfos[0].DateFrom)
	assert.Equal(t, "2020-06-20", reservations[0].StayInfos[0].DateTo)
}

func TestCanParse2Reservations(t *testing.T) {

	asByteArray := []byte(twoReservationsMock)

	reservations, errParsing := parseListReservationsResponse(asByteArray)

	assert.Nil(t, errParsing)
	assert.NotNil(t, reservations)
	assert.Len(t, reservations, 2)
	assert.Equal(t, ru.ReservationID(136485319), reservations[1].ReservationID)

}

const reservationsPayloadMock = `<Reservations>
        <Reservation>
            <ReservationID>136468606</ReservationID>
            <StatusID>1</StatusID>
            <LastMod>2020-06-18 05:40:26</LastMod>
            <StayInfos>
                <StayInfo>
                    <PropertyID>2650975</PropertyID>
                    <XmlApartmentID>-1</XmlApartmentID>
                    <DateFrom>2020-06-18</DateFrom>
                    <DateTo>2020-06-20</DateTo>
                    <NumberOfGuests>1</NumberOfGuests>
                    <Costs>
                        <RUPrice>123.0000</RUPrice>
                        <ClientPrice>162.0000</ClientPrice>
                        <AlreadyPaid>162.0000</AlreadyPaid>
                    </Costs>
                    <ResapaID>163221461</ResapaID>
                    <Comments>Cool booking!</Comments>
                    <Units>1</Units>
                </StayInfo>
            </StayInfos>
            <CustomerInfo>
                <Name>Tom</Name>
                <SurName>Cruise</SurName>
                <Email />
                <Phone />
                <MobilePhone />
                <SkypeID />
                <Address />
                <ZipCode />
                <Passport />
            </CustomerInfo>
            <Creator>jose@homa.co</Creator>
            <IsArchived>false</IsArchived>
        </Reservation>
    </Reservations>`

const twoReservationsMock = `<Reservations>
        <Reservation>
            <ReservationID>136468606</ReservationID>
            <StatusID>1</StatusID>
            <LastMod>2020-06-18 05:40:00</LastMod>
            <StayInfos>
                <StayInfo>
                    <PropertyID>2650975</PropertyID>
                    <XmlApartmentID>-1</XmlApartmentID>
                    <DateFrom>2020-06-18</DateFrom>
                    <DateTo>2020-06-20</DateTo>
                    <NumberOfGuests>1</NumberOfGuests>
                    <Costs>
                        <RUPrice>123.0000</RUPrice>
                        <ClientPrice>162.0000</ClientPrice>
                        <AlreadyPaid>162.0000</AlreadyPaid>
                    </Costs>
                    <ResapaID>163221461</ResapaID>
                    <Comments>Cool booking!</Comments>
                    <Units>1</Units>
                </StayInfo>
            </StayInfos>
            <CustomerInfo>
                <Name>Tom</Name>
                <SurName>Cruise</SurName>
                <Email />
                <Phone />
                <MobilePhone />
                <SkypeID />
                <Address />
                <ZipCode />
                <Passport />
            </CustomerInfo>
            <Creator>jose@homa.co</Creator>
            <IsArchived>false</IsArchived>
        </Reservation>
        <Reservation>
            <ReservationID>136485319</ReservationID>
            <StatusID>1</StatusID>
            <LastMod>2020-06-26 02:37:00</LastMod>
            <StayInfos>
                <StayInfo>
                    <PropertyID>2650975</PropertyID>
                    <XmlApartmentID>-1</XmlApartmentID>
                    <DateFrom>2020-10-22</DateFrom>
                    <DateTo>2020-10-31</DateTo>
                    <NumberOfGuests>1</NumberOfGuests>
                    <Costs>
                        <RUPrice>123.0000</RUPrice>
                        <ClientPrice>123.0000</ClientPrice>
                        <AlreadyPaid>123.0000</AlreadyPaid>
                    </Costs>
                    <ResapaID>163237794</ResapaID>
                    <Comments />
                    <Units>1</Units>
                </StayInfo>
            </StayInfos>
            <CustomerInfo>
                <Name>Tom</Name>
                <SurName>Cruise</SurName>
                <Email />
                <Phone />
                <MobilePhone /> 
                <SkypeID />
                <Address />
                <ZipCode />
                <Passport />
            </CustomerInfo>
            <Creator>jose@homa.co</Creator>
            <IsArchived>false</IsArchived>
        </Reservation>
    </Reservations>`
