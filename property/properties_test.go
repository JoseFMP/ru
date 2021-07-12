package property

import (
	"testing"

	"dev.azure.com/noon-homa/RentalsUnitedGoClient/_git/ru"

	"github.com/stretchr/testify/assert"
)

func TestGenerateReqPayload(t *testing.T) {
	payload := generateListPropertiesRequestPayload(false)
	assert.NotNil(t, payload)
	assert.Contains(t, payload, "<IncludeNLA>false</IncludeNLA>")
	assert.NotContains(t, payload, "<IncludeNLA>true</IncludeNLA>")

	payload2 := generateListPropertiesRequestPayload(true)
	assert.NotNil(t, payload2)
	assert.Contains(t, payload2, "<IncludeNLA>true</IncludeNLA>")
	assert.NotContains(t, payload2, "<IncludeNLA>false</IncludeNLA>")

}

func TestPropertyParseSimplePayload(t *testing.T) {

	payloadAsByteSlice := []byte(simplePayload)
	properties, errParsing := parsePropertiesPayload(payloadAsByteSlice)

	assert.Nil(t, errParsing, "There should not be an error parsing a simple payload")
	assert.NotNilf(t, properties, "Properties object cannot be nil: +%v", properties)
	assert.Lenf(t, properties, 2, "The number of properties must be 2 but %+v found.", properties)

	firstProperty := (properties)[0]
	assert.NotNil(t, firstProperty)
	assert.Equalf(t, "One Bedroom Apartment", firstProperty.Name, "%+v", firstProperty)
	assert.Equal(t, uint64(553968), firstProperty.OwnerID)
	assert.Equal(t, "2019-03-05", firstProperty.DateCreated)
	assert.NotNil(t, firstProperty.LastMod)
	assert.Equal(t, false, firstProperty.LastMod.NLA)
	assert.Equal(t, "2019-04-26 06:53:45", firstProperty.LastMod.LastMod)

	assert.NotNil(t, firstProperty.ProviderUID)
	assert.Equal(t, int64(-1), firstProperty.ProviderUID.BuildingID)
	assert.Equal(t, ru.BasePropertyID(273641), firstProperty.ProviderUID.ID, "ProviderUID ID")

	assert.NotNil(t, firstProperty.ID)
	assert.Equal(t, int64(-1), firstProperty.ID.BuildingID)
	assert.Equal(t, ru.BasePropertyID(2209989), firstProperty.ID.ID, "ID ID")

}

const simplePayload = `
<Properties>
        <Property>
            <PUID BuildingID="-1">273641</PUID>
            <ID BuildingID="-1">2209989</ID>
            <Name>One Bedroom Apartment</Name>
            <OwnerID>553968</OwnerID>
            <DetailedLocationID TypeID="4">25591</DetailedLocationID>
            <LastMod NLA="false">2019-04-26 06:53:45</LastMod>
            <DateCreated>2019-03-05</DateCreated>
		</Property>
		<Property>
            <PUID BuildingID="-1">12312</PUID>
            <ID BuildingID="-1">12323</ID>
            <Name>Classic</Name>
            <OwnerID>553968</OwnerID>
            <DetailedLocationID TypeID="4">25591</DetailedLocationID>
            <LastMod NLA="false">2019-04-26 06:53:45</LastMod>
            <DateCreated>2019-03-05</DateCreated>
		</Property>
</Properties>
`

const realPayload = `
<Properties>
<Property><ID BuildingID="-1">2650975</ID><Name>Classic</Name><OwnerID>566143</OwnerID><DetailedLocationID TypeID="4">6139</DetailedLocationID><LastMod NLA="false">2020-06-15 07:08:13</LastMod><DateCreated>2020-05-19</DateCreated><UserID>0</UserID></Property><Property><ID BuildingID="-1">2660540</ID><Name>The Glass XL</Name><OwnerID>566143</OwnerID><DetailedLocationID TypeID="4">6139</DetailedLocationID><LastMod NLA="false">2020-06-15 08:08:41</LastMod><DateCreated>2020-06-15</DateCreated><UserID>0</UserID></Property>
</Properties>
`

func TestRealPayload(t *testing.T) {
	payloadAsByteSlice := []byte(realPayload)
	properties, errParsing := parsePropertiesPayload(payloadAsByteSlice)

	assert.Nil(t, errParsing)
	assert.NotNil(t, properties)
}
